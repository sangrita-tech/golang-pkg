package kube

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
)

func New(cfg Config) (*kubernetes.Clientset, error) {
	restCfg, err := buildRESTConfig(cfg)
	if err != nil {
		return nil, err
	}

	if cfg.QPS > 0 {
		restCfg.QPS = cfg.QPS
	}
	if cfg.Burst > 0 {
		restCfg.Burst = cfg.Burst
	}
	if cfg.Timeout > 0 {
		restCfg.Timeout = cfg.Timeout
	}

	cs, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		return nil, fmt.Errorf("kube: create clientset: %w", err)
	}
	return cs, nil
}
