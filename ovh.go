package ovh

import (
	"github.com/libdns/ovh"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *ovh.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "dns.providers.ovh",
		New: func() caddy.Module {
			return &Provider{new(ovh.Provider)}
		},
	}
}

// Provision implements the Provisioner interface to initialize the OVH client
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.Endpoint = repl.ReplaceAll(p.Provider.Endpoint, "")
	p.Provider.ApplicationKey = repl.ReplaceAll(p.Provider.ApplicationKey, "")
	p.Provider.ApplicationSecret = repl.ReplaceAll(p.Provider.ApplicationSecret, "")
	p.Provider.ConsumerKey = repl.ReplaceAll(p.Provider.ConsumerKey, "")

	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// ovh {
//     endpoint <string>
//     application_key <string>
//     application_secret <string>
//     consumer_key <string>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "endpoint":
				if d.NextArg() {
					p.Provider.Endpoint = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "application_key":
				if d.NextArg() {
					p.Provider.ApplicationKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "application_secret":
				if d.NextArg() {
					p.Provider.ApplicationSecret = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "consumer_key":
				if d.NextArg() {
					p.Provider.ConsumerKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}

	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)