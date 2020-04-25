package storage

import "github.com/mkawserm/flamed/pkg/iface"

type Storage struct {
	configuration iface.IConfiguration
}

func (s *Storage) SetConfiguration(configuration iface.IConfiguration) {
	s.configuration = configuration
}

func (s *Storage) IsConfigurationValid() bool {

	return false
}
