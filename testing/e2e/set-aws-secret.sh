#!/bin/bash
# This file gets the s3user credentials from the secret and sets as the aws-cli credentials.
# Usage: ./script.sh PROFILE_NAME SECRET_NAME
# Example: ./script.sh ceph-test s3-sample-admin-secret

if [ "$#" -ne 2 ]; then
  echo "Usage: $0 PROFILE_NAME SECRET_NAME"
  exit 1
fi

PROFILE_NAME="$1"
SECRET_NAME="$2"
REGION="us-east-1"

# Get the access key and secret key from Kubernetes secret
ACCESS_KEY=$(kubectl get secret $SECRET_NAME -n s3-test -o jsonpath="{.data.accessKey}" | base64 --decode)
SECRET_ACCESS_KEY=$(kubectl get secret $SECRET_NAME -n s3-test -o jsonpath="{.data.secretKey}" | base64 --decode)

# Update the existing profile or create a new one
aws configure set aws_access_key_id $ACCESS_KEY --profile $PROFILE_NAME
aws configure set aws_secret_access_key $SECRET_ACCESS_KEY --profile $PROFILE_NAME
aws configure set region $REGION --profile $PROFILE_NAME