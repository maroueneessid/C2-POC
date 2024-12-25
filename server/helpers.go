package main

import "fmt"

// Anon function to check if port already listening
func KillListenerHelper(c map[int]chan bool, target int) {

	for port, sigChan := range c {
		if port == target {
			fmt.Printf("[!] Port equals")
			sigChan <- false
			close(sigChan)
			delete(c, port)
		}
	}
}

func ListenerPresence(c map[int]chan bool, target int) bool {
	_, exists := c[target]
	return exists
}
