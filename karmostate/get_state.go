package karmostate

import "strings"

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

// GetWeightValue translate String code to Weight value
func GetWeightValue(value string) float64 {
	switch strings.ToLower(value) {
	case "small":
		return .5
	case "medium":
		return 1
	case "large":
		return 2
	case "half":
		return 4
	case "full":
		return 8
	}
	return 0
}
