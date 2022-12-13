# Tech Book API

To Build

```
go build
```

To Run

```
go run main.go 
```


## Run API via Postman or RapidAPI or view in Browser

Run the localhost url in your browser

[http://localhost:8080](http://localhost:8080)

or see run the api via:

```
curl http://localhost:8080
```


## Routes

Route         | Method        | Call
------------- | ------------- | -------------
/books  | GET  | getBooks
/books/:id  | GET  | bookById
/books  | POST  | createBook
/checkout  | PATCH  | checkoutBook
/return  | PATCH  | returnBook

For PATCH use Query param ?id=(n)



## Prerequisites

**Make sure GoLang is installed globally**

add Go path shortcut to zshrc or bashrc

```
export PATH=$PATH:$(go env GOPATH)/bin
```

 

 