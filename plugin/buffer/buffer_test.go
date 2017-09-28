package buffer

import (
	"testing"

	"github.com/coredns/coredns/core/dnsserver"

	"github.com/mholt/caddy"
)

func TestSetupBuffers(t *testing.T) {
	c := caddy.NewTestController("dns", `buffer 1`)
	err := setupBuffers(c)
	if err != nil {
		t.Fatalf("Expected no errors, but got: %v", err)
	}

	cfg := dnsserver.GetConfig(c)
	if got, want := cfg.UDPRxBuffer, 1; got != want {
		t.Errorf("Expected the config's UDPRxBuffer to be %d, was %d", want, got)
	}
}

func TestBuffers(t *testing.T) {
	c := caddy.NewTestController("dns", `buffer`)
	err := setupBuffers(c)
	if err == nil {
		t.Fatalf("Expected errors, but got none")
	}
}
