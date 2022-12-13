# MyGoReminders

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

Run the localhost url in your browser

/books => getBooks (GET)
/books/:id => bookById (GET)
/books => createBook (POST)
/checkout => checkoutBook (PATCH use Query param ?id=(n)) 
/return => returnBook (PATCH use Query param ?id=(n)) 



## Prerequisites

**Make sure GoLang is installed globally**

add Go path shortcut to zshrc or bashrc

```
export PATH=$PATH:$(go env GOPATH)/bin
```

 

 