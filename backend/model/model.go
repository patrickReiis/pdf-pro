package model

// Returns the user by its Api Key
// If the user does not exists returns an empty string
func GetUserByApiKey(apiKey string) (user string) {
	return getUserByApiKeyImpl(apiKey)
}

func getUserByApiKeyImpl(apiKey string) string {
	return "implement"
}
