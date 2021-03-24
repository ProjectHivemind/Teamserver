package main

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/listeners/tcp"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/rest"
)

func main() {
	go tcp.StartListener("1234")
	rest.Start("4321")
}
