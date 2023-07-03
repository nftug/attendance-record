package repository

import (
	"attendance-record/domain/config"
	"attendance-record/domain/interfaces"
	"attendance-record/infrastructure/localpath"
	"encoding/json"
	"log"
	"os"
)

type configRepository struct {
	filename string
}

const ConfigFile = "config.json"

func NewConfigRepository(lp *localpath.LocalPathService) interfaces.IConfigRepository {
	return &configRepository{filename: lp.GetJoinedPath(ConfigFile)}
}

func (r *configRepository) LoadConfig() (*config.Config, error) {
	if err := r.initConfig(); err != nil {
		return nil, err
	}

	f, err := os.Open(r.filename)
	if err != nil {
		log.Fatal("failed to load config.json")
		return nil, err
	}

	defer f.Close()

	var cfg config.Config
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

func (r *configRepository) SaveConfig(cfg config.Config) error {
	f, err := os.Create(r.filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(cfg)
}

func (r *configRepository) initConfig() error {
	if _, err := os.Stat(r.filename); err == nil {
		return nil
	}

	f, err := os.Create(r.filename)
	if err != nil {
		log.Fatal("failed to create config.json")
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(config.DefaultConfig)
}
