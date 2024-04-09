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
	yamlData = regexPlaybook(yamlData)

	err = os.WriteFile("updatePlaybook.yaml", yamlData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	println("updatePlaybook.yaml created successfully!")
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

func regexPlaybook(yamlData []byte) []byte {
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