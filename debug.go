package jira_hooks

import (
	"log"
)

var DebugRequest = false

func debugRequest(payload []byte) {
	log.Println("=== BEGIN DEBUG REQUEST ===")
	log.Println(string(payload))
	log.Println("=== END DEBUG REQUEST ===")
}
