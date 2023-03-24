#!/bin/bash

if [ -z "$1" ]
then
  PORT=8080
else
  PORT="$1"
fi

USERID="Ub4601ea7f9eb45d86b3c1e57214675a2"

echo "=== Test sever apis on port: $PORT ==="

echo -e "Test send message to line\n"
curl -X POST localhost:$PORT/line_message/send \
   -H 'Content-Type: application/json' \
   -d "{\"target_id\":\"$USERID\",\"message\":\"Hello World\"}"

echo -e "Test get all user message \n"
curl -X GET localhost:$PORT/line_message/Ub4601ea7f9eb45d86b3c1e57214675a2