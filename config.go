package oconfig

import (
	"strings"
)

const (
	ConfigVersion = "0.1-beta"

	LoggingTypeGlobal    = "global"
	LoggingTypePerPlayer = "per_player"

	DefaultShutdownMessage = "<red>Server is restarting.</red>"
)

type Config struct {
	Version string `json:"version"`

	AuthKey string `json:"auth_key"`
	Branch  string `json:"branch"`

	LocalAddress  string `json:"local_addr"`
	RemoteAddress string `json:"remote_addr"`

	LoggingType string `json:"logging_type"`
	LogFile     string `json:"log_file"`

	ShutdownMessage string `json:"shutdown_message"`

	Resource ResourceOpts `json:"resource_opts"`

	Movement MovementOpts `json:"movement_opts"`

	Moderators []string `json:"moderators"`
}

type ResourceOpts struct {
	ResourceFolder   string `json:"resource_folder"`
	RequirePacks     bool   `json:"require_packs"`
	FetchPacksRemote bool   `json:"fetch_packs_remote"`
}

type MovementOpts struct {
	CorrectionThreshold float64 `json:"correction_threshold"`
}

var DefaultConfig = Config{
	Version: ConfigVersion,

	AuthKey: strings.Repeat("0", 64),
	Branch:  "stable",

	LocalAddress:  ":19132",
	RemoteAddress: ":20000",

	LoggingType: LoggingTypePerPlayer,
	LogFile:     "oomph.log",

	ShutdownMessage: DefaultShutdownMessage,

	Resource: ResourceOpts{
		ResourceFolder:   "resources/",
		RequirePacks:     true,
		FetchPacksRemote: true,
	},

	Movement: MovementOpts{
		CorrectionThreshold: 0.3,
	},

	Moderators: []string{},
}
