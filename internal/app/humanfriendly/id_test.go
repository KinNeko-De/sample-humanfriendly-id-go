package humanfriendly

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const allowedChars = "ABCDEFGHJKLMNPQRSTUVWXYZ123456789"

func TestNewHumanFriendlyId_Id_IsGenerated(t *testing.T) {
	tests := []struct {
		length int
	}{
		{1},
		{16},
	}

	for _, test := range tests {
		id, err := NewHumanFriendlyId(test.length)

		assert.Nil(t, err)
		assert.Equal(t, test.length, len(id.Id))
		assert.Regexp(t, regexp.MustCompile("^["+allowedChars+"]*$"), id.Id)
	}
}

func TestNewHumanFriendlyId_DisplayId_IsGroupedBySpaces(t *testing.T) {
	tests := []struct {
		length                  int
		expectedDisplayIdLength int
	}{
		{1, 1},
		{4, 4},
		{5, 6},
		{8, 9},
		{9, 11},
		{12, 14},
		{13, 16},
	}

	for _, test := range tests {
		id, err := NewHumanFriendlyId(test.length)

		assert.Nil(t, err)
		assert.Regexp(t, regexp.MustCompile(createRegex(test.length)), id.DisplayId)
	}
}

func createRegex(length int) string {
	rest := length % 4
	expectedBlocks := length / 4

	regexTemplate := "[" + allowedChars + "]{4}"
	lastRegex := ""
	if rest > 0 {
		lastRegex = "[" + allowedChars + "]{" + strconv.Itoa(rest) + "}"
	}

	var regex strings.Builder
	regex.WriteString("^")

	for i := 0; i < expectedBlocks; i++ {
		if i != 0 {
			regex.WriteString(" ")
		}
		regex.WriteString(regexTemplate)
	}

	if expectedBlocks > 0 && rest > 0 {
		regex.WriteString(" ")
	}
	regex.WriteString(lastRegex + "$")

	return regex.String()
}

func TestParseHumanFriendlyId_Id_IsNormalized(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"ABCDEFGHJKLM", "ABCDEFGHJKLM"},
		{"abcdefghjklm", "ABCDEFGHJKLM"},
		{"I", "1"},
		{"i", "1"},
		{"ABCD EFGH JKLM", "ABCDEFGHJKLM"},
	}

	for _, test := range tests {
		id, err := ParseHumanFriendlyId(test.input)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, id.Id)
	}

}
