package auth

func extractSession() string {
	return "Loggedin"
}

func GetSession() string {
	return extractSession()
}
