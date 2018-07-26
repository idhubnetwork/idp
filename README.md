# IDHub Identity Provider（IDHub IdP）

From this moment on, you only need to have an IDHub identity. You no longer need to say hello to the chief, you can easily shuttle in the Amazon rainforest, just one click you can do it.

## Feature Overview

- IDHub Identity
- SAML Protocol
- AWS Services

## Requirements

- [Go](https://golang.org)
- [Node.js](https://nodejs.org/en/)
- [MySQL](https://www.mysql.com)
- [IDHub SAML API](https://gitlab.com/idhub/saml_idhub_api)
- [MetaMask](https://metamask.io)

## Installation

1. Create a MySQL database with name "idp" and services at localhost:3306.

2. Table Schema with SQL statements in [db.sql](https://gitlab.com/idhub/idp/blob/master/dbconn/db.sql)

3. Generate RS256 key for JWT

```sh
$ ./jwt_rs256.sh
$ ls keystore
```

4. Frontend build

```sh
$ npm i
$ npm run build
```

5. Install go packages

- using go get

```sh
$ go get github.com/labstack/echo
$ go get golang.org/x/sync
$ go get github.com/ethereum/go-ethereum
$ go get gopkg.in/go-playground/validator.v9
$ go get github.com/gobuffalo/pop
$ go get github.com/dgrijalva/jwt-go
$ go get github.com/parnurzeal/gorequest
```

- using vendor tool ([dep](https://github.com/golang/dep))

```sh
$ dep ensure
```

6. Startup

```sh
$ go run main.go
```

7. [Try it](http://127.0.0.1:8080/#/)

## License

[Apache License 2.0](https://gitlab.com/idhub/idp/blob/master/LICENSE)