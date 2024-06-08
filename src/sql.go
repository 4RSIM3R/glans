package src

func IsReadQuery(query string) bool {
	return len(query) >= 6 && query[:6] == "SELECT"
}
