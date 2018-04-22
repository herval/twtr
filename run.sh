#!/usr/bin/env bash

export $(cat env | xargs)
cd src
go run *.go
