#!/usr/bin/env bash

set -e
set -o nounset
set -o pipefail

SCRIPT_DIR=$(cd "$(dirname "$0")" ; pwd -P)
PROJECT_ROOT="${SCRIPT_DIR}/.."

goal_setup() {
  pushd "${SCRIPT_DIR}" > /dev/null
    config_file="${PROJECT_ROOT}/okta-config.txt"
    rm -f $config_file && touch $config_file
    echo "Adding okta-config.txt to your .gitignore"
    echo "okta-config.txt" >> ../.gitignore

    read -p "What is your username (e.g. meow@catnip.com)? " username_response
    if [ -z "${username_response}" ]; then
      echo "Username not provided"
      exit 1
    fi
    process_read "OKTA_USERNAME" "${username_response}" $config_file

    read -p "What is your appId? " app_id_response
    if [ -z "${app_id_response}" ]; then
      echo "AppID not provided"
      exit 1
    fi
    process_read "OKTA_APP_ID" "${app_id_response}" $config_file

    read -p "Please provide your Okta server hostname: " okta_server_host
    if [ -z "${okta_server_host}" ]; then
      echo "Okta Server not provided"
      exit 1
    fi
    process_read "OKTA_SERVER" "${okta_server_host}" $config_file

    read -p "Please provide your Application Type: " okta_app_type
    if [ -z "${okta_app_type}" ]; then
      echo "Application Type not provided"
      exit 1
    fi
    process_read "OKTA_APP_TYPE" "${okta_app_type}" $config_file
  popd > /dev/null

}

process_read() {
  field=$1
  response=$2
  output_file=$3

  usage="Usage: <func> field response output_file"

  if [ -z "${field}" ]; then
    "Field is not set. ${usage}"
  fi

  if [ -z "${response}" ]; then
    "Response is not set. ${usage}"
  fi

  if [ -z "${output_file}" ]; then
    "Output file is not set. ${usage}"
  fi

  if [ ! -z "${response}" ]; then
    echo "${field}=${response}" >> $output_file
  fi
}

goal_build() {
  pushd "${SCRIPT_DIR}" > /dev/null
      docker build -t okta-aws-portal  .
  popd > /dev/null
}

goal_run() {
  pushd "${SCRIPT_DIR}" > /dev/null
    mounted_dir=$(cd ${PROJECT_ROOT}; pwd)
    echo "Mounting ${mounted_dir}"

    docker run -it \
      --env-file "${mounted_dir}/okta-config.txt" \
      -v "${mounted_dir}:/app" \
      okta-aws-portal ./get-token.sh
  popd > /dev/null
}

TARGET=${1:-}
if type -t "goal_${TARGET}" &>/dev/null; then
  "goal_${TARGET}" ${@:2}
else
  echo "Usage: $0 <goal>

goal:
    setup                   - Set up OKTA config
    build                   - Builds container
    run                     - Runs container
"
  exit 1
fi
