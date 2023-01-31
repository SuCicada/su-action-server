#!/bin/bash

echo ${WEBHOOK_URL}
env_file_contents=$(cat .env)
echo ${env_file_contents}

act -j deploy -s WEBHOOK_URL=${WEBHOOK_URL} \
  -s ENV_FILE_CONTENT="$env_file_contents" \
  --container-architecture linux/arm64 \
