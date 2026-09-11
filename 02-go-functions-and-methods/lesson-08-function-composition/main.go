package main

import (
	"fmt"
	"strings"
)

func trimUser(username string) string {
	return strings.TrimSpace(username)
}

func isLongEnough(username string) bool {
	return len(username) >= 4
}

func containsNoSpace(username string) bool {
	return !strings.ContainsAny(username, " \t\n\r")
}

func usernameExists(usernames []string, username string) bool {
	for _, candidate := range usernames {
		if candidate == username {
			return true
		}
	}
	return false
}

func validUsername(usernames []string, raw string) bool {
	username := trimUser(raw)
	return isLongEnough(username) && containsNoSpace(username) && !usernameExists(usernames, username)
}

func main() {
	usernames := []string{"adam", "sara"}
	candidate := "  med  "
	fmt.Println(validUsername(usernames, candidate))
}
