package vipconfig

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Adds --hetzner-cred-file flag and wires env + default.
// Priority remains: flag > env > yaml.
// Env key: VIP_HETZNER_CRED_FILE
// YAML key: hetzner-cred-file
func init() {
	// CLI flag (default keeps backward compatibility)
	pflag.String("hetzner-cred-file", "/etc/hetzner", "Path to Hetzner Robot credentials file (default: /etc/hetzner)")

	// Default + ENV binding
	viper.SetDefault("hetzner-cred-file", "/etc/hetzner")
	_ = viper.BindEnv("hetzner-cred-file", "VIP_HETZNER_CRED_FILE")

	// Bind the flag to viper (cobra/pflag parsing happens in main/config init)
	_ = viper.BindPFlag("hetzner-cred-file", pflag.Lookup("hetzner-cred-file"))
}