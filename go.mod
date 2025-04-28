module github.com/caddy-dns/dreamhost

go 1.16g

require (
	github.com/caddyserver/caddy/v2 v2.10.1
	github.com/libdns/dreamhost v0.1.1
)

replace github.com/libdns/dreamhost => github.com/Jesssullivan/libdns-dreamhost v0.2.0
