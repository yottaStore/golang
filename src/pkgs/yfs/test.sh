#!/bin/bash

curl -X POST http://localhost:8080/store/read \
   -H 'Content-Type: application/json' \
   -d '{"Path":"/home/mamluk/yotta/yottaStore-go/src/libs/yfs/drivers/direct/test/test.txt"}' \


   curl -X POST http://localhost:8080/store/append \
      -H 'Content-Type: application/json' \
      -d '{"Path":"/home/mamluk/yotta/yottaStore-go/src/pkgs/yfs/drivers/direct/test/test.txt","Data":"Http Append"}'


