#!/bin/bash

curl -X POST http://localhost:8080/yfs/read \
   -H 'Content-Type: application/json' \
   -d '{"Path":"test.txt"}' \
   --output -


   curl -X POST http://localhost:8080/yfs/append \
      -H 'Content-Type: application/json' \
      -d '{"Path":"test.txt","Data":"Http Append"}' \
      --output -


