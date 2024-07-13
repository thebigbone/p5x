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

	username := config.Node.Username
	password := config.Node.Password

	credentials := proxmox.Credentials{
		Username: username,
		Password: password,
	}

	client := proxmox.NewClient("https://192.168.31.102:8006/api2/json",
		proxmox.WithHTTPClient(&insecureHTTPClient),
		proxmox.WithCredentials(&credentials),
	)
	var node_names []string

	node_names = append(node_names, "local")
	node_names = append(node_names, "home")

	for i := range node_names {
		nodes, err := client.Node(context.TODO(), node_names[i])
		if err != nil {
			panic(err)
		}

		fmt.Println(nodes.CPUInfo.CPUs)
	}
}
