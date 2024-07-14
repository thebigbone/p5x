package main

import (
	"crypto/tls"
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

	mapVM(config, client)
}
