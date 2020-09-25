package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommitMessageFormat(t *testing.T) {
	pairOne := User{"gthreepwood",
		"Guybrush Threepwood",
		"gthreepwood@monkeyisland.com"}

	pairTwo := User{"emarley",
		"Elaine Marley",
		"emarley@monkeyisland.com"}

	pairs := []User{}

	tests := []struct {
		name                     string
		commitMessageContent     MessageContent
		expectedFormattedMessage string
	}{
		{
			"Plain commit message",
			MessageContent{Message: "Plain commit message"},
			"Plain commit message\n"},
		{
			"Plain commit message with ticket number",
			MessageContent{Message: "Plain commit message", JiraTag: "MYTEAM", JiraNumber: 1234},
			"[MYTEAM-1234] Plain commit message\n"},
		{
			"Plain commit message with ticket number one pair",
			MessageContent{
				Collaborators: append(pairs, pairOne),
				Message:       "Plain commit message",
				JiraTag:       "MYTEAM", JiraNumber: 1234},
			"[MYTEAM-1234] Plain commit message\n\n\n" +
				"Co-Authored-by: Guybrush Threepwood <gthreepwood@monkeyisland.com>\n"},
		{
			"Plain commit message with ticket number multiple pairs pair",
			MessageContent{
				Collaborators: append(pairs, pairOne, pairTwo),
				Message:       "Plain commit message",
				JiraTag:       "MYTEAM", JiraNumber: 1234},
			"[MYTEAM-1234] Plain commit message\n\n\n" +
				"Co-Authored-by: Guybrush Threepwood <gthreepwood@monkeyisland.com>\n" +
				"Co-Authored-by: Elaine Marley <emarley@monkeyisland.com>\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := format(tt.commitMessageContent)
			assert.Equal(t, tt.expectedFormattedMessage, got)
		})
	}
}
