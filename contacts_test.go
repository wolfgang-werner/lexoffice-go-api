package lexoffice

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func getApiKey() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiKey, defined := os.LookupEnv("LEXOFFICE_API_KEY")
	if !defined || apiKey == "" {
		return "", errors.New("define non-empty environment variable LEXOFFICE_API_KEY (may be located in .env file")
	}

	return apiKey, nil
}

func createClient() *Client {
	apiKey, _ := getApiKey()
	return NewClient(apiKey)
}

func TestClient_CreateContact(t *testing.T) {
	client := createClient()

	// var _ *struct{} = nil

	var contact = Contact{
		ID:             "",
		OrganizationID: "",
		Version:        0,
		Roles: Roles{
			Customer: &Customer{},
		},
		// Company: nil,
		Person: &Person{
			LastName: "Müller",
		},
		Addresses: &Addresses{
			Billing: []Address{
				{
					Supplement:  "",
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
			Office:   nil,
			Private:  nil,
			Other:    nil,
		},
		PhoneNumbers: &PhoneNumbers{
			Business: nil,
			Office:   nil,
			Mobile:   nil,
			Private:  nil,
			Fax:      nil,
			Other:    nil,
		},
		Note:     "golang TEST",
		Archived: false,
	}

	contactResponse, err := client.CreateContact(&contact)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, contactResponse, "expecting non-nil result")

	fmt.Printf("%+v\n", contactResponse)
}

func TestGetContact(t *testing.T) {
	client := createClient()

	// ctx := context.Background()
	//contact, err := client.GetContact("4b5f0e33-59da-4caa-8721-e5d5b35cd4d4")
	contact, err := client.GetContact("4b5f0e33-59da-4caa-8721-e5d5b35cd4d4")

	// show go structure
	fmt.Printf("%+v\n", contact)

	// show as json
	res2B, _ := json.Marshal(contact)
	fmt.Println(string(res2B))

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, contact, "expecting non-nil result")

	// assert.Equal(t, 1, contact.Count, "expecting 1 face found")
	// assert.Equal(t, 1, contact.PagesCount, "expecting 1 PAGE found")

	// assert.Equal(t, "integration_face_id", contact.Faces[0].FaceID, "expecting correct face_id")
	// assert.NotEmpty(t, contact.Faces[0].FaceToken, "expecting non-empty face_token")
	// assert.Greater(t, len(contact.Faces[0].FaceImages), 0, "expecting non-empty face_images")
}
