package repository

import (
	"attendance-record/domain/config"
	"encoding/json"
	"log"
	"os"
)

type configRepository struct{}

const ConfigFile = "config.json"

func NewConfigRepository() config.IConfigRepository {
	return &configRepository{}
}

func (r *configRepository) LoadConfig() (*config.Config, error) {
	f, err := r.openFile()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg config.Config
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

func (r *configRepository) SaveConfig(cfg config.Config) error {
	f, err := r.openFile()
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(cfg)
}

func (r *configRepository) openFile() (*os.File, error) {
	if err := r.initConfig(); err != nil {
		return nil, err
	}

	f, err := os.Open(ConfigFile)
	if err != nil {
		log.Fatal("failed to load config.json")
		return nil, err
	}

	return f, err
}

func (r *configRepository) initConfig() error {
	if _, err := os.Stat(ConfigFile); err == nil {
		return nil
	}

	f, er := os.Create(ConfigFile)
	if er != nil {
		log.Fatal("failed to create config.json")
		return er
	}
	defer f.Close()

	cfg := config.Config{WorkHours: 8}
	return json.NewEncoder(f).Encode(cfg)
}
