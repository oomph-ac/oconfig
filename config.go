package oconfig

type Config struct {
	AuthKey [64]byte
	Branch  string

	LocalAddress  string
	RemoteAddress string
}

var DefaultConfig = Config{
	AuthKey: [64]byte{},
	Branch:  "stable",

	LocalAddress:  ":19132",
	RemoteAddress: ":20000",
}
