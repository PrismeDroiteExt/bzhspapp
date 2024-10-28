#!/bin/sh

go mod download ;
go build -o ./cmd/web/main ./cmd/web/main.go ; 
go run ./cmd/web/main.go ;