#!/bin/bash

go get .
go build -o "build/${PWD##*/}"
