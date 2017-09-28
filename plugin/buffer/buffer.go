package buffer

import "github.com/mholt/caddy"

func init() {
	caddy.RegisterPlugin("buffer", caddy.Plugin{
		ServerType: "dns",
		Action:     setupBuffer,
	})
}
