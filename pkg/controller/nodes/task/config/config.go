package config

import (
	"time"

	"github.com/lyft/flytestdlib/config"
	"k8s.io/apimachinery/pkg/util/sets"
)

//go:generate pflags Config --default-var defaultConfig

const SectionKey = "tasks"

var (
	defaultConfig = &Config{
		TaskPlugins: TaskPluginConfig{EnabledPlugins: []string{
			"container",
		}},
		MaxPluginPhaseVersions: 1000,
		BarrierConfig: BarrierConfig{
			Enabled:   true,
			CacheSize: 5000,
			CacheTTL:  config.Duration{Duration: time.Minute * 10},
		},
	}

	section = config.MustRegisterSection(SectionKey, defaultConfig)
)

type Config struct {
	TaskPlugins            TaskPluginConfig `json:"task-plugins" pflag:",Task plugin configuration"`
	MaxPluginPhaseVersions int32            `json:"max-plugin-phase-versions" pflag:",Maximum number of plugin phase versions allowed for one phase."`
	BarrierConfig          BarrierConfig    `json:"barrier" pflag:",Config for Barrier implementation"`
}

type BarrierConfig struct {
	Enabled   bool            `json:"enabled" pflag:",Enable Barrier transitions using inmemory context"`
	CacheSize int             `json:"cache-size" pflag:",Max number of barrier to preserve in memory"`
	CacheTTL  config.Duration `json:"cache-ttl" pflag:", Max duration that a barrier would be respected if the process is not restarted. This should account for time required to store the record into persistent storage (across multiple rounds."`
}

type TaskPluginConfig struct {
	EnabledPlugins []string
}

func (p TaskPluginConfig) GetEnabledPluginsSet() sets.String {
	return sets.NewString(p.EnabledPlugins...)
}

func GetConfig() *Config {
	return section.GetConfig().(*Config)
}
