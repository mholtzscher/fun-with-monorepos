#!/usr/bin/env bash
CONTAINERS=`bazelisk query --output=label 'kind("container_push", //...)'`

for container in $CONTAINERS; do
  bazelisk run -- ${container}
done
