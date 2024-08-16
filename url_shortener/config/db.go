package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type gormPgConfig struct {
	host, username, password, name string
	port                           int
	ssl                            bool
	timezone                       *time.Location
}

func (c gormPgConfig) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.host,
		c.username,
		c.password,
		c.name,
		c.port,
		func() string {
			if c.ssl {
				return "enable"
			}
			return "disable"
		}(),
		c.timezone,
	)
}

func lookupString(name string, defaultValue string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}
	return defaultValue
}

func lookupInt(name string, defaultValue int) int {
	value := defaultValue
	if val, ok := os.LookupEnv(name); ok {
		if val_, err := strconv.Atoi(val); err != nil {
			value = val_
		}
	}

	return value
}

// func lookupFloat(name string, defaultValue float64) float64 {
// 	value := defaultValue
// 	if val, ok := os.LookupEnv(name); ok {
// 		if val_, err := strconv.ParseFloat(val, 64); err != nil {
// 			value = val_
// 		}
// 	}

// 	return value
// }

func lookupBool(name string, defaultValue bool) bool {
	value := defaultValue
	if val, ok := os.LookupEnv(name); ok {
		if val_, err := strconv.ParseBool(val); err != nil {
			value = val_
		}
	}

	return value
}

func GetDbConfig() gormPgConfig {
	return gormPgConfig{
		host:     lookupString("POSTGRES_HOST", "localhost"),
		username: lookupString("POSTGRES_USER", ""),
		password: lookupString("POSTGRES_PASS", ""),
		name:     lookupString("POSTGRES_DB", ""),
		port:     lookupInt("POSTGRES_PORT", 5432),
		ssl:      lookupBool("POSTGRES_SSL", false),
		timezone: func() *time.Location {
			if val, ok := os.LookupEnv("TZ"); ok {
				if tz, err := time.LoadLocation(val); err != nil {
					return tz
				}
			}
			return time.UTC
		}(),
	}
}
