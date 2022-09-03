#!/bin/bash

cd src/yottafs
go build -o ../../dist/yottafs.out \
 cmd/yottafs.go

cd ../../src/yottastore
go build -o ../../dist/yottastore.out \
 cmd/yottastore.go
