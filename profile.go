package lexoffice

import (
	"fmt"
	"net/http"
)

// Profile is the response for the lexoffice /profile request, delivering information regarding your account
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

//GetProfile returns the profile information
//
//https equivalent request/response (sensitive information replaced by e.g. [api-key])
//
//  ~ https -p hbHB GET "api.lexoffice.io/v1/profile" "Authorization: Bearer [api-key]"
//  GET /v1/profile HTTP/1.1
//  Accept: */*
//  Accept-Encoding: gzip, deflate
//  Authorization: Bearer [api-key]
//  Connection: keep-alive
//  Host: api.lexoffice.io
//  User-Agent: HTTPie/2.3.0
//
//  HTTP/1.1 200 OK
//  Connection: keep-alive
//  Content-Length: 355
//  Content-Type: application/json
//  Date: Sun, 13 Dec 2020 10:05:34 GMT
//  X-Amzn-Trace-Id: Root=1-5fd5e76e-20d04acc1f000e3573c5a77c
//  x-amz-apigw-id: XfEZPE9zFiAFXTA=
//  x-amzn-RequestId: 56045147-6a01-44d4-8182-6784ea49b19f
//
//  {
//    "companyName": "[profile name]",
//    "connectionId": "[connectionId]",
//    "created": {
//      "date": "2020-12-03T10:33:21.597+01:00",
//      "userEmail": "[email]",
//      "userId": "[userId]",
//      "userName": "[userName]"
//    },
//    "organizationId": "[organizationId]",
//    "smallBusiness": false,
//    "taxType": "net"
//  }
//
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
