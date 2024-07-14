package main

import (
	"context"
	"fmt"
	"log"

	"github.com/luthermonson/go-proxmox"
)

func mapVM(config *Config, client *proxmox.Client) error {
	vmMap := make(map[string]proxmox.StringOrUint64)

	for _, val := range config.Nodes {
		nodes, err := client.Node(context.TODO(), val)
		if err != nil {
			log.Fatal(err)
		}

		names, err := nodes.VirtualMachines(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		for i := range names {
			vmMap[names[i].Name] = names[i].VMID
		}

		for vmid, name := range vmMap {
			fmt.Printf("vmname: %s and vmid: %d\n", vmid, name)
		}
	}

	return nil
}
