#!/usr/bin/env bash

# Basic HAL config
export HAL_NAME=""
export HAL_ALIAS=""
export HAL_ADAPTER=slack
export HAL_STORE=memory
export HAL_PORT=9000
export HAL_LOG_LEVEL=debug
export HAL_DEBUG_ENDPOINT=1

# Auth Config
export HAL_AUTH_ENABLED=1
export HAL_AUTH_ADMIN=""

# Slack config
export HAL_SLACK_TOKEN=
export HAL_SLACK_CHANNELS=""
export HAL_SLACK_BOTNAME=""

chmod a+x ./build/hal
./build/hal

exit 0