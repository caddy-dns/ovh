# OVH module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with OVH accounts.

## Caddy module name

```
dns.providers.ovh
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "ovh",
        "endpoint": "{env.OVH_ENDPOINT}",
        "application_key": "{env.OVH_APPLICATION_KEY}",
        "application_secret": "{env.OVH_APPLICATION_SECRET}",
        "consumer_key": "{env.OVH_CONSUMER_KEY}",
      }
    }
  }
}
```

or with the Caddyfile:

```
tls {
  dns ovh {
    endpoint {$OVH_ENDPOINT}
    application_key {$OVH_APPLICATION_KEY}
    application_secret {$OVH_APPLICATION_SECRET}
    consumer_key {$OVH_CONSUMER_KEY}
  }
}
```

You can replace `{$*}` or `{env.*}` with the actual values if you prefer to put it directly in your config instead of an environment variable.

## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/ovh#authenticating) for important information about credentials.
