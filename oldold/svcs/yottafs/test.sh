#!/bin/bash

curl -X POST http://localhost:8081/yottafs/ \
      -H 'Content-Type: application/json' \
      -d '{"Path":"test.txt","Method":"write","Data":"Http Write"}' \
      --output -

curl -X POST http://localhost:8081/yottafs/read \
   -H 'Content-Type: application/json' \
   -d '{"Path":"test.txt","Method":"read"}' \
   --output -

curl -X POST http://localhost:8081/yottafs/append \
      -H 'Content-Type: application/json' \
      -d '{"Method":"append","Path":"test.txt","Data":"Http Append"}' \
      --output -



