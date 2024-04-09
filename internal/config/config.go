package config

type Fact struct {
	IsCentOS string `yaml:"is_centos,omitempty"`
	IsUbuntu string `yaml:"is_ubuntu,omitempty"`
}

type TaskConfig struct {
	Name         string `yaml:"name,omitempty"`
	Command      string `yaml:"command,omitempty"`
	When         string `yaml:"when,omitempty"`
	Set_fact     []Fact `yaml:"set_fact,omitempty"`
	IgnoreErrors bool   `yaml:"ignore_errors,omitempty"`
}

// Playbook structs
type UpdateConfigs struct {
	Name  string       `yaml:"name,omitempty"`
	Hosts string       `yaml:"hosts,omitempty"`
	Tasks []TaskConfig `yaml:"tasks,omitempty"`
}

type Update struct {
	Updates []UpdateConfigs
}
