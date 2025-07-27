package oconfig

const (
	ConfigVersion          uint64 = 1
	DefaultShutdownMessage        = "§cServer is restarting."
)

type Config struct {
	SpectrumAPIToken string `json:"spectrum_api_token" comment:"The Spectrum API token used to authenticate with a Spectrum API instance.\nIf you do not know what this is - don't set anything here."`

	LocalAddress  string `json:"local_addr" comment:"The address the proxy listens on for incoming connections. For the most part, just using a colon followed by the port is fine."`
	RemoteAddress string `json:"remote_addr" comment:"The address the proxy first connects to for the remote server."`
	BackupAddress string `json:"backup_addr" comment:"The address the proxy will connect to if the inital connection fails to the remote address."`

	ShutdownMessage string `json:"shutdown_message" comment:"The message players are disconnected with when the proxy is shut down and there is no available reconnect address."`
	ReconnectIP     string `json:"reconnect_ip" comment:"The IP address players connected to the proxy are transferred to in the event of a shutdown.\nIf this option is empty, players will be disconnected instead."`
	ReconnectPort   int    `json:"reconnect_port" comment:"The port players connected to the proxy are transferred to in the event of a shutdown.\nIf this option is empty, players will be disconnected instead."`

	UseDebugCommands bool `json:"allow_debug_commands" comment:"This option signifies wether debug commands should be enabled on the proxy. If this is disabled, then\n any attempt to run Oomph debug commands (!oomph_debug) will not be handled."`

	Resource ResourceOpts `json:"resource_opts" comment:"Options for your resource packs."`
	Movement MovementOpts `json:"movement_opts" comment:"Options for configuring movement policies and strictness for Oomph."`
	Combat   CombatOpts   `json:"combat_opts" comment:"Options for configuring combat policies and strictness for Oomph."`
	Mem      MemOpts      `json:"memory_opts" comment:"Memory options to be used by the proxy process."`
}

type ResourceOpts struct {
	ResourceFolder string `json:"resource_folder" comment:"The folder where resource packs are stored. If your resource pack requires a content key, list them in a JSON file with a map of the pack UUID to the content key."`
	RequirePacks   bool   `json:"require_packs" comment:"Set to true if you require players to download the resource packs stored on the proxy."`
}

type MovementOpts struct {
	// The correction threshold represents the amount of distance in blocks between the client and Oomph's
	// predicted position is required to trigger a correction.
	CorrectionThreshold float32 `json:"correction_threshold" comment:"The amount of blocks between the client and Oomph's prediction required to trigger a correction. It is recommended to keep this option at around 0.2-0.5 blocks to avoid large noticable corrections."`
	// PersuasionThreshold is the amount of blocks per tick Oomph should move towards the client position (given that there are no pending corrections and the player's movement has been valid for a long enough period of time).
	// Note that this persuasion is not applied on the Y-axis.
	PersuasionThreshold float32 `json:"persuasion_threshold" comment:"The amount of block per tick Oomph's position moves towards the client's position. Note that the persuasion is not applied on the Y-axis.\nIncreasing this value above the default option will result in slight movement bypasses."`
	// AcceptClientPosition is a boolean that represents if the Oomph proxy should accept the client's position if
	// their position is within the opts.PositionAcceptanceThreshold and the player has no pending corrections (
	// and some other factors, such as immobile, etc.)
	// By default, this is disabled as it may result in small movement bypasses, but enabling it can help
	// reduce the amount of false corrections sent to the client.
	AcceptClientPosition bool `json:"accept_client_position" comment:"Should Oomph accept the client's position completely if it's within the PositionAcceptanceThreshold?\nEnabling this option will result in small movement bypasses."`
	// PositionAcceptanceThreshold (which is only used if AcceptClientPosition is TRUE) is the maximum allowed distance in blocks required for Oomph to accept the client's position.
	PositionAcceptanceThreshold float32 `json:"position_acception_threshold" comment:"The distance between the client and Oomph's position (in blocks) required for Oomph to accept the client's position. This option is only applied if AcceptClientPosition is set to true."`
	// AcceptClientVelocity is a boolean that represents if the Oomph proxy should accept the client's velocity in
	// PlayerAuthInputPacket if the difference between the client's and server predicted end-frame velocity in blocks is
	// within the opts.VelocityAcceptanceThreshold and the player has no pending corrections (and some other factors, such
	// as immobile, etc.)
	// By default, this is disabled as it may result in small movement bypasses, but enabling it can help
	// reduce the amount of false corrections sent to the client.
	AcceptClientVelocity bool `json:"accept_client_velocity" comment:"Should Oomph accept the client's velocity completely if it's within the VelocityAcceptanceThreshold?\nMay result in small movement bypasses."`
	// VelocityAcceptanceThreshold (which is only used if AcceptClientVelocity is TRUE) is the maximum allowed difference
	// in blocks required for Oomph to accept the client's velocity. Note that this perusasion is not applied on the Y-axis.
	VelocityAcceptanceThreshold float32 `json:"velocity_acception_threshold" comment:"The difference between the client and Oomph's velocity (in blocks) required for Oomph to accept the client's velocity.\nNOTE: Increasing this above 0.07 may result in movement bypasses."`
}

