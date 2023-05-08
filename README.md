# flight-tracking
A simple microservice API that can help us understand and track how a particular person's flight path may be queried

Go version: 1.20

The microservice is built by using [gin framework](https://github.com/gin-gonic/gin)

### Endpoints
This microservice contains only one endpoint, i.e /calculate
```
// Reqeust body: 
{
    "flights": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] 
}
// where flights is a 2D array of inputs of a list of flights according to the requirements

// Response body:
{
    "result": [
        "SFO",
        "EWR"
    ]
}
// where result is an array represents the output according to the requirements
```
Curl command
``` bash
curl --location 'localhost:8080/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "flights": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] 
}' 
```

### How to run
Simply go to root of the project folder and run below command
```bash
go run main.go
```

### Unit Test
Please run below make command to run the unit test, it will generate a **coverage.html** in the root of the project folder, which is a report to show the coverage 
```bash
make unit-test
```
So far only the code with implementations are unit tests covered
```
?       flight-tracking [no test files]
?       flight-tracking/internal/constant       [no test files]
?       flight-tracking/routes  [no test files]
ok      flight-tracking/handler 0.984s  coverage: 100.0% of statements
ok      flight-tracking/internal/flighttracker  0.614s  coverage: 100.0% of statements
```

### Linting
[golangci-lint](https://github.com/golangci/golangci-lint) is used for linting
Please run below make command to execute lint checking
```bash
make lint
```