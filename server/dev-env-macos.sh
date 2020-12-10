#!/usr/bin/bash

SCRIPT_PATH="`dirname $0`";
SCRIPT_PATH=`grealpath ${SCRIPT_PATH}`

export DOCLOCKER_DEPLOYMENT_ID="default";
export DOCLOCKER_CONFIG_DIR=`grealpath "${SCRIPT_PATH}/../config"`
export DOCLOCKER_DEPLOYMENT_ID=`cat ${DOCLOCKER_CONFIG_DIR}/latest/deployment-id`
export DOCLOCKER_SETUP_TEMPLATES_DIR="${SCRIPT_PATH}/templates"
