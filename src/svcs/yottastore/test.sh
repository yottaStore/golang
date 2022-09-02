#!/bin/bash

curl -X POST http://localhost:8080/store/write \
      -H 'Content-Type: application/json' \
      -d '{"Record":"default@testTable/testRecord","Data":"Http Store Write"}' \
      --output -


curl -X POST http://localhost:8080/store/read \
      -H 'Content-Type: application/json' \
      -d '{"Record":"default@testTable/testRecord"}' \
      --output -


curl -X POST http://localhost:8080/store/append \
      -H 'Content-Type: application/json' \
      -d '{"Record":"default@testTable/testRecord","Data":"Http Store Append"}' \
      --output -



