package main

import (
	"bufio"
	"embed"
	"io"
	"strings"
)

//go:embed "windowsRegistry.txt"
var windowsRegistry embed.FS
func WindowsRegistryCfg() {
	if Init_WindowsRegistry {
		FolderCheck("util/winReg", "util/winReg", "[WINDOWSREGISTRY] ")
		
		// 先解析文件内容
		data, _ := windowsRegistry.ReadFile("windowsRegistry.txt")
		reader := bufio.NewReader(strings.NewReader(string(data)))
		var results []string

		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			
			lineStr := string(line)
			results = append(results, lineStr)
			
			switch lineStr {
			case "//WindowsRegistry_Config_Read_End":
				WindowsRegistry_Config_Read = strings.Join(results, "\n")
				results = []string{}
			case "//WindowsRegistry_Config_Write_End":
				WindowsRegistry_Config_Write = strings.Join(results, "\n")
				results = []string{}
			case "//WindowsRegistry_Config_Start_End":
				WindowsRegistry_Config_Start = strings.Join(results, "\n")
				results = []string{}
			}
		}

		// 然后写入文件
		WriteContentToConfigYaml(WindowsRegistry_Config_Read, "util/winReg/windowsRegistryRead.go", "[WINDOWSREGISTRY] ")
		WriteContentToConfigYaml(WindowsRegistry_Config_Write, "util/winReg/windowsRegistryWrite.go", "[WINDOWSREGISTRY] ")
		WriteContentToConfigYaml(WindowsRegistry_Config_Start, "util/winReg/windowsRegistryStart.go", "[WINDOWSREGISTRY] ")
	}
}

var WindowsRegistry_Config_Read = ""

var WindowsRegistry_Config_Write = ""

var WindowsRegistry_Config_Start =""
