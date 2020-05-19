package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	// plug in plugins here, for example:
	_ "github.com/devwak/caddyfilebroswer"
)

var run = caddymain.Run // replaced for tests

func main() {
	run()
}
