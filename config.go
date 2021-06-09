package main

import (
	"os"

	t "milkteapool.com/pool-server/types"

	"gopkg.in/yaml.v2"
)

func NewConfig(configPath string) (*t.Config, error) {
	config := &t.Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
