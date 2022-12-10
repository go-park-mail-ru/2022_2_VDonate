#!/bin/bash

PREFIX=zeronethunter/vdonate

# building Docker images
for f in $(find .. -name 'Dockerfile')
do
  echo "BUILD $f"
  BASE=$(basename "$(dirname "${f}")")
  docker pull ${PREFIX}-"${BASE}":latest
done