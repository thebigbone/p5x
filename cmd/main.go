package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/luthermonson/go-proxmox"
)

func main() {
	config, err := parseConfig("../config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	insecureHTTPClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	username := config.Credentials.Username
	password := config.Credentials.Password

	credentials := proxmox.Credentials{
		Username: username,
		Password: password,
	}

	client := proxmox.NewClient(config.Credentials.Url,
		proxmox.WithHTTPClient(&insecureHTTPClient),
		proxmox.WithCredentials(&credentials),
	)

	for _, val := range config.Nodes {
		nodes, err := client.Node(context.TODO(), val)
		if err != nil {
			panic(err)
		}

		fmt.Println(nodes.CPUInfo.CPUs)
	}
}
