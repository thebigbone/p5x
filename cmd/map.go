package main

import (
	"context"
	"log"

	"github.com/luthermonson/go-proxmox"
)

func mapVM(vmName string, config *Config, client *proxmox.Client) (*proxmox.VirtualMachine, error) {
	// vmName := con.Args().First()

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

		vmID := vmMap[vmName]

		vm, err := nodes.VirtualMachine(context.Background(), int(vmID))
		if err != nil {
			log.Fatal(err)
		}
		// _, err = vm.Shutdown(context.Background())

		return vm, nil
	}

	return &proxmox.VirtualMachine{}, nil
}
