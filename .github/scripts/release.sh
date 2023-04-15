#!/usr/bin/env bash
CONTAINERS=`bazelisk query --output=label 'kind("container_push", //...)'`

for container in $CONTAINERS; do
  bazelisk run ${container} --stamp --workspace_status_command="echo RELEASE_TAG $RELEASE_TAG"
done
