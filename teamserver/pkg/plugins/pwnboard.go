package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var ENABLED bool = false
var PWNBOARD_URL string = "http://pwnboard.win"
var PWNBOARD_PORT string = "80"

type PwnBoard struct {
	IPs  string `json:"ip"`
	Type string `json:"type"`
}

func SetPwnboardConfig(enabled bool, url, port string) {
	ENABLED = enabled
	PWNBOARD_URL = url
	PWNBOARD_PORT = port
}

func UpdatepwnBoard(ip, implantName string) {
	// Returns if it isn't enabled.
	if !ENABLED {
		return
	}

	// Create the struct
	data := PwnBoard{
		IPs:  "172.16.3.22",
		Type: implantName,
	}

	// Marshal the data
	sendit, err := json.Marshal(data)
	if err != nil {
		fmt.Println("\n[-] ERROR SENDING POST:", err)
		return
	}

	// Send the post to pwnboard
	resp, err := http.Post(PWNBOARD_URL+":"+PWNBOARD_PORT+"/generic", "application/json", bytes.NewBuffer(sendit))
	if err != nil {
		fmt.Println("[-] ERROR SENDING POST:", err)
		return
	}

	defer resp.Body.Close()
}
