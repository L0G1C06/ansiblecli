package ansiblecli

import (
	"bytes"
	"log"
	"os"

	"github.com/L0G1C06/ansiblecli/internal/config"


	"gopkg.in/yaml.v3"
)

func UpdatePlaybook(update *config.Update) {
	var b bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(1)

	err := yamlEncoder.Encode(&update)
	if err != nil {
		log.Fatal(err)
	}

	yamlData := b.Bytes()
	yamlData = regexUpdatePlaybook(yamlData)

	err = os.WriteFile("update.yaml", yamlData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	println("update.yaml created successfully!")
}

func Inventory(inventory *config.Inventory){
	var b bytes.Buffer 
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(1)

	err := yamlEncoder.Encode(&inventory)
	if err != nil{
		log.Fatal(err)
	}

	yamlData := b.Bytes()
	
	err = os.WriteFile("inventory.yaml", yamlData, 0644)
	if err != nil{
		log.Fatal(err)
	}

	println("inventory.yaml created successfully")
}

func MonitorPlaybook(monitor *config.Monitoring){
	var b bytes.Buffer 
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(1)

	err := yamlEncoder.Encode(&monitor)
	if err != nil{
		log.Fatal(err)
	}

	yamlData := b.Bytes()
	yamlData = regexMonitorPlaybook(yamlData)
	
	err = os.WriteFile("monitor.yaml", yamlData, 0644)
	if err != nil{
		log.Fatal(err)
	}

	println("monitor.yaml created successfully") 
}

func regexUpdatePlaybook(yamlData []byte) []byte {
	yamlData = bytes.Replace(yamlData, []byte("updates:\n"), []byte{}, 1)
	yamlData = bytes.Replace(yamlData, []byte("  - name:"), []byte("- name:"), -1)
	yamlData = bytes.Replace(yamlData, []byte("  - "), []byte("- "), -1)
	yamlData = bytes.Replace(yamlData, []byte("   "), []byte(" "), -1)
	yamlData = bytes.Replace(yamlData, []byte("  - is_centos:"), []byte("    is_centos:"), -1)
	yamlData = bytes.Replace(yamlData, []byte("  - is_ubuntu:"), []byte("    is_ubuntu:"), -1)
	yamlData = bytes.Replace(yamlData, []byte(`'centos'`), []byte("'centos'"), -1)
	yamlData = bytes.Replace(yamlData, []byte(`'ubuntu'`), []byte("'ubuntu'"), -1)
	return yamlData
}

func regexMonitorPlaybook(yamlData []byte) []byte {
	yamlData = bytes.Replace(yamlData, []byte("monitor:\n"), []byte{}, 1)
	yamlData = bytes.Replace(yamlData, []byte("  - name:"), []byte("- name:"), -1)
	yamlData = bytes.Replace(yamlData, []byte("  - "), []byte("- "), -1)
	yamlData = bytes.Replace(yamlData, []byte("   "), []byte(" "), -1)
	yamlData = bytes.Replace(yamlData, []byte("  - msg:"), []byte("    msg:"), -1)
	yamlData = bytes.Replace(yamlData, []byte("  - formatted_metrics:"), []byte("    formatted_metrics:"), -1)
	yamlData = bytes.Replace(yamlData, []byte("  - webhook_id:"), []byte("    webhook_id:"), -1)
	yamlData = bytes.Replace(yamlData, []byte("  webhook_token:"), []byte("    webhook_token:"), -1)
	yamlData = bytes.Replace(yamlData, []byte("  content:"), []byte("    content:"), -1)
	return yamlData
}