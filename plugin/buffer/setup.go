package buffer

import (
	"strconv"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"

	"github.com/mholt/caddy"
)

func setupBuffer(c *caddy.Controller) error {
	config := dnsserver.GetConfig(c)
	for c.Next() {
		args := c.RemainingArgs()
		if len(args) == 0 || len(args) > 2 {
			return plugin.Error("buffer", c.ArgErr())
		}

		size, err := strconv.Atoi(args[0])
		if err != nil {
			return plugin.Error("buffer", err)
		}

		config.UDPRxSize = size
	}
	return nil
}
