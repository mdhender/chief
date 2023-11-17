// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package config

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	Env    string
	Games  map[string]*Game `json:"games,omitempty"`
	Server Server           `json:"server,omitempty"`
}

type Game struct {
	Id    string           `json:"id"`
	Turn  GameTurn         `json:"turn,omitempty"`
	Clans map[string]*Clan `json:"clans"`
}

type GameTurn struct {
	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
}

type Clan struct {
	Id   string `json:"id"`
	Root string `json:"root,omitempty"`
	Docs string `json:"docs,omitempty"`
}

// Default returns a Config that has been initialized with
// default values.
func Default() *Config {
	cfg := Config{
		Env:    "unknown",
		Games:  make(map[string]*Game),
		Server: defaultServer(),
	}
	return &cfg
}

// Load updates the Config from JSON.
// It returns the first error encountered (opening, reading, or parsing).
// If there are errors opening the file, the Config is not updated.
// If there are errors reading or parsing, the Configis left in an unknown state.
func (c *Config) Load(path string) error {
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(fp).Decode(c); err != nil {
		return err
	}
	return nil
}

type Server struct {
	Host         string `json:"host,omitempty"`
	Port         string `json:"port,omitempty"`
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func defaultServer() Server {
	return Server{
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
}
