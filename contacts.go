package lexoffice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Contact struct {
	ID             *string         `json:"id"`
	OrganizationID *string         `json:"organizationId"`
	Version        int             `json:"version"`
	Roles          Roles           `json:"roles"`
	Company        *Company        `json:"company"`
	Person         *Person         `json:"person"`
	Addresses      *Addresses      `json:"addresses"`
	EmailAddresses *EmailAddresses `json:"emailAddresses"`
	PhoneNumbers   *PhoneNumbers   `json:"phoneNumbers"`
	Note           *string         `json:"note"`
	Archived       bool            `json:"archived"`
}

type Roles struct {
	Customer *Customer `json:"customer"`
	Vendor   *Vendor   `json:"vendor"`
}

type Customer struct {
	Number *int `json:"number"`
}

type Vendor struct {
	Number *int `json:"number"`
}

type Company struct {
	Name                 string           `json:"name"`
	TaxNumber            *string          `json:"taxNumber"`
	VatRegistrationID    *string          `json:"vatRegistrationId"`
	AllowTaxFreeInvoices bool             `json:"allowTaxFreeInvoices"`
	ContactPersons       []ContactPersons `json:"contactPersons"`
}

type ContactPersons struct {
	Salutation   *string `json:"salutation"`
	FirstName    *string `json:"firstName"`
	LastName     string  `json:"lastName"`
	Primary      bool    `json:"primary"`
	EmailAddress *string `json:"emailAddress"`
	PhoneNumber  *string `json:"phoneNumber"`
}

type Person struct {
	Salutation *string `json:"salutation"`
	FirstName  *string `json:"firstName"`
	LastName   string  `json:"lastName"`
}

type Addresses struct {
	Billing  []Address `json:"billing"`
	Shipping []Address `json:"shipping"`
}

type Address struct {
	Supplement  *string `json:"supplement"`
	Street      *string `json:"street"`
	Zip         *string `json:"zip"`
	City        *string `json:"city"`
	CountryCode string  `json:"countryCode"`
}

type EmailAddresses struct {
	Business []string `json:"business"`
	Office   []string `json:"office"`
	Private  []string `json:"private"`
	Other    []string `json:"other"`
}

type PhoneNumbers struct {
	Business []string `json:"business"`
	Office   []string `json:"office"`
	Mobile   []string `json:"mobile"`
	Private  []string `json:"private"`
	Fax      []string `json:"fax"`
	Other    []string `json:"other"`
}

type CreateContactResponse struct {
	ID          string    `json:"id"`
	ResourceURI string    `json:"resourceUri"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	Version     int       `json:"version"`
}

type LookupContactsResponse struct {
	Content          []Contact `json:"content"`
	First            bool      `json:"first"`
	Last             bool      `json:"last"`
	Number           int       `json:"number"`
	NumberOfElements int       `json:"numberOfElements"`
	Size             int       `json:"size"`
	Sort             []Sort    `json:"sort"`
	TotalElements    int       `json:"totalElements"`
	TotalPages       int       `json:"totalPages"`
}

type Sort struct {
	Ascending    bool   `json:"ascending"`
	Direction    string `json:"direction"`
	IgnoreCase   bool   `json:"ignoreCase"`
	NullHandling string `json:"nullHandling"`
	Property     string `json:"property"`
}

// GetContact returns contact object by ContactID
func (c *Client) GetContact(contactID string) (*Contact, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/contacts/%s", c.baseURL, url.PathEscape(contactID)), nil)
	if err != nil {
		return nil, err
	}

	res := Contact{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateContact creates a new contact (person or company)
func (c *Client) CreateContact(contact *Contact) (*CreateContactResponse, error) {
	jsonValue, err := json.Marshal(contact)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/contacts", c.baseURL), bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var res CreateContactResponse
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// UpdateContact creates a new contact (person or company)
func (c *Client) UpdateContact(contact *Contact) (*CreateContactResponse, error) {
	jsonValue, err := json.Marshal(contact)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/contacts/%s", c.baseURL, *contact.ID), bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var res CreateContactResponse
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type Filter struct {
	name  string
	value string
}

// NewFilter creates a new filter
func NewFilter(k, v string) Filter {
	if !(k == "email" || k == "name" || k == "number" || k == "customer" || k == "vendor") {
		fmt.Printf("invalid filter %s)%s", k, v)
	}

	return Filter{
		name:  k,
		value: v,
	}
}

type Pagination struct {
	page int
	size int
}

func buildFilterParameter(filters []Filter, pagination Pagination) string {
	if filters == nil || len(filters) == 0 {
		return "?page=" + strconv.Itoa(pagination.page) + "&size=" + strconv.Itoa(pagination.size)
	}

	var parameter string
	for ndx, filter := range filters {
		if ndx == 0 {
			parameter = "?" + filter.name + "=" + url.PathEscape(filter.value)
		} else {
			parameter = parameter + "&" + filter.name + "=" + url.PathEscape(filter.value)
		}
	}

	return parameter + "&page=" + strconv.Itoa(pagination.page) + "&size=" + strconv.Itoa(pagination.size)
}

func (c *Client) LookupContacts(filters []Filter, pagination Pagination) (*LookupContactsResponse, error) {
	filter := buildFilterParameter(filters, pagination)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/contacts%s", c.baseURL, filter), nil)
	if err != nil {
		return nil, err
	}

	var res LookupContactsResponse
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
