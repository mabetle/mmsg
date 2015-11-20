package mmsg

var userLocale = make(map[string]string)

func PutUserLocale(username, locale string) {
	userLocale[username] = locale
}

func GetUserLocale(username string) string {
	if locale, ok := userLocale[username]; ok {
		return locale
	}
	return DefaultLocale
}

// UserMessage
func UserMessage(username string, key string, args ...interface{}) string {
	return Message(GetUserLocale(username), key, args...)
}
