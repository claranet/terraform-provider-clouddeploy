#!/bin/bash

go get .
go build -o "${PWD##*/}"
