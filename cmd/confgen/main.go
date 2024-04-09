package main

import (
	"github.com/L0G1C06/ansiblecli/internal/config"
	"github.com/L0G1C06/ansiblecli/pkg/ansiblecli"
)

func main() {
	updates := config.Update{
		Updates: []config.UpdateConfigs{
			{
				Name:  "updateJenkinsMachinesPlaybook",
				Hosts: "jenkins_hosts",
				Tasks: []config.TaskConfig{
					{
						Name: "Check if the system is CentOS",
						Set_fact: []config.Fact{
							{
								IsCentOS: "{{ ansible_distribution | lower == 'centos' }}",
							},
						},
						IgnoreErrors: true,
					},
					{
						Name: "Check if the system is Ubuntu",
						Set_fact: []config.Fact{
							{
								IsUbuntu: "{{ ansible_distribution | lower == 'ubuntu' }}",
							},
						},
						IgnoreErrors: true,
					},
					{
						Name:    "Update CentOS Machines",
						Command: "sudo dnf update",
						When:    "is_centos | bool",
					},
					{
						Name:    "Update Ubuntu Machines",
						Command: "sudo apt update",
						When:    "is_ubuntu | bool",
					},
				},
			},
		},
	}
	ansiblecli.UpdatePlaybook(&updates)
}
