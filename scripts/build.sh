#!/bin/bash

PREFIX=zeronethunter/vdonate

# building Docker image
docker build -f ../deployments/dev/api/Dockerfile ../ --tag ${PREFIX}-api
