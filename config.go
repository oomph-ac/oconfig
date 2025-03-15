package oconfig

const (
	ConfigVersion = "1.0-beta"

	DefaultShutdownMessage = "<red>Server is restarting.</red>"
)

type Config struct {
	Version string `json:"version"`
	AuthKey string `json:"auth_key"`
	Branch  string `json:"branch"`

	ShutdownMessage string `json:"shutdown_message"`

	LocalAddress  string `json:"local_addr"`
	RemoteAddress string `json:"remote_addr"`
	BackupAddress string `json:"backup_addr"`
	SpectrumKey   string `json:"spectrum_key"`

	Resource ResourceOpts `json:"resource_opts"`
	Movement MovementOpts `json:"movement_opts"`
	Mem      MemOpts      `json:"memory_opts"`
}

type ResourceOpts struct {
	ResourceFolder string `json:"resource_folder"`
	RequirePacks   bool   `json:"require_packs"`
}

type MovementOpts struct {
	CorrectionThreshold float64 `json:"correction_threshold"`
	PersuasionThreshold float64 `json:"persuasion_threshold"`
}

type MemOpts struct {
	GCPercent    int `json:"gc_percent"`
	MemThreshold int `json:"mem_threshold"`
}

var DefaultConfig = Config{
	Version: ConfigVersion,

	AuthKey: "your_auth_key_here",
	Branch:  "stable",

	LocalAddress:  ":19132",
	RemoteAddress: ":20000",

	ShutdownMessage: DefaultShutdownMessage,

	Resource: ResourceOpts{
		ResourceFolder: "resources/",
		RequirePacks:   true,
	},

	Movement: MovementOpts{
		CorrectionThreshold: 0.3,
		PersuasionThreshold: 0.002,
	},

	Mem: MemOpts{
		GCPercent:    -1,
		MemThreshold: 1000 * 1000 * 1000,
	},
}
