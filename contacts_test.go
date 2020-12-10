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
	var zip = "20000"
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
	var zip = "20000"
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
			Office: []string{"mail@example.com"},
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

func TestClient_UpdateContact(t *testing.T) {
	client := NewClient(getApiKey())

	// create and retrieve new contact
	initialContact := createCustomerCompany()
	contactResponse, _ := client.CreateContact(initialContact)
	contact, _ := client.GetContact(contactResponse.ID)

	// assert initial data is as expected
	assert.Nil(t, contact.PhoneNumbers, "phone number should be nil")
	assert.Equal(t, safeString(initialContact.Note), safeString(contact.Note), "expecting initial note")

	// update contact with new note and a new phone number
	var note = "UPDATED " + safeString(contact.Note)
	var phone = "+49 000 000 00 00"
	contact.Note = &note
	contact.PhoneNumbers = &PhoneNumbers{Mobile: []string{phone}}

	updateResponse, err := client.UpdateContact(contact)
	assert.NoError(t, err, "error calling GetContact")
	assert.NotNil(t, updateResponse, "no response")
	if client.debug {
		// fmt.Printf("%#v\n", contact)
		fmt.Println(prettyPrintJson(updateResponse))
	}

	// retrieve new contact, test for change
	updatedContact, _ := client.GetContact(updateResponse.ID)
	assert.NotNil(t, updatedContact.PhoneNumbers.Mobile, "phone number expected")
	assert.Equal(t, 1, len(updatedContact.PhoneNumbers.Mobile), "one number expected")
	assert.Equal(t, phone, updatedContact.PhoneNumbers.Mobile[0], "new phone number expected")
	assert.Equal(t, note, safeString(updatedContact.Note), "new note expected")
}

// the guid has to be changed to an existing contact guid!
func TestGetContact(t *testing.T) {
	client := NewClient(getApiKey())

	contact, err := client.GetContact("6a5ff2f9-3fe2-4ea1-ae05-a279d28cb55b")
	assert.NoError(t, err, "error calling GetContact")
	assert.NotNil(t, contact, "no response")
	if client.debug {
		// fmt.Printf("%#v\n", contact)
		fmt.Println(prettyPrintJson(contact))
	}
}

func TestLookupAll(t *testing.T) {
	client := NewClient(getApiKey())
	client.debug = true

	filters := []Filter{
		NewFilter("number", "10001"),
		NewFilter("customer", "true"),
	}
	pagination := Pagination{page: 0, size: 5}

	response, err := client.LookupContacts(filters, pagination)
	assert.NoError(t, err, "error calling LookupContact")
	assert.NotNil(t, response, "no response")
	if client.debug {
		// fmt.Printf("%#v\n", response)
		fmt.Println(prettyPrintJson(response))
	}
}
