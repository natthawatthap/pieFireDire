How to start api 
```go run main.go```
```curl --location 'http://localhost:3000/beef/summary'```

How to run test 
```cd meatcounter```
```go test```

How to run performance test script
```k6 run --out json=result.json performance-test-script.js > result.txt```