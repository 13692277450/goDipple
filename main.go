package main

/*
Version: 0.12
Author: Mang Zhang, Shenzhen China
Release Date: 2024-10-07
Project Name: GoDipple
Description: A tool to help developers quickly create a project structure and initiallize function in Go project.
Copy Rights: MIT License
Email: m13692277450@outlook.com
Mobile: +86-13692277450
HomePage: www.pavogroup.top

*/
import (
	"flag"
	"fmt"
	"os"
)

var (
	CurrentVersion        = "0.12"
	NewVersionIsAvailable = ""
	IsUpgrade             = flag.Bool("upgrade", false, "Run with -upgrade to upgrade new version of GoDipple")
)

func main() {
	// Check and add to PATH if needed
	// if !isInSystemPath() {
	// 	fmt.Println("Adding GoDipple to system PATH...")
	// 	if err := addToSystemPath(); err != nil {
	// 		fmt.Printf("Warning: Failed to add to system PATH (admin rights needed): %v\n", err)
	// 	} else {
	// 		fmt.Println("Successfully added to system PATH")
	// 	}
	// }

	Init()
	go NewVersionCheck()                                 // check for new version
	SignalString += SignalString + NewVersionIsAvailable // check for new version
	if *IsUpgrade {
		DownloadUpgrade() // download new version

		os.Exit(0)
	}
	if _, err := os.Stat("GoDipple.exe.old"); os.IsNotExist(err) {
		fmt.Printf("The old application was not found.\n")
	} else {
		os.Remove("GoDipple.exe.old")
		fmt.Printf("The old application was removed success.\n")
	}

	MainMenu() // main menu
	os.Exit(0)

}
func Init() {
	flag.Parse()
}
