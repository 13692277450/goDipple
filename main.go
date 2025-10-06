package main

//tjikko, dipple, panda
import (
	"flag"
	"fmt"
	"os"
)

var (
	CurrentVersion        = "0.1"
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

// var upgrade string
// var help string

// var rootCmd = &cobra.Command{
// 	Use:   "-u",
// 	Short: "upgrade goDipple.exe",
// 	Long:  `Run goDipple.exe -u to upgrade`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		DownloadUpgrade()
// 	},
// }

// func Execute() {
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		panic(err)
// 		// os.Exit(1)
// 	}
// }

// var helpCmd = &cobra.Command{
// 	Use:   "-h",
// 	Short: "Run goDipper -u to upgrade goDipple",
// 	Long:  `Run goDipple.exe -u to upgrade, for issue, pls access www.pavogroup.top or https://github.com/13692277450/`,
// 	Run: func(cmd *cobra.Command, args []string) {

// 	},
// }

//	func ExecuteHelp() {
//		err := helpCmd.Execute()
//		if err != nil {
//			os.Exit(1)
//		}
//	}
func Init() {
	flag.Parse()
}
