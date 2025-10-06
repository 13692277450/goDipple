package main

import (
	"log"
	"os"

	"go.uber.org/zap"
)

func FolderCheck(FolderPath, FolderName, LogParameter string) {
	log.SetPrefix(green.Render(LogParameter))
	_, errFolder := os.Stat(FolderName)
	if os.IsNotExist(errFolder) {
		os.MkdirAll(FolderName, 0755)
		log.Printf("%s folder created. \n", FolderName)
		zap.L().Info("folder created." + FolderName)

	} else {
		log.Printf("%s folder exist, bypass. \n", FolderName)
		zap.L().Info("folder exist, bypass: "+FolderName, zap.String("path", FolderPath))
	}

}

func FileSettingsCheck(FolderPath, FileName, LogParameter string) {
	log.SetPrefix(green.Render(LogParameter))
	_, errFile := os.Stat(FolderPath + "/settings/" + FileName)
	if os.IsNotExist(errFile) {
		file, err := os.Create("./settings/" + FileName)
		if err != nil {
			log.Println(red.Render("%s file creation error: \n")+FileName, err)
			zap.L().Info("file creation error: "+FileName, zap.Error(err))
		}
		defer file.Close()
		file.Write([]byte(settingsGo))
		log.Printf("%s file created. \n", FileName)
		zap.L().Info("file created: " + FileName)

	} else {
		log.Printf("%s file exist, bypass. \n", FileName)
		zap.L().Info("file exist, bypass: "+FileName, zap.String("path", FolderPath))
	}

}

func FileConfigYamlCheck(FolderPath, FileName, LogParameter string) {
	log.SetPrefix(green.Render(LogParameter))
	_, errConfigFile := os.Stat(FolderPath + "/" + FileName)
	if os.IsNotExist(errConfigFile) {
		file, err := os.Create("./" + FileName)
		if err != nil {
			log.Printf(red.Render("%s file creation error: \n")+FileName, err)
			zap.L().Info("file creation error: "+FileName, zap.Error(err))
		}

		defer file.Close()
		fileInfo, err := file.Stat()
		if err == nil && fileInfo.Size() == 0 {
			file.Write([]byte(configYaml))
		} else {
			file.Write([]byte("\n\n" + configYaml))
		}

		log.Printf("%s file wrote done. \n", FileName)
		zap.L().Info("file created: " + FileName)
		return
	} else {
		file, err := os.OpenFile("./"+FileName, os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Printf("%s file openfailure. \n", FileName)
		}
		log.Printf("%s file exist, bypass file create. \n", FileName)
		file.Write([]byte("\n\n" + configYaml))
		defer file.Close()
		log.Printf("%s file created. \n", FileName)
		zap.L().Info("file exist, bypass create file: "+FileName, zap.String("path", FolderPath))
	}
}
func WriteContentToConfigYaml(Content string, FileName, LogParameter string) {
	log.SetPrefix(green.Render(LogParameter))
	file, err := os.OpenFile("./"+FileName, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(red.Render("%s file open error. \n")+FileName, err)
		zap.L().Info("file open error: "+FileName, zap.Error(err))
	}
	defer file.Close()
	// if FileName == "mongoDB.yaml" {
	// 	file.Write([]byte(Content))
	// 	return
	// }
	// if LogParameter == "prettyshutdown.go" {
	// 	file.Write([]byte(Content))
	// 	return
	// }
	fileInfo, err := file.Stat()
	if err == nil && fileInfo.Size() == 0 {
		file.Write([]byte(Content))
	} else {
		file.Write([]byte("\n\n" + Content))
	}
	log.Printf("Write %s configuration to %s success. \n", LogParameter, FileName)
}
