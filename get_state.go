package main

// GetState : Recupère l'état courant
func GetState(value string) string {
	switch value {
	case "un":
		return "1"
	case "deux":
		return "2"
	}
	return "0"
}
