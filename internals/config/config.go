package config

import "time"

// Config holds all application configuration
type Config struct {
	ClientAddr string       `yaml:"client"`
	ServerAddr string       `yaml:"server"`
	Timing     TimingConfig `yaml:"timing"`
	Protocol   string       `yaml:"protocol"` // this will be the starting protocol
}

type TimingConfig struct {
	Delay  Duration `yaml:"delay"`  // Base delay between cycles
	Jitter int      `yaml:"jitter"` // Jitter percentage (0-100)}
}

// Duration wraps time.Duration for YAML parsing
type Duration struct {
	time.Duration
}

// UnmarshalYAML implements yaml.Unmarshaler
func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	duration, err := time.ParseDuration(s)
	if err != nil {
		return err
	}

	d.Duration = duration
	return nil
}
