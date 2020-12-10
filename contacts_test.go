package lexoffice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createCustomerPerson() *Contact {
	var note = "creator: lexoffice-go-api"
	var supplement = "billing address"
	var street = "Teststraße 123"
	var zip = "20355"
	var city = "Hamburg"

	return &Contact{
		Version: 0,
		Roles: Roles{
			Customer: &Customer{},
		},
		Person: &Person{
			LastName: "Müller",
		},
		Addresses: &Addresses{
			Billing: []Address{
				{
					Supplement:  &supplement,
					Street:      &street,
					Zip:         &zip,
					City:        &city,
					CountryCode: "DE",
				},
			},
		},
		EmailAddresses: &EmailAddresses{
			Office: []string{"wwerner@outlook.com"},
		},
		Note:     &note,
		Archived: false,
	}
}

func createCustomerCompany() *Contact {
	var note = "creator: lexoffice-go-api"
	var supplement = "billing address"
	var street = "Teststraße 123"
	var zip = "20355"
	var city = "Hamburg"

	return &Contact{
		Version: 0,
		Roles: Roles{
			Customer: &Customer{},
		},
		Company: &Company{
			Name: "ACME Ltd",
		},
		Addresses: &Addresses{
			Billing: []Address{
				{
					Supplement:  &supplement,
					Street:      &street,
					Zip:         &zip,
					City:        &city,
					CountryCode: "DE",
				},
			},
		},
		EmailAddresses: &EmailAddresses{
			Office: []string{"wwerner@outlook.com"},
		},
		Note:     &note,
		Archived: false,
	}
}

func TestClient_CreateCustomerPerson(t *testing.T) {
	client := NewClient(getApiKey())
	contactResponse, err := client.CreateContact(createCustomerPerson())

	assert.NoError(t, err, "error calling CreateContact")
	assert.NotNil(t, contactResponse, "no response")

	if client.debug {
		// fmt.Printf("%#v\n", contactResponse)
		fmt.Println(prettyPrintJson(contactResponse))
	}
}

func TestClient_CreateCustomerCompany(t *testing.T) {
	client := NewClient(getApiKey())
	contactResponse, err := client.CreateContact(createCustomerCompany())

	assert.NoError(t, err, "error calling CreateContact")
	assert.NotNil(t, contactResponse, "no response")

	if client.debug {
		// fmt.Printf("%#v\n", contactResponse)
		fmt.Println(prettyPrintJson(contactResponse))
	}
}

func TestGetContact(t *testing.T) {
	client := NewClient(getApiKey())
	// contact, err := client.GetContact("9919cd4d-0a0c-4fd0-857f-d3963335fc5a")
	contact, err := client.GetContact("6a5ff2f9-3fe2-4ea1-ae05-a279d28cb55b")

	assert.NoError(t, err, "error calling GetContact")
	assert.NotNil(t, contact, "no response")

	if client.debug {
		// fmt.Printf("%#v\n", contact)
		fmt.Println(prettyPrintJson(contact))
	}
}
