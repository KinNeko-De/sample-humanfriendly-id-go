package humanfriendly

import (
	"crypto/rand"
	"strings"
)

const charset = "ABCDEFGHJKLMNPQRSTUVWXYZ123456789"

// HumanFriendlyId is a struct that represents a human-friendly identifier.
type HumanFriendlyId struct {
	// Id is the human-friendly identifier in the form for internal use in the system.
	Id string
	// DisplayId is the human-friendly identifier in the form for display to users.
	DisplayId string
}

func ParseHumanFriendlyId(userInput string) (HumanFriendlyId, error) {
	normalizedId := strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(userInput, " ", ""), "I", "1"))
	return HumanFriendlyId{
		Id:        normalizedId,
		DisplayId: userInput,
	}, nil
}

func NewHumanFriendlyId(length int) (HumanFriendlyId, error) {
	idChars := make([]byte, length)
	displayLength := length + (length-1)/4
	displayIdChars := make([]byte, displayLength)
	_, err := rand.Read(idChars)
	if err != nil {
		return HumanFriendlyId{}, err
	}
	d := 0
	for k, v := range idChars {
		if k != 0 && k%4 == 0 {
			displayIdChars[d] = ' '
			d++
		}
		idChars[k] = charset[v%byte(len(charset))]
		displayIdChars[d] = idChars[k]
		d++
	}

	return HumanFriendlyId{
		Id:        string(idChars),
		DisplayId: string(displayIdChars),
	}, nil
}

func ToDisplayId(s string) {

}
