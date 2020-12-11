package lexoffice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	client := NewClient(getApiKey())
	client.debug = true

	profile, err := client.GetProfile()
	assert.NoError(t, err, "error calling GetContact")
	assert.NotNil(t, profile, "no response")
	if client.debug {
		// fmt.Printf("%#v\n", contact)
		fmt.Println(prettyPrintJson(profile))
	}
}
