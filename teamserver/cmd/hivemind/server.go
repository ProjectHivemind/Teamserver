package main

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/conf"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/listeners/tcp"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/rest"
)

func main() {
	// Reads the config file
	var configOptions conf.ConfOptions
	configOptions.GetConf("../../config/config.yaml")

	// Sets the database parameters
	m := configOptions.Database
	crud.SetDatabaseOptions(m["uri"], m["port"], m["dbuser"], m["password"], m["sslmode"])

	// Start the different listeners
	for _, val := range configOptions.Listeners {
		for k, v := range val {
			switch k {
			case "tcp":
				if v["enabled"] == "true" {
					go tcp.StartListener(v["port"])
				}
			default:
				break
			}
		}
	}

	rest.Start(configOptions.Restapi["port"])
}
