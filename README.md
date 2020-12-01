# toyyibpay-go
Toyyibpay unnofficial client SDK written in Golang


## Coverage (WIP)
## Usage
### Create client
```go
import "github.com/qhkm/toyyibpay"

// Create client instance
client, err := toyyibpay.NewClient(secretKey)
```

### Create single bill
```go
params := toyyibpay.CreateBillParams{}
code, err := client.CreateSingleBill(params)
```

### Run bill
```go
params := toyyibpay.RunBillParams{}
htmlTag, err := client.RunBill(rbp)
```

### Get bill Transaction
```go
params := toyyibpay.CreateBillTransactions{}
resp, err := client.GetTransactions(params)

```

### Work in progress ...


## How to Contribute
   - Fork a repository
   - Add/Fix something
   - Check that tests are passing
   - Create PR

## Tests

   - Unit tests: go test
   - Integration tests: go test -tags=integration
