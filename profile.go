package lexoffice

import (
	"fmt"
	"net/http"
)

type Profile struct {
	OrganizationID string `json:"organizationId"`
	CompanyName    string `json:"companyName"`
	Created        struct {
		UserID    string `json:"userId"`
		UserName  string `json:"userName"`
		UserEmail string `json:"userEmail"`
		Date      string `json:"date"`
	} `json:"created"`
	ConnectionID       string   `json:"connectionId"`
	Features           []string `json:"features"`
	SubscriptionStatus string   `json:"subscriptionStatus"`
	TaxType            string   `json:"taxType"`
	SmallBusiness      bool     `json:"smallBusiness"`
}

// GetProfile returns the profile information
func (c *Client) GetProfile() (*Profile, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/profile", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	res := Profile{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
