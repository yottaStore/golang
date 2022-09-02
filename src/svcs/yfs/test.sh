#!/bin/bash

curl -X POST http://localhost:8081/yfs/ \
      -H 'Content-Type: application/json' \
      -d '{"Path":"test.txt","Method":"write","Data":"Http Write"}' \
      --output -

curl -X POST http://localhost:8081/yfs/read \
   -H 'Content-Type: application/json' \
   -d '{"Path":"test.txt","Method":"read"}' \
   --output -

curl -X POST http://localhost:8081/yfs/append \
      -H 'Content-Type: application/json' \
      -d '{"Method":"append","Path":"test.txt","Data":"Http Append"}' \
      --output -



