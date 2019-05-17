package config

import (
	"errors"
	"strings"

	"github.com/lsgrep/gostatus/addon"
	"github.com/lsgrep/gostatus/utils"
	"github.com/spf13/viper"
)

var logger = utils.NewLogger()

type barConfig struct {
	Addons []map[string]interface{} `json:"addons"`
}

func ReadConfig(configPath string) ([]*addon.Addon, error) {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var cfg barConfig
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	var addons []*addon.Addon
	for _, m := range cfg.Addons {
		name, ok := m["name"].(string)
		if !ok || strings.TrimSpace(name) == "" {
			return nil, errors.New("Invalid Addon Name")
		}

		name = strings.ToLower(name)

		if name == "pomodoro" {
			addons = append(addons, addon.NewPomodoroAddon())
			continue
		}

		if name == "volume" {
			addons = append(addons, addon.NewVolumeAddon())
			continue
		}

		if name == "cpu" {
			addons = append(addons, addon.NewCPUAddon())
			continue
		}

		if name == "time" {
			addons = append(addons, addon.NewTimeAddon())
			continue
		}

		if name == "github" {
			username, ok := m["username"].(string)
			if !ok || strings.TrimSpace(username) == "" {
				return nil, errors.New("Invalid Github Username")
			}
			addons = append(addons, addon.NewGithubNotificationsAddon(username))
			continue
		}

		if name == "network" {
			iface, ok := m["interface"].(string)
			if !ok || strings.TrimSpace(iface) == "" {
				return nil, errors.New("Invalid Network Interface")
			}
			addons = append(addons, addon.NewNetworkAddon(iface))
			continue
		}

		if name == "ip" {
			iface, ok := m["interface"].(string)
			if !ok || strings.TrimSpace(iface) == "" {
				return nil, errors.New("Invalid Network Interface")
			}
			addons = append(addons, addon.NewIpAddon(iface))
			continue
		}

		if name == "ip_ext" {
			addons = append(addons, addon.NewIpExtAddon())
			continue
		}

		if name == "ipv6_ext" {
			logger.Error("ipv6....")
			addons = append(addons, addon.NewIpv6ExtAddon())
			continue
		}

		if name == "disk" {
			path, ok := m["path"].(string)
			if !ok || strings.TrimSpace(path) == "" {
				return nil, errors.New("Invalid Mount Path")
			}
			addons = append(addons, addon.NewDiskAddon(path))
			continue
		}

		if name == "ping" {
			addr, ok := m["address"].(string)
			if !ok || strings.TrimSpace(addr) == "" {
				return nil, errors.New("Invalid Address")
			}

			proto, ok := m["proto"].(string)
			if !ok {
				proto = "ipv4"
			} else if strings.TrimSpace(proto) == "" || proto != "ipv4" && proto != "ipv6" {
				return nil, errors.New("Invalid Proto")
			}
			addons = append(addons, addon.NewPinger(addr, proto))
			continue
		}
	}

	return addons, nil
}
