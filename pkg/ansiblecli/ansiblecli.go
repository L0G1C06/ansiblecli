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
	yamlData = bytes.Replace(yamlData, []byte("updates:\n"), []byte{}, 1)
	yamlData = bytes.Replace(yamlData, []byte("  - name:"), []byte("- name:"), -1)
	yamlData = bytes.Replace(yamlData, []byte("  - "), []byte("- "), -1)
	yamlData = bytes.Replace(yamlData, []byte("   "), []byte(" "), -1)

	err = os.WriteFile("updatePlaybook.yaml", yamlData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	println("YAML file created successfully!")
}
