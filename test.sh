#!/bin/bash

CWD=$(pwd)

# YottaFS tests
cd $CWD/svcs/yottafs
go test ./...