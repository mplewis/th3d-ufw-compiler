#!/bin/bash
# http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail
IFS=$'\n\t'

(cd apisrv; GOOS=linux go build)
docker build . --tag th3d_ufw_compiler
