package main

import (
	"log"
	"math"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/luthermonson/go-proxmox"
	"github.com/urfave/cli/v2"
)

func basicOperationHelper(vmName string, con *cli.Context, config *Config, client *proxmox.Client) *proxmox.VirtualMachine {
	vm, err := mapVM(vmName, config, client)
	if err != nil {
		log.Fatal(err)
	}

	return vm
}

func tableHelper() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"VMID", "Name", "CPUs", "Disk (GB)", "Mem (GB)", "Status", "Uptime"})
	t.SetStyle(table.StyleLight)

	return t
}

func byteConversion(bytes float64) float64 {
	gibibytes := bytes / math.Pow(1024, 3)
	return gibibytes
}
