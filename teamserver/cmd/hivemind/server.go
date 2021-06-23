package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/conf"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/crud"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/listeners/simplehttp"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/listeners/tcp"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/plugins"
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/rest"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("missing argument: need path to config file")
	// 	os.Exit(1)
	// }
	// configFilePath := os.Args[1]
	configFilePath := "../../config/config.yaml"

	// Reads the config file
	var configOptions conf.ConfOptions
	configOptions.GetConf(configFilePath)

	// Sets the database parameters
	m := configOptions.Database
	if m == nil {
		fmt.Println("there was a problem connecting to the database")
		os.Exit(1)
	}
	crud.SetDatabaseOptions(m["uri"], m["port"], m["dbuser"], m["password"], m["sslmode"])

	var db crud.DatabaseModel
	db.Open()

	// Adds the configured users
	for _, val := range configOptions.Users {
		permission, err := strconv.Atoi(val["permission"])
		if err != nil {
			fmt.Println(err)
		}
		operator := model.Operator{
			Username:   val["username"],
			Password:   val["password"],
			Permission: permission,
		}
		db.InsertOperator(operator)
	}
	db.Close()

	// Start the different listeners
	for _, val := range configOptions.Listeners {
		for k, v := range val {
			switch k {
			case "tcp":
				if v["enabled"] == "true" {
					go tcp.StartListener(v["port"])
				}
			case "simplehttp":
				if v["enabled"] == "true" {
					go simplehttp.StartListener(v["port"])
				}
			default:
				break
			}
		}
	}

	// Starts any plugins
	for _, val := range configOptions.Plugins {
		for k, v := range val {
			switch k {
			case "pwnboard":
				if v["enabled"] == "true" {
					plugins.SetPwnboardConfig(true, v["url"], v["port"])
				}
			default:
				break
			}
		}
	}

	rest.Start(configOptions.Restapi["port"])
}
