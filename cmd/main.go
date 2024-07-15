package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/luthermonson/go-proxmox"
	"github.com/urfave/cli/v2"
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

	app := &cli.App{
		Name:  "p5x",
		Usage: "proxmox tui and cli",
		Commands: []*cli.Command{
			{
				Name:  "stop",
				Usage: "stop the VM with name",
				Action: func(con *cli.Context) error {
					return stop(con, config, client)
				},
			},
			{
				Name:  "start",
				Usage: "start the VM with name",
				Action: func(con *cli.Context) error {
					return start(con, config, client)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
