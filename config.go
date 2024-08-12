package oconfig

import (
	"encoding/json"
	"strings"
)

const (
	ConfigVersion = "1.0.0-beta"

	LoggingTypeGlobal    = "global"
	LoggingTypePerPlayer = "per_player"

	LatencyReportTypeRaknet = "raknet"
	LatencyReportTypeStack  = "nsl"

	AuthorityNone = "none"
	AuthoritySemi = "semi-auth"
	AuthorityFull = "full-auth"
)

type Config struct {
	Version string `json:"version"`

	AuthKey string `json:"auth_key"`
	Branch  string `json:"branch"`

	LocalAddress   string `json:"local_addr"`
	RemoteAddress  string `json:"remote_addr"`
	Authentication bool   `json:"authentication"`

	LoggingType string `json:"logging_type"`
	LogFile     string `json:"log_file"`

	Resource ResourceOpts `json:"resource_opts"`

	Authority AuthorityOpts `json:"authority_opts"`
}

type ResourceOpts struct {
	ResourceFolder   string `json:"resource_folder"`
	RequirePacks     bool   `json:"require_packs"`
	FetchPacksRemote bool   `json:"fetch_packs_remote"`
}

type AuthorityOpts struct {
	CombatAuthority     string  `json:"combat_authority"`
	MovementAuthority   string  `json:"movement_authority"`
	CorrectionThreshold float64 `json:"correction_threshold"`
	ClientPredictsSpeed bool    `json:"client_predicts_speed"`
}

var DefaultConfig = Config{
	Version: ConfigVersion,

	AuthKey: strings.Repeat("0", 64),
	Branch:  "stable",

	LocalAddress:   ":19132",
	RemoteAddress:  ":20000",
	Authentication: true,

	LoggingType: LoggingTypePerPlayer,
	LogFile:     "oomph.log",

	Resource: ResourceOpts{
		ResourceFolder:   "resources/",
		RequirePacks:     true,
		FetchPacksRemote: true,
	},

	Authority: AuthorityOpts{
		CombatAuthority:   AuthoritySemi,
		MovementAuthority: AuthoritySemi,

		CorrectionThreshold: 1_000_000.0,
		ClientPredictsSpeed: false,
	},
}

var DefaultConfigMap = map[string]interface{}{}

func init() {
	enc, err := json.Marshal(DefaultConfig)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(enc, &DefaultConfigMap); err != nil {
		panic(err)
	}
}
