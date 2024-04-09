package main

import (
	"github.com/L0G1C06/ansiblecli/pkg/ansiblecli"
	"github.com/L0G1C06/ansiblecli/internal/dataInput"
)

func main() {
	//updates := datainput.CreateUpdate()
	//ansiblecli.UpdatePlaybook(&updates)
	inventory := datainput.CreateInventory()
	ansiblecli.Inventory(&inventory)
}
