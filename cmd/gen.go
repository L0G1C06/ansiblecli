/*
Copyright © 2024 NAME HERE eduwmaldaner@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/L0G1C06/ansiblecli/internal/dataInput"
	"github.com/L0G1C06/ansiblecli/pkg/ansiblecli"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "This command will generate a .yaml file to ansible",
	Long: `This command will generate an ansible .yaml file from 3 options, monitorPlaybook, updatePlaybook and inventory`,
	Run: func(cmd *cobra.Command, args []string) {
		option := ""
		if len(args) >= 1 && args[0] != ""{
			option = args[0]
			if option == "update"{
				updates := datainput.CreateUpdate()
				ansiblecli.UpdatePlaybook(&updates)
			} else if option == "monitor"{
				monitor := datainput.CreateMonitor()
				ansiblecli.MonitorPlaybook(&monitor)
			} else if option == "inventory"{
				inventory := datainput.CreateInventory()
				ansiblecli.Inventory(&inventory)
			} else {
				fmt.Println("Error: " + option + " not exists\n")
				fmt.Println("The options are: \n update\n monitor\n inventory")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
