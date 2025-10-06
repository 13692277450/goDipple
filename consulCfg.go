package main

func ConsulCfg() {

	if Init_Consul {

		FolderCheck("micro/consul", "micro/consul", "[CONSUL] ")
		WriteContentToConfigYaml(Consul_Init_Content, "micro/consul/consul.go", "[CONSUL] ")
		WriteContentToConfigYaml(Consul_Config_Yaml, "micro/consul/config.yaml", "[CONSUL] ")
	}
}

var (
	Consul_Init_Content = `
package main
import (

	"fmt"
	"github.com/hashicorp/consul/api"

)
func registerService() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file

	// Create Consul client
	config := api.DefaultConfig()
	config.Scheme = "http" // Consul scheme (http or https)
	config.Address = fmt.Sprintf("%s:%d", viper.GetString("client_address"), viper.GetInt("port"))
	//config.Address = "127.0.0.1:8500" // Consul address
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	// Define registration
	registration := &api.AgentServiceRegistration{
		ID:      viper.GetString("service_id"),      // Service ID
		Name:    viper.GetString("service_name"),    // Service Name
		Address: viper.GetString("service_address"), // Service Address
		Port:    viper.GetString("service_port"),    // Service Port
		Check: &api.AgentServiceCheck{ // Health check
			HTTP:     viper.GetString("service_health_check_url"),
			Interval: viper.GetString("service_health_check_interval"),
			Timeout:  viper.GetString("service_health_check_timeout"),
		},
	}
	//Register service to  Consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	fmt.Println("Service registered successfully!")
}
func main() {
	registerService()
}
`
	Consul_Config_Yaml = `consul: 
  name: "consul"
  client_address: "127.0.0.1"
  client_port: 8500
  client_scheme: "http"
  client_timeout: 5000
  client_retry: 3
  client_max_retry: 5
  service_name: "consul"
  service_id: "consul"
  service_tags: ["consul"]
  service_address: "127.0.0.1"
  service_port: 8080
  service_health_check_interval: 10
  service_health_check_timeout: 5
  service_health_check_url: "http://192.168.1.100:8080/health"`
)
