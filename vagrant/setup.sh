#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace

export DEBIAN_FRONTEND=noninteractive

apt-get update -y
apt-get install -y \
  linux-image-4.18.0-16-generic \
  linux-headers-4.18.0-16-generic \
  make \
  docker.io
