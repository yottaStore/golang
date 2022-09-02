#!/bin/bash

go build -o dist/yfs.out \
 src/yfs.go

go build -o dist/store.out \
  src/store.go