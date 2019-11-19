# Example go-kit micro service intergrate with postgres db

## Server start on port 8080

`$ go run main.go`

`$ curl -XPOST -d'{"email": "sanprasirt@pccth.com", "password": "password"}' localhost:8080/user`

`$ curl localhost:8080/user/f8b9ea14-2d41-4637-9e66-5ac3781f1e4c`
