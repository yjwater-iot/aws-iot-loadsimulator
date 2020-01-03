#!/bin/bash

set -x

STACK_NAME=aws-iot-loadsimulator
ACCOUNT_ID=$(aws sts get-caller-identity | jq -r '.Account')
S3_BUCKET="${STACK_NAME}-${ACCOUNT_ID}"

if aws s3 ls "s3://$S3_BUCKET" 2>&1 | grep -q 'NoSuchBucket'
then
  echo "$S3_BUCKET does not exist. Creating now..."
  aws s3 mb s3://$S3_BUCKET
fi

./scripts/build-engine-lambda.sh
./scripts/build-worker-lambda.sh

sam package --template-file ./scripts/template.yml --s3-bucket $S3_BUCKET --output-template-file ./scripts/packaged.yml
sam deploy --template-file ./scripts/packaged.yml --stack-name $STACK_NAME --capabilities CAPABILITY_NAMED_IAM
