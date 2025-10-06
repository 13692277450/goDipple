package main

func MongoDBCfg() {

	if Init_MongoDB {

		FolderCheck("dao/mongodb", "dao/mongodb", "[MONGODB] ")
		WriteContentToConfigYaml(MongoDB_Init_Content, "dao/mongodb/mongodb.go", "[MONGODB] ")
		WriteContentToConfigYaml(MongoDB_Config_Yaml, "mongoDB.yaml", "[MONGODB] ")
	}
}

var (
	MongoDB_Config_Yaml = `systemLog: 
 destination: file
 path: "/var/log/mongo/mongod.log"
 quiet: true
 logAppend: true
 timeStampFormat: iso8601-utc
storage: 
 dbPath: "/var/lib/mongo"
 directoryPerDB: true
 indexBuildRetry: false
 preallocDataFiles: true
 nsSize: 16
# quota:
#  enforced: false
#  maxFilesPerDB: 8
 smallFiles: false
 syncPeriodSecs: 60
# repairPath: "/var/lib/mongo/_tmp"
 journal:
  enabled: true
#  debugFlags: 1
  commitIntervalMs: 100
processManagement: 
 fork: true
 pidFilePath: "/var/run/mongodb/mongod.pid"
net: 
 bindIp: 192.168.1.111
 port: 27017
 http:
  enabled: true
  RESTInterfaceEnabled: false 
# ssl:
#  mode: "requireSSL"
#  PEMKeyFile: "/etc/ssl/mongodb.pem"
operationProfiling:
 slowOpThresholdMs: 100 
 mode: "slowOp"
security:
 keyFile: "/var/lib/mongo/mongodb-keyfile"
 clusterAuthMode: "keyFile"
 authorization: "disabled"
replication:
 oplogSizeMB: 50
 replSetName: "repl_test"
 secondaryIndexPrefetch: "all"
 `
	MongoDB_Init_Content = `package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func initDB() (err error) {
	clientOptions := options.Client().ApplyURI("mongodb://ip:port")
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Failure to initiallize！,err:%v\n", err)
		return
	} else {
		fmt.Println("Connected to MongoDB!")
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		fmt.Println("Disconnected from MongoDB!")
	}()
}
`
)
