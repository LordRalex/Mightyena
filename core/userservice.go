package core

var userCache = make(map[string]User)

func GetUser(nickname string) (User, error) {
	return userCache[nickname], nil
}