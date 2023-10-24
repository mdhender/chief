// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"encoding/json"
	"os"
)

func LoadStarting(path string) (*Starting, error) {
	s := &Starting{}
	return s, s.Load(path)
}

func (s *Starting) Load(path string) error {
	s.Lock()
	defer s.Unlock()
	if data, err := os.ReadFile(path); err != nil {
		return err
	} else if err = json.Unmarshal(data, s); err != nil {
		return err
	}
	return nil
}
