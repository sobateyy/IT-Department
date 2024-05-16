package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type Result struct {
	Hostname    string
	BiosVersion string
	UserProfile string
}

func getWMIC(hostname, class, property string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "wmic", "/node:"+hostname, class, "get", property)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func main() {
	hostnames := make([]string, 0, 21)
	for i := 1000; i <= 1500; i++ {
		hostnames = append(hostnames, fmt.Sprintf("HOST%d", i))
	}
	for i := 5000; i <= 5500; i++ {
		hostnames = append(hostnames, fmt.Sprintf("HOST%d", i))
	}
	hostnames = append(hostnames, fmt.Sprintf("HOST3000"))
	hostnames = append(hostnames, fmt.Sprintf("HOST4000"))

	outputFile := "BIOSinventory.csv"
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"HOSTNAME", "BIOS_VERSION", "USERPROFILE"})
	if err != nil {
		fmt.Println("Cannot write to file", err)
		return
	}

	results := make(chan Result, len(hostnames))
	var wg sync.WaitGroup

　// Limit the number of goroutines to be executed at the same time
　// Match the specs of the host you want to run on
	sem := make(chan struct{}, 100)

	for _, hostname := range hostnames {
		wg.Add(1)
		go func(hostname string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }() 

			biosVersion, err := getWMIC(hostname, "bios", "SMBIOSBIOSVersion")
			if err != nil {
				biosVersion = "ERROR: " + err.Error()
			}

			userProfile, err := getWMIC(hostname, "computersystem", "username")
			if err != nil {
				userProfile = "ERROR: " + err.Error()
			} else {
				split := strings.Split(userProfile, "\\")
				userProfile = split[len(split)-1]
			}

			results <- Result{Hostname: hostname, BiosVersion: biosVersion, UserProfile: userProfile}
		}(hostname)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		err = writer.Write([]string{result.Hostname, result.BiosVersion, result.UserProfile})
		if err != nil {
			fmt.Println("Cannot write to file", err)
			return
		}
	}
}
