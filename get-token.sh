#!/usr/bin/env bash

if [ -z "${OKTA_USERNAME}" ]; then
  echo "OKTA_USERNAME not set. Try exiting this container and rerunning ./go setup"
  exit 1
fi

if [ -z "${OKTA_SERVER}" ]; then
  echo "OKTA_SERVER not set. Try exiting this container and rerunning ./go setup"
  exit 1
fi

if [ -z "${OKTA_APPTYPE}" ]; then
  echo "OKTA_APPTYPE not set. Try exiting this container and rerunning ./go setup"
  exit 1
fi

if [ -z "${OKTA_APP_ID}" ]; then
  echo "OKTA_APP_ID not set. Try exiting this container and rerunning ./go setup"
  exit 1
fi


oktaauth \
  --username $OKTA_USERNAME \
  --server $OKTA_SERVER \
  --apptype $OKTA_APPTYPE \
  --appid $OKTA_APP_ID | \
  aws_role_credentials saml --profile default

cat  /root/.aws/credentials