package config

type DiscordNotifier struct{
	ID string `yaml:"webhook_id,omitempty"`
	Token string `yaml:"webhook_token,omitempty"`
	Content string `yaml:"content,omitempty"`
}
type DebugConfig struct {
	Msg string `yaml:"msg,omitempty"`
}
type Fact struct {
	IsCentOS string `yaml:"is_centos,omitempty"`
	IsUbuntu string `yaml:"is_ubuntu,omitempty"`
}

type TaskConfig struct {
	Name string `yaml:"name,omitempty"`
	Command string `yaml:"command,omitempty"`
	Register string `yaml:"register,omitempty"`
	When string `yaml:"when,omitempty"`
	Set_fact []Fact `yaml:"set_fact,omitempty"`
	IgnoreErrors bool `yaml:"ignore_errors,omitempty"`
	Debug []DebugConfig `yaml:"debug,omitempty"`
	DiscordMessage []DiscordNotifier `yaml:"community.general.discord,omitempty"`
}

// Playbook structs
type UpdateConfigs struct {
	Name  string `yaml:"name,omitempty"`
	Hosts string `yaml:"hosts,omitempty"`
	Tasks []TaskConfig `yaml:"tasks,omitempty"`
}

type Update struct {
	Updates []UpdateConfigs
}
