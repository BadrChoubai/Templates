#!/usr/bin/env bash

set -euo pipefail

./scripts/clean.sh

./scripts/build.sh

./scripts/push.sh

