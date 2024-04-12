package datainput

import (
	"github.com/L0G1C06/ansiblecli/internal/config"
)

func CreateInventory() config.Inventory{
	config := config.Inventory{
		HostGroupName: struct {
			Hosts map[string]struct {
				AnsibleConnection       string `yaml:"ansible_connection"`
				AnsibleUser             string `yaml:"ansible_user"`
				AnsibleSSHPrivateKeyFile string `yaml:"ansible_ssh_private_key_file"`
			} `yaml:"hosts"`
		}{
			Hosts: map[string]struct {
				AnsibleConnection       string `yaml:"ansible_connection"`
				AnsibleUser             string `yaml:"ansible_user"`
				AnsibleSSHPrivateKeyFile string `yaml:"ansible_ssh_private_key_file"`
			}{
				"your-host.com": {
					AnsibleConnection:       "",
					AnsibleUser:             "",
					AnsibleSSHPrivateKeyFile: "",
				},
				"your-host2.com": {
					AnsibleConnection:       "",
					AnsibleUser:             "",
					AnsibleSSHPrivateKeyFile: "",
				},
			},
		},
	}
	return config
}