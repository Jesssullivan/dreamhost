module github.com/caddy-dns/dreamhost

go 1.16

require (
	github.com/caddyserver/caddy/v2 v2.10.0
	github.com/libdns/dreamhost v0.1.1
)

replace github.com/libdns/dreamhost => ../libdns-dreamhost 
