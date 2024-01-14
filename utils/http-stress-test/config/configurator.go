package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Target struct {
	URL string
}

type Phase struct {
	TargetRps int `json:"targetRps"` // TargetRps requests amount to send per second for this phase./
	TimeIdx   int `json:"timeIdx"`   // TimeIdx where the RPS should be reached in s starting at 0.
}

type Configuration struct {
	Users   int       `json:"users"`   // Users to simulate or workers to run concurrently.
	Targets []*Target `json:"targets"` // Targets to address randomly
	Phases  []*Phase  `json:"phases"`  // Phases which are configured for this test scenario.
}

func GetConfig(path string) (*Configuration, error) {
	viper.SetConfigFile(path)
	viper.SetDefault("users", 10)
	viper.SetDefault("targets", []*Target{
		{URL: "https://google.de:443"},
	})
	viper.SetDefault("phases", []*Phase{
		{
			TargetRps: 0,
			TimeIdx:   0,
		},
		{
			TargetRps: 500,
			TimeIdx:   10,
		},
		{
			TargetRps: 500,
			TimeIdx:   20,
		},
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	configuration := &Configuration{}
	if err := viper.Unmarshal(configuration); err != nil {
		return nil, err
	}

	if configuration.Phases == nil || len(configuration.Phases) == 0 {
		panic("No phases specified")
	}

	return configuration, nil
}
