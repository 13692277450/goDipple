package main

import "fmt"

func ProjectFolderStructureCfg() {

	if Init_ProjectFolderStructure {

		FolderCheck("cmd", "cmd", "[PROJECT FOLDER] ")
		FolderCheck("config", "config", "[PROJECT FOLDER] ")
		FolderCheck("handler", "handler", "[PROJECT FOLDER] ")
		FolderCheck("middleware", "middleware", "[PROJECT FOLDER] ")
		FolderCheck("static", "static", "[PROJECT FOLDER] ")
		FolderCheck("assets", "assets", "[PROJECT FOLDER] ")
		FolderCheck("api", "api", "[PROJECT FOLDER] ")
		FolderCheck("controller", "controller", "[PROJECT FOLDER] ")
		FolderCheck("model", "model", "[PROJECT FOLDER] ")
		FolderCheck("router", "router", "[PROJECT FOLDER] ")
		FolderCheck("service", "service", "[PROJECT FOLDER] ")
		FolderCheck("util", "util", "[PROJECT FOLDER] ")
		FolderCheck("pkgs", "pkgs", "[PROJECT FOLDER] ")
		FolderCheck("test", "test", "[PROJECT FOLDER] ")
		fmt.Println("Project folders structure created. ")
	}
}
