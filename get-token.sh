#!/usr/bin/env bash

oktaauth \
  --username $OKTA_USERNAME \
  --server $OKTA_SERVER \
  --apptype $OKTA_APPTYPE \
  --appid $OKTA_APP_ID | \
  aws_role_credentials saml --profile default

cat  /root/.aws/credentials