type CombatOpts struct {
	// MaxRewind is the maximum amount of positions Oomph will store for each entity for combat rewind and simulation.
	MaxRewind int `json:"max_rewind" comment:"The maximum amount of positions Oomph should store for each entity for combat rewind and simulation.\nThis value is capped at 20 ticks (1000ms).\nThis option is not applied if FullAuthoritative is set to false."`
	// MaximumAttackAngle is the maximum angle in degrees that Oomph will allow for an attack to be considered valid.
	MaximumAttackAngle float32 `json:"maximum_attack_angle" comment:"The maximum angle in degrees that Oomph will allow for an attack to be considered valid."`
	// EnableClientEntityTracking is a boolean that indicates if the proxy should also enable it's client-sided entity tracking to perfectly lag compensate for the client view of entities. This is primarily used for
	// detecting and taking action against reach/killaura. This option is not neccessary for combat rewind to work properly, but should be enabled if you need precise information for herustics or whatnot.
	EnableClientEntityTracking bool `json:"enable_client_entity_tracking" comment:"This option is used to enable Oomph's client-sided entity tracking to perfectly lag compensate for the client view of entities. If you want to enable reach detections, this should be enabled."`
	// AllowNonMobileTouch is a boolean indicating if the proxy should allow non-mobile players to use the touch input mode. Although some devices on Windows may allow for the touch input mode,
	// it can allow for tiny combat gains when specifically being abused.
	AllowNonMobileTouch bool `json:"allow_non_mobile_touch" comment:"Should Oomph allow non-mobile players to use the touch input mode? It is recommended to keep this option disabled as it may allow for\ntiny combat gains from players spoofing their input mode to touch."`
	// AllowSwitchInputMode is a boolean indicating if the proxy should allow players to switch their input mode, or if it should enforce one input mode to be used unless the player re-joins the server. This can primarily
	// be used to prevent players from using exploits that switch the input mode to touch only when enabled to obtain a combat advantage. This option is recommended to be set to false.
	AllowSwitchInputMode bool `json:"allow_switch_input_mode" comment:"Should Oomph allow players to switch their input mode?\nIt is recommended to keep this option disabled as it may allow for tiny combat gains from players switching their input mode to touch."`
}

type MemOpts struct {
	GCPercent    int `json:"gc_percent" comment:"Golang's garbage collection percentage. If set to -1 (the default value), the proxy will only run garbage collection when reaching the memory soft limit.\nWe recommend NOT changing this value unless you know what you're doing."`
	MemThreshold int `json:"mem_threshold" comment:"A soft-limit for how much memory the Oomph proxy should use in megabytes. The default value is 1GB (1024MB).\nIf you are running Oomph on a container, we recommend setting this to roughly ~500MB lower to avoid OOM errors.\nIncrease this as neccessary to reduce garbage collection cycles."`
}

var (
	Cfg Config

	DefaultConfig = Config{
		SpectrumAPIToken: "api_token_here",

		LocalAddress:  ":19132",
		RemoteAddress: ":20000",

		ShutdownMessage: DefaultShutdownMessage,

		Resource: ResourceOpts{
			ResourceFolder: "resources/",
			RequirePacks:   true,
		},

		Movement: MovementOpts{
			CorrectionThreshold:         0.3,
			PersuasionThreshold:         0.002,
			AcceptClientPosition:        false,
			PositionAcceptanceThreshold: 0.09,
			AcceptClientVelocity:        false,
			VelocityAcceptanceThreshold: 0.03,
		},

		Combat: CombatOpts{
			MaxRewind:                  6,
			MaximumAttackAngle:         85.0,
			EnableClientEntityTracking: true,
		},

		Mem: MemOpts{
			GCPercent:    -1,
			MemThreshold: 1 * 1024 * 1024 * 1024,
		},
	}
)

func Resource() ResourceOpts {
	return Cfg.Resource
}

func Movement() MovementOpts {
	return Cfg.Movement
}

func Combat() CombatOpts {
	return Cfg.Combat
}

func Mem() MemOpts {
	return Cfg.Mem
}
