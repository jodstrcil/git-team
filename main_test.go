package main

import (
	"github.com/stretchr/testify/assert"
	"os"
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
			got, _ := format(tt.commitMessageContent)
			assert.Equal(t, tt.expectedFormattedMessage, got)
		})
	}
}

func TestLoadDefaultConfigWhenEnvironmentalIsEmpty(t *testing.T) {
	os.Clearenv()
	expectedPath := "./config.yml"
	result := getConfigPath()
	assert.Equal(t, expectedPath, result)
}

func TestLoadConfigDefinedInEnvironmental(t *testing.T) {
	envVarPath := "/some/path/config.yml"
	os.Setenv("GIT_TEAM_CONFIG_PATH", envVarPath)
	expectedPath := envVarPath
	result := getConfigPath()
	assert.Equal(t, expectedPath, result)
}