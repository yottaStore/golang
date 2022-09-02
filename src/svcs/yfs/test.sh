#!/bin/bash

curl -X POST http://localhost:8081/yfs/write \
      -H 'Content-Type: application/json' \
      -d '{"Path":"test.txt","Data":"Http Write"}' \
      --output -

curl -X POST http://localhost:8081/yfs/read \
   -H 'Content-Type: application/json' \
   -d '{"Path":"test.txt"}' \
   --output -

curl -X POST http://localhost:8081/yfs/append \
      -H 'Content-Type: application/json' \
      -d '{"Path":"test.txt","Data":"Http Append"}' \
      --output -



