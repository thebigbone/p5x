package main

import (
	"context"
	"fmt"
	"log"

	"github.com/luthermonson/go-proxmox"
	"github.com/urfave/cli/v2"
)

func stop(con *cli.Context, config *Config, client *proxmox.Client) error {
	vmName := con.Args().First()

	fmt.Printf("stopping %s\n", vmName)

	vm, err := mapVM(vmName, config, client)
	if err != nil {
		log.Fatal(err)
	}

	_, err = vm.Stop(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func start(con *cli.Context, config *Config, client *proxmox.Client) error {
	vmName := con.Args().First()

	fmt.Printf("starting %s\n", vmName)

	vm, err := mapVM(vmName, config, client)
	if err != nil {
		log.Fatal(err)
	}

	_, err = vm.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
