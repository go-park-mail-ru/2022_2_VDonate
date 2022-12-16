#!/bin/bash

PREFIX=zeronethunter/vdonate

# building Docker image
docker build -f ../deployments/deploy/api/Dockerfile ../ --tag ${PREFIX}-api
