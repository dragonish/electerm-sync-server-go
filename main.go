// Project entry program.
package main

import (
	"electerm/internal/daemon"
	"electerm/internal/parse"
)

func main() {
	parse.Parse()
	daemon.StartDaemon()
}
