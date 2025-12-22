package kube

import "time"

type Config struct {
	KubeConfigPath string        `yaml:"kubeConfigPath" env:"KUBE_CONFIG_PATH"`
	KubeContext    string        `yaml:"kubeContext" env:"KUBE_CONTEXT"`
	QPS            float32       `yaml:"qps" env:"QPS" env-default:"20"`
	Burst          int           `yaml:"burst" env:"BURST" env-default:"40"`
	Timeout        time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"30s"`
}
