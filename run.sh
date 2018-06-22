#!/usr/bin/env bash

set -ex
export $(cat env | xargs)
./build.sh
./twtr
