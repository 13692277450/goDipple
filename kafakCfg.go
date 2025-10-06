package main

import (
	"embed"
	"log"
	"os"

	"go.uber.org/zap"
)

//go:embed "kafkaCfgTxt.txt"

var staticFiles embed.FS

func KafkaCfg() {

	if Init_Kafka {
		FolderCheck("util/kafka", "util/kafka", "[KAFKA] ")
		FileKafkaCheck("util/kafka", "kafka.go", "[KAFKA] ")
	}

}

func FileKafkaCheck(FolderPath, FileName, LogParameter string) {
	log.SetPrefix(green.Render(LogParameter))
	data, _ := staticFiles.ReadFile("kafkaCfgTxt.txt")
	_, errFile := os.Stat(FolderPath + "/" + FileName)
	if os.IsNotExist(errFile) {
		file, err := os.Create("./util/kafka/" + FileName)
		if err != nil {
			log.Println(red.Render("%s file creation error: \n")+FileName, err)
			zap.L().Info("file creation error: "+FileName, zap.Error(err))
		}
		defer file.Close()
		file.Write([]byte(data))
		log.Printf("%s file created. \n", FileName)
		zap.L().Info("file created: " + FileName)
	} else {
		log.Printf("%s file exist, bypass. \n", FileName)
		zap.L().Info("file exist, bypass: "+FileName, zap.String("path", FolderPath))
	}

}
