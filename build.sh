#!/bin/bash +vx


# build server
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/mic_server mic_srv_office

# build api_server
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/api_server mic_srv_office_api/bin

# docker-compose down
docker-compose down


# docker-compose up
docker-compose -f docker-compose.yml  up 