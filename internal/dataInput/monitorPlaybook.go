package datainput

import "github.com/L0G1C06/ansiblecli/internal/config"

func CreateMonitor() config.Monitoring{
	monitor := config.Monitoring{
		Monitor: []config.MonitoringConfigs{
			{
				Name: "ConectaFacensMonitoring",
				Hosts: "conecta_facens",
				Tasks: []config.TaskConfig{
					{
						Name: "Check Disk",
						Command: "df -hT",
						Register: "disk_memory",
					},
				},
			},
		},
	}
	return monitor
}