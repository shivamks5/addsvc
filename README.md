# addsvc - Simple Add Service using Go Kit

`addsvc` is a minimal microservice built using **Go Kit** that exposes an HTTP endpoint to add two numbers.

## ▶️ How to Run

Make sure you have [Go](https://golang.org/) installed.

Then run the service using:

```bash
go run main.go
```

This will start the server at:

http://localhost:8080

## 📥 Example Request

Once the server is running, you can make a GET request to:

GET http://localhost:8080/add?a=10&b=20

## 📤 Example Response

```json
{
  "result": 30
}
```
