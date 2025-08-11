package oconfig

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

func Combat() CombatOpts {
	return Global.Combat
}
