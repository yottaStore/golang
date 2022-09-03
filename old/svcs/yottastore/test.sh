#!/bin/bash

curl -X POST http://localhost:8080/yottastore/ \
      -H 'Content-Type: application/json' \
      -d '{"Record":"default@testTable/testRecord","Data":"Http Store Write", "Method":"write"}' \
      --output -


curl -X POST http://localhost:8080/yottastore/ \
      -H 'Content-Type: application/json' \
      -d '{"Record":"default@testTable/testRecord", "Method":"read"}' \
      --output -


curl -X POST http://localhost:8080/store/append \
      -H 'Content-Type: application/json' \
      -d '{"Record":"default@testTable/testRecord","Data":"Http Store Append"}' \
      --output -



