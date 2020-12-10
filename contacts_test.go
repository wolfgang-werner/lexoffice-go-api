package lexoffice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_CreateContact(t *testing.T) {
	var note = "creator: lexoffice-go-api"
	var contact = Contact{
		// read only ID:             nil,
		// read only OrganizationID: nil,
		Version: 0,
		Roles: Roles{
			Customer: &Customer{},
		},
		Company: nil,
		// Company: nil,
		Person: &Person{
			LastName: "Müller",
		},
		Addresses: &Addresses{
			Billing: []Address{
				{
					Supplement:  "billing address",
					Street:      "Teststraße 123",
					Zip:         "20355",
					City:        "Hamburg",
					CountryCode: "DE",
				},
			},
			Shipping: nil,
		},
		EmailAddresses: &EmailAddresses{
			Business: nil,
			Office:   []string{"wwerner@outlook.com"},
			Private:  nil,
			Other:    nil,
		},
		/*
			PhoneNumbers: &PhoneNumbers{
				Business: []string{"+49 170 810 92 41"},
				Office:   nil,
				Mobile:   nil,
				Private:  nil,
				Fax:      nil,
				Other:    nil,
			},
		*/
		Note:     &note,
		Archived: false,
	}

	client := NewClient(getApiKey())
	contactResponse, err := client.CreateContact(&contact)

	assert.NoError(t, err, "error calling CreateContact")
	assert.NotNil(t, contactResponse, "no response")

	if client.debug {
		// fmt.Printf("%#v\n", contactResponse)
		fmt.Println(prettyPrintJson(contactResponse))
	}
}

func TestGetContact(t *testing.T) {
	client := NewClient(getApiKey())
	contact, err := client.GetContact("9919cd4d-0a0c-4fd0-857f-d3963335fc5a")

	assert.NoError(t, err, "error calling GetContact")
	assert.NotNil(t, contact, "no response")

	if client.debug {
		// fmt.Printf("%#v\n", contact)
		fmt.Println(prettyPrintJson(contact))
	}
}
