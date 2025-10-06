package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

func NewVersionCheck() {
	tr := &http.Transport{
		MaxIdleConns:          5,
		IdleConnTimeout:       30 * time.Second,
		DisableCompression:    true,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 3 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
	}
	client := &http.Client{Transport: tr}
	url := "http://www.pavogroup.top/software/godipple/version.json"
	resp, err := client.Get(url)
	if err != nil {
		log.Error("New version check Error:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var data map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			log.Info("New version Error decoding JSON:", err)
			return
		}
		NewVersion, ok := data["version"]
		if !ok {
			log.Info("New version check: Invalid version format")
			return
		}
		if NewVersion != CurrentVersion {
			NewVersionIsAvailable = "A new version is available, pls run GoDipple -upgrade to update."
		}
	}
}
