package config

import (
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type Config struct {
	Latitude  float64
	Longitude float64
	Country   string
	Region    string
}

func defaultCfg() Config {
	return Config{
		Latitude:  0.0,
		Longitude: 0.0,
		Country:   "",
		Region:    "",
	}
}

func configPath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		return ""
	}

	path := filepath.Join(dir, "clawea")
	_ = os.MkdirAll(path, 0755)

	return filepath.Join(path, "config.conf")
}

func Load() Config {
	cfg := defaultCfg()
	path := configPath()

	if _, err := os.Stat(path); err != nil {
		return cfg
	}

	iniFile, err := ini.Load(path)
	if err != nil {
		return cfg
	}

	sec := iniFile.Section("location")
	cfg.Latitude = sec.Key("latitude").MustFloat64(cfg.Latitude)
	cfg.Longitude = sec.Key("longitude").MustFloat64(cfg.Longitude)
	cfg.Country = sec.Key("country").MustString(cfg.Country)
	cfg.Region = sec.Key("region").MustString(cfg.Region)

	return cfg
}

func EnsureConfig() {
	path := configPath()

	if _, err := os.Stat(path); err == nil {
		return
	}

	content := `# Latitude and longitude data must be in float64
# Country and region data must be in string (it will show on top of the weather stats)
# If latitude and longitude are 0.0 and longitude are 0.0, it will use ip-api.com to get the location

[location]
latitude  = 0.0
longitude = 0.0
country   = 
region    = 
`

	_ = os.WriteFile(path, []byte(content), 0644)
}
