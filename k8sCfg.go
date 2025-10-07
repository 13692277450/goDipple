package main

func K8sCfg() {

	if Init_K8s {

		FolderCheck("k8s/config", "k8s/config", "[K8S] ")
		WriteContentToConfigYaml(K8s_KubeConfig, "k8s/config/kubeConfig.go", "[K8S] ")
		WriteContentToConfigYaml(K8s_Config_Yaml, "k8s/config/config.yaml", "[K8S] ")
		WriteContentToConfigYaml(K8s_ClientSet, "k8s/config/client.go", "[K8S] ")
		WriteContentToConfigYaml(K8s_Deployment_Yaml, "k8s/config/deployment.yaml", "[K8S] ")
		WriteContentToConfigYaml(K8s_DynamicClient, "k8s/config/dynamicClient.go", "[K8S] ")
		WriteContentToConfigYaml(K8s_DiscoveryClient, "k8s/config/discoveryClient.go", "[K8S] ")
	}
}

var (
	K8s_Config_Yaml     = ``
	K8s_DiscoveryClient = `package main

import (
        "fmt"
        "k8s-clientset/config"
)

func K8s_DiscoveryClient() {
        client := config.NewK8sConfig().InitDiscoveryClient()
        // Read gvr
        preferredResources, _ := client.ServerPreferredResources()
        for _, pr := range preferredResources {
                fmt.Println(pr.String())
        }

        // _, _, _ = client.ServerGroupsAndResources()

}`
	K8s_DynamicClient = `package main

import (
   "context"
   _ "embed"
   "k8s-clientset/config"
   metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
   "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
   "k8s.io/apimachinery/pkg/runtime/schema"
   "k8s.io/apimachinery/pkg/util/yaml"
   "log"
)


// Deployment model file
//go:embed k8s/config/deployment.yaml
var deployTpl string

// dynamic client client Deploy
func K8s_DynamicClient()  {

   // Dynamic client init
   dynamicCli := config.NewK8sConfig().InitDynamicClient()

   // create deploy
   deployGVR := schema.GroupVersionResource{
      Group: "apps",
      Version: "v1",
      Resource: "deployments",
   }

   deployObj := &unstructured.Unstructured{}
   if err := yaml.Unmarshal([]byte(deployTpl), deployObj); err != nil {
       log.Fatalln(err)
   }

   if _, err = dynamicCli.
      Resource(deployGVR).
      Namespace("default").
      Create(context.Background(), deployObj, metav1.CreateOptions{}); 
      err != nil {
      log.Fatalln(err)
   }

   log.Println("Deploy created successfully!")
}`
	K8s_Deployment_Yaml = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: mynginx
  namespace: default
spec:
  selector:
    matchLabels:
      app: mynginx
  replicas: 1
  template:
    metadata:
      labels:
        app: mynginx
    spec:
      containers:
        - name: myngx-container
          image: nginx:1.2-alpine
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: 8080`

	K8s_ClientSet = `func K8s_ClientSet() {
    // Loading configures and clients

    cliset := NewK8sConfig().InitClient()
    configMaps, err := cliset.CoreV1().ConfigMaps(ns).List(metav1.ListOptions{})
    if err != nil {
       panic(err)
    }
    for _, cm := range configMaps.Items {
       fmt.Printf("configName: %v, configData: %v \n", cm.Name, cm.Data)
    }
    return nil
}`

	K8s_KubeConfig = `package config

import (
	"log"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const kubeConfigFilePath = "./config.yaml"

type K8sConfig struct {
}

func NewK8sConfig() *K8sConfig {
        return &K8sConfig{}
}
// Read kubeconfig file
func (this *K8sConfig) K8sRestConfig() *rest.Config {
        config, err := clientcmd.BuildConfigFromFlags("", kubeConfigFilePath)

        if err != nil {
                log.Fatal(err)
        }

        return config
}
// Initiallize clientSet
func (this *K8sConfig) InitClient() *kubernetes.Clientset {
        c, err := kubernetes.NewForConfig(this.K8sRestConfig())

        if err != nil {
                log.Fatal(err)
        }

        return c
}

// Initiallize dynamicClient
func (this *K8sConfig) InitDynamicClient() dynamic.Interface {
        c, err := dynamic.NewForConfig(this.K8sRestConfig())

        if err != nil {
                log.Fatal(err)
        }

        return c
}

// Initiallize DiscoveryClient
func (this *K8sConfig) InitDiscoveryClient() *discovery.DiscoveryClient {
        return discovery.NewDiscoveryClient(this.InitClient().RESTClient())
}`
)