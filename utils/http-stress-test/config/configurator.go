package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Target struct {
	URL string
}

type Configuration struct {
	Users    int       `json:"users"`    // Users to simulate or workers to run concurrently.
	Targets  []*Target `json:"targets"`  // Targets to address randomly
	Time     []int     `json:"time"`     // Timestamps where the RPS should be reached in s starting at 0.
	Requests []int     `json:"requests"` // Requests amount to send per second.
}

func GetConfig(path string) (*Configuration, error) {
	viper.SetConfigFile(path)
	viper.SetDefault("users", 10)
	viper.SetDefault("targets", []*Target{
		{URL: "https://google.de:443"},
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	configuration := &Configuration{}
	if err := viper.Unmarshal(configuration); err != nil {
		return nil, err
	}

	if len(configuration.Time) != len(configuration.Requests) {
		return nil, fmt.Errorf("length of 'time' and 'requests' should be equal")
	}

	return configuration, nil
}
