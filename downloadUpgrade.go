package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/cheggaaa/pb/v3"
)

var cyan2 = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF"))

func downloadFile(url, filepath string) error {
	// HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Download new file was failure.")

		return err
	}
	defer resp.Body.Close()

	// Get file size
	size := resp.ContentLength

	// creat progress bar
	bar := pb.Full.Start64(size)
	defer bar.Finish()

	// create temp file first
	//tempFile := filepath + ".tmp"
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Create tmp file was failure.")

		return err
	}
	defer file.Close()

	// create writer with progress bar
	writer := bar.NewProxyWriter(file)

	// write file
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		fmt.Println("Create new file was failure.")
		return err
	}
	time.Sleep(2 * time.Second)
	return nil
}

// DownloadUpgrade
func DownloadUpgrade() {
	url := "http://www.pavogroup.top/software/godipple/GoDipple.exe"
	filepath := "./GoDipple.exe"
	go func() {
		fmt.Println(cyan2.Render("Starting download upgrade from: ", url+"\n"))
		for i := 1; i < 10; i++ {
			fmt.Print(".")
			time.Sleep(500 * time.Millisecond)
		}
	}()
	tempFile := filepath + ".tmp" // save version as .tmp file
	err := downloadFile(url, tempFile)
	if err != nil {
		fmt.Println("Download new file error.")
	}

	// Rename current executable to .old and rename the new one to current executable
	oldFile := filepath + ".old"
	os.Remove(oldFile) // Remove old backup if exists
	if err := os.Rename(filepath, oldFile); err != nil {
		fmt.Println("Move to old file was failure.")
	}

	// Create update batch script
	batchContent := `@echo off
timeout /t 2 /nobreak >nul
move /Y "` + tempFile + `" "` + filepath + `"
del "%~f0"
`
	batchFile := "update.bat"
	if err := os.WriteFile(batchFile, []byte(batchContent), 0755); err != nil {
		fmt.Println("Running batch file was failure.")
	}
	// Run the batch file and wait for completion
	fmt.Println("Executing update script...")
	cmd := exec.Command("cmd.exe", "/C", batchFile)
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error executing update script: %v\n", err)
	}
	time.Sleep(4 * time.Second)

	// Verify the new file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Printf("Error: New file %s was not created\n", filepath)
	}

	fmt.Println("Update completed successfully. Old version saved as " + oldFile + ", the old version will automaticaly be removed when application launch next time.")
	os.Exit(0) // Exit the program after successful update
}
