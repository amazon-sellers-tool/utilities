#!/bin/bash
cd $1
rm go.mod
rm go.sum
go mod init $1
go mod tidy
