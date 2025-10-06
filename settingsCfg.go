package main

import (
	"log"
	"os"

	"go.uber.org/zap"
)

func SettingsCfg() {
	if Init_Settings {
		// TODO: Implement settings configuration logic
		log.SetPrefix(green.Render("[SETTINGS] "))
		getWD, err := os.Getwd()
		//fmt.Println("getwd: ", getWD)
		if err != nil {
			zap.L().Info("settings configuration read error: ", zap.Error(err))
			log.Println("settings configuration read error: ", err)

		}
		FolderCheck(getWD, "settings", "[SETTINGS] ")
		FileSettingsCheck(getWD, "settings.go", "[SETTINGS] ")
		FileConfigYamlCheck(getWD, "config.yaml", "[SETTINGS] ")

		// _, errFolder := os.Stat("./settings")
		// if os.IsNotExist(errFolder) {
		// 	os.Mkdir("./settings", 0755)
		// 	log.Println("settings folder created.")
		// 	zap.L().Info("settings folder created.")

		// } else {
		// 	log.Println("settings folder exist, bypass: ", getWD)
		// 	zap.L().Info("settings folder exist, bypass: ", zap.String("path", getWD))
		// }

		// _, errFile := os.Stat(getWD + "/settings/settings.go")
		// if os.IsNotExist(errFile) {
		// 	file, err := os.Create("./settings/settings.go")
		// 	if err != nil {
		// 		log.Println("settings.go file creation error: ", err)
		// 		zap.L().Info("settings.go file creation error: ", zap.Error(err))
		// 	}
		// 	defer file.Close()
		// 	file.Write([]byte(settingsGo))
		// 	log.Println("settings.go file created.")
		// 	zap.L().Info("settings.go file created.")

		// }
		// _, errConfigFile := os.Stat(getWD + "/config.yaml")
		// if os.IsNotExist(errConfigFile) {
		// 	file, err := os.Create("./config.yaml")
		// 	if err != nil {
		// 		log.Println("config.yaml file creation error: ", err)
		// 		zap.L().Info("config.yaml file creation error: ", zap.Error(err))
		// 	}
		// 	defer file.Close()
		// 	file.Write([]byte(configYaml))
		// 	log.Println("config.yaml file created.")
		// 	zap.L().Info("config.yaml file created.")
		// 	return
		// }

	} else {
		// zap.L().Info("settings initialized configuration was not selected.")
		// log.Panicln("settings initialized configuration was not selected.")
	}
}

var (
	configYaml string = `app:
  name: "web_app"
  mode: "dev"
  port: 8080`
	settingsGo string = `package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("viper.ReadInConfig() was error: %s \n", err))
	}
	r := gin.Default()
	if err := r.Run(":8080"); err != nil {
		fmt.Sprintf("Error starting server: %s\n", err)
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file was changed: ", in.Name)
	})
	return err
}`
)
