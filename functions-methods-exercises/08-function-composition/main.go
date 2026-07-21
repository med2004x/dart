package main

func trimUser(username []string) {
	trimmedUsername := username
	for i := range trimmedUsername {
		if username[i] == "" {
			trimmedUsername = append(username[:i], username[:i+1]...)
		}
	}
}

func isLongEnough(username []string) bool {
	if len(username) != 0 {
		return true
	}
	return false
}

func main() {

}
