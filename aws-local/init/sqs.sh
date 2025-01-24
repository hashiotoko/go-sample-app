#!/bin/bash

if [ -z "$AWS_DEFAULT_REGION" ]; then
  echo "Error: AWS_DEFAULT_REGION is not set"
  exit 1
fi

if [ -z "$QUEUE_NAME_FOR_WELCOME_MAIL" ]; then
  echo "Using default queue name for welcome mail"
  QUEUE_NAME_FOR_WELCOME_MAIL="welcome-mail"
fi

queue_names=(
  $QUEUE_NAME_FOR_WELCOME_MAIL
)

for queue_name in "${queue_names[@]}"; do
  # キュー作成
  awslocal sqs create-queue --queue-name ${queue_name}
  echo "Created queue: ${queue_name} in region: ${AWS_DEFAULT_REGION}"

  # キュー作成確認
  if awslocal sqs list-queues --queue-name-prefix ${queue_name} | grep -q ${queue_name}; then
    echo "Queue verification successful"
  else
    echo "Error: Failed to create ${queue_name} queue"
    exit 1
  fi
done
