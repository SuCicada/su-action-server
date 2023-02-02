#!/bin/bash

echo ${WEBHOOK_URL}
env_file_contents=$(cat .env)
echo ${env_file_contents}
architecture="linux/amd64"
if [ "$(uname -m)" == "arm64" ]; then
  architecture="linux/arm64"
fi

act -j deploy \
  --container-architecture $architecture \
  --secret-file ~/.config/gh/.env.secret \
  -s WEBHOOK_URL=${WEBHOOK_URL} \
  -s DEPLOY_PATH=$DEPLOY_PATH \
  -s SSH_KEY="$(cat ~/.ssh/id_rsa)" \
  --insecure-secrets \
#  --use-gitignore true \
#  -s DEBUG=true \
#  -v


