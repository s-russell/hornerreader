#!/bin/bash
hash=$(git rev-parse --short HEAD)

docker build -t horner:$hash .
docker tag horner:$hash 0x7374657665/horner:$hash
docker push 0x7374657665/horner:$hash

docker tag horner:$hash 0x7374657665/horner:latest
docker push 0x7374657665/horner:latest