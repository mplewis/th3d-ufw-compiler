#!/bin/bash
# http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail
IFS=$'\n\t'

docker run -it -p 8080:8080 th3d_ufw_compiler
