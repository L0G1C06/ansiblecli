package datainput

import "github.com/L0G1C06/ansiblecli/internal/config"

func CreateMonitor() config.Monitoring{
	monitor := config.Monitoring{
		Monitor: []config.MonitoringConfigs{
			{
				Name: "MonitoringTask",
				Hosts: "InventoryHostName",
				Tasks: []config.TaskConfig{
					{
						Name: "Check Disk",
						Command: "df -hT",
						Register: "disk_memory",
					},
					{
						Name: "Format DIsk Metrics",
						Set_fact: []config.Fact{
							{
								Formatted_metrics: `{% for line in disk_memory.stdout_lines[1:] %} {{ line.split()[5] }}{mount="{{ line.split()[6] }}"} {{ line.split()[1] | regex_replace('G', '') | float}}`,
							},
						},
					},
					{
						Name: "Print Disk Metrics",
						Debug: []config.DebugConfig{
							{
								Message: `{{ disk_memory.stdout_lines[1].split()[5] | int}}`,
							},
						},
					},
					{
						Name: "Send Alert",
						DiscordMessage: []config.DiscordNotifier{
							{
								ID: `DiscordChatID`,
								Token: `DiscordChatToken`,
								Content: `Your service is using more than 95% of the disk`,
							},
						},
						When: "disk_memory.stdout_lines[1].split()[5] | int >= 95",
					},
					{
						Name: "Print normal disk size",
						Debug: []config.DebugConfig{
							{
								Message: `Disk is not full. Disk usage: {{ disk_memory.stdout_lines[1].split()[5] }}`,
							},
						},
						When: "disk_memory.stdout_lines[1].split()[5] | int < 95", 
					},
				},
			},
		},
	}
	return monitor
}