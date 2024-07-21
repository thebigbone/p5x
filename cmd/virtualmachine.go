package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
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
			t.AppendRows([]table.Row{
				{vm[i].VMID, vm[i].Name, vm[i].CPUs, vm[i].Disk, vm[i].Mem, vm[i].Status, vm[i].Uptime},
			})
		}

		t.Render()
	case *proxmox.VirtualMachine:
		t := tableHelper()
		t.AppendRows([]table.Row{
			{vm.VMID, vm.Name, vm.CPUs, vm.Disk, vm.Mem, vm.Status, vm.Uptime},
		})
		t.Render()
	}
}

func tableHelper() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"VMID", "Name", "CPUs", "Disk", "Mem", "Status", "Uptime"})
	t.SetStyle(table.StyleLight)

	return t
}
