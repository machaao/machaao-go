# machaao-go

![MessengerX logo](https://www.messengerx.io/img/logo.png)  

[![GoDoc](https://godoc.org/github.com/machaao/machaao-go?status.svg)](https://godoc.org/github.com/machaao/machaao-go) [![Go Report Card](https://goreportcard.com/badge/github.com/machaao/machaao-go)](https://goreportcard.com/report/github.com/machaao/machaao-go) 

machaao-go is a [go](https://golang.org) (or 'golang' for search engine friendliness) based library for building chatbot using [MessengerX](https://messengerx.io) APIs.

*This repository is community-maintained. We gladly accept pull requests. Please see the  [Machaao API Docs](https://ganglia.machaao.com/api-docs/#/) for all supported endpoints.*

The API docs for the [MessengerX.io](https://messengerx.io) is availabe at [MessengerX API Docs](https://messengerx.readthedocs.io/en/latest/) and [Machaao API Docs](https://ganglia.machaao.com/api-docs/#/).

### Install 

```bash
go get github.com/machaao/machaao-go
```

### Usage

1. Visit [MessengerX Docs](https://messengerx.readthedocs.io/en/latest/) and [Wit.ai](https://wit.ai) to get FREE API Token.

2. Put MachaaoAPIToken, MachaaoBaseURL and WitAPIToken on path. Follow [this](https://www.poftut.com/how-to-set-environment-variables-for-linux-windows-bsd-and-macosx/) article for more details on Environment Variable.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/machaao/machaao-go"
)

func main() {
	machaao.Server(messageHandler)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	//This function reads the request Body and saves to body as bytes.
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}

	//converts bytes to string
	var bodyData string = string(body)

	//incoming JWT Token
	var tokenString string = bodyData[8:(len(bodyData) - 2)]

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(machaao.MachaaoAPIToken), nil
	})

	_ = token

	if err != nil {
		fmt.Println(err)
	}

	//Do Something
}

```

