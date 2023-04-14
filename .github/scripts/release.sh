#!/usr/bin/env bash
CONTAINERS=`bazelisk query --output=label 'kind("container_push", //...)'`

for container in $CONTAINERS; do
  bazelisk run ${container} --stamp --workspace_status_command="echo GITHUB_SHA $SHORT_SHA"
done
