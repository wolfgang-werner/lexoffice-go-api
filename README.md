# lexoffice Go API

This module implements a Go REST client for the public API from [lexoffice](https://www.lexoffice.de).
With the help of the API one can work with entities from the cloud-based accounting software. 

![lexoffice](https://www.lexoffice.de/wp-content/uploads/lexoffice-09-devices-dashboard-siegel.jpg)


## Setup

To work with the lexoffice API, you need to setup a lexoffice test instance and get an API key.  
Thankfully these steps can be done online (without credit card) in some minutes.

To get the test instance (provided 30 days free of charge by lexoffice), 
just hit https://www.lexoffice.de and create your test-account. 
You have to verify your email address and setup your password on first login.

After you have successfully logged in, just go to https://app.lexoffice.de/settings/#/public-api
to create your API key, attached to your instance and valid for 24 months. 

You can find the vendor explanation for this [in the cookbook](https://developers.lexoffice.io/cookbooks/public-api/#lexoffice-api-kochbuch-public-api-first-steps).

This module needs this API key to connect. You'll provide the API key as an environment 
variable or in an .env file (you can also hard-code it in your project if you dare).

To export the API key in an environment variable, just define it, e.g.

````shell
~ export LEXOFFICE_API_KEY=########-####-####-####-############
````

Alternatively, To create the .env file, just run:
````shell
~ echo "LEXOFFICE_API_KEY=########-####-####-####-############" > .env
````

You can check the file content with: 
````shell
~ cat .env
LEXOFFICE_API_KEY=########-####-####-####-############
````

Of course you can provide the API key in every command, also, eg.:
````shell
LEXOFFICE_API_KEY=########-####-####-####-############ go test -v -tags=integration
````


## Install module

```
go get -u github.com/wolfgang-werner/lexoffice-go-api
```


## Usage

```go
import "github.com/wolfgang-werner/lexoffice-go-api"

// Replace API_KEY with your real key
client := lexoffice.NewClient("API_KEY")

// enable debug output
client.debug = true


```


## Testing

Run integration tests with real API Key.

```shell
LEXOFFICE_API_KEY=########-####-####-####-############ go test -v -tags=integration
```


## Links

The following resources have been very helpful in writing this module.

- lexoffice [_API Documentation_](https://developers.lexoffice.io/docs/) 
  and [_Getting Started_](https://developers.lexoffice.io/cookbooks/public-api/#lexoffice-api-kochbuch-public-api) 

- Helpful advice from [Alex Pliutau](https://dev.to/plutov) - thank you Alex! - 
  good read [_Writing REST API Client in Go_](https://dev.to/plutov/writing-rest-api-client-in-go-3fkg)  
  and the corresponding GitHub repository [facest / factest-go](https://github.com/facest/facest-go) 

- [_JSON-to-Go_](https://mholt.github.io/json-to-go/), a cool online tool to convert JSON to Go type definition 

- Some in-deep knowledge - how to deal with dynamic JSON structures in Go.  
  Good starting points are [_Go JSON Cookbook_](https://eli.thegreenplace.net/2019/go-json-cookbook/) 
  and [_Representing JSON structures in Go_](https://eli.thegreenplace.net/2020/representing-json-structures-in-go/), both from _Eli Bendersky_ 
  and [_Dynamic JSON in Go_](https://eagain.net/articles/go-dynamic-json/) from _Tommi Virtanen_
  