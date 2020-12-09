# lexoffice-go-api

## Install

```
go get -u github.com/wolfgang-werner/lexoffice-go-api
```

## Usage

```
import "github.com/wolfgang-werner/lexoffice-go-api"

// Replace API_KEY with your real key
client := lexoffice.NewClient("API_KEY")
```

## Testing

Run integration tests with real API Key.

```
LEXOFFICE_API_KEY=XXX go test -v -tags=integration
```
