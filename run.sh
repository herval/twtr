#!/usr/bin/env bash

export $(cat env | xargs)
cd src
go build -o twtr .
./twtr
