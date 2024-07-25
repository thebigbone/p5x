package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/luthermonson/go-proxmox"
	"github.com/urfave/cli/v2"
)

func stop(con *cli.Context, config *Config, client *proxmox.Client) error {
	vmName := con.Args().First()
	fmt.Printf("stopping %s\n", vmName)

	vm := basicOperationHelper(vmName, con, config, client)

	_, err := vm.Stop(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func shutdown(con *cli.Context, config *Config, client *proxmox.Client) error {
	vmName := con.Args().First()
	fmt.Printf("shutting down %s\n", vmName)

	vm := basicOperationHelper(vmName, con, config, client)

	_, err := vm.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func start(con *cli.Context, config *Config, client *proxmox.Client) error {
	vmName := con.Args().First()
	fmt.Printf("starting %s\n", vmName)

	vm := basicOperationHelper(vmName, con, config, client)

	_, err := vm.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func info(con *cli.Context, config *Config, client *proxmox.Client) error {
	vmName := con.Args().First()

	if vmName == "" {
		for _, val := range config.Nodes {
			names, _ := vmNames(client, val)
			tableDisplay(names)
		}
	} else {
		vm, err := mapVM(vmName, config, client)
		if err != nil {
			log.Fatal(err)
		}

		tableDisplay(vm)
	}
	return nil
}

func tableDisplay(x interface{}) {
	switch vm := x.(type) {
	case proxmox.VirtualMachines:
		t := tableHelper()

		for i := range vm {
			mem := fmt.Sprintf("%.2f", byteConversion(float64(vm[i].MaxMem)))
			disk := fmt.Sprintf("%.2f", byteConversion(float64(vm[i].MaxDisk)))

			t.AppendRows([]table.Row{
				{vm[i].VMID, vm[i].Name, vm[i].CPUs, disk, mem, vm[i].Status, vm[i].Uptime},
			})
		}

		t.Render()
	case *proxmox.VirtualMachine:
		t := tableHelper()

		mem := fmt.Sprintf("%.2f", byteConversion(float64(vm.MaxMem)))
		disk := fmt.Sprintf("%.2f", byteConversion(float64(vm.MaxDisk)))

		t.AppendRows([]table.Row{
			{vm.VMID, vm.Name, vm.CPUs, disk, mem, vm.Status, vm.Uptime},
		})
		t.Render()
	}
}
