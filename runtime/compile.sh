#!/bin/bash
# http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail
IFS=$'\n\t'

cp /config/platformio.ini /build
cp /config/Configuration.h /build/src
cd /build
platformio run