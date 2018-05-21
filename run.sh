#!/usr/bin/env bash

export $(cat env | xargs)
go build -o twtr .
./twtr
