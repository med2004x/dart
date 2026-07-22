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
	if len(username) >= 4 {
		return true
	}
	return false
}
func containNoSpace(username []string) bool {
	for _,letter:= range username {
		if letter == ""{
			return false
		}
	}
	return true
}
func userNameExist(usernames []string, username string ) bool {
	for _,compare := range usernames {
		if compare == username {
			return true
		}
	}
	return false
}
func combine() {
	if isLongEnough() == true && containNoSpace() == true && userNameExist() == true {
		
	}
}

func main() {

}
