package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/validator.v2"

	"code.cloudfoundry.org/lager/v3"
)

type Netmon struct {
	PollInterval      int    `json:"poll_interval" validate:"min=1"`
	MetronAddress     string `json:"metron_address" validate:"nonzero"`
	InterfaceName     string `json:"interface_name" validate:"nonzero"`
	LogLevel          string `json:"log_level"`
	LogPrefix         string `json:"log_prefix" validate:"nonzero"`
	IPTablesLockFile  string `json:"iptables_lock_file" validate:"nonzero"`
	TelemetryEnabled  bool   `json:"telemetry_enabled"`
	TelemetryInterval int    `json:"telemetry_interval"`
}

func (n Netmon) ParseLogLevel() (lager.LogLevel, error) {
	switch strings.ToLower(n.LogLevel) {
	case "debug":
		return 0, nil
	case "info":
		return 1, nil
	case "error":
		return 2, nil
	case "fatal":
		return 3, nil
	}
	return 0, fmt.Errorf(`unknown log level %q`, n.LogLevel)
}

func (c *Netmon) Validate() error {
	if c.TelemetryEnabled && c.TelemetryInterval <= 0 {
		return errors.New("telemetry_interval must be set to a positive, non-zero value if telemetry_enabled is true")
	}
	return validator.Validate(c)
}

func New(path string) (*Netmon, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("file does not exist: %s", err)
	}
	jsonBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config: %s", err) // not tested
	}

	cfg := Netmon{}
	err = json.Unmarshal(jsonBytes, &cfg)
	if err != nil {
		return nil, fmt.Errorf("parsing config: %s", err)
	}

	if err := cfg.Validate(); err != nil {
		return &cfg, fmt.Errorf("invalid config: %s", err)
	}

	return &cfg, nil
}
