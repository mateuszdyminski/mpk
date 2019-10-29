#!/usr/bin/env bash

set -e

usage() {
	cat <<EOF
Usage: $(basename $0) <command>

Wrappers around core binaries:
    start                  Starts MPK tracker.
EOF
	exit 1
}

if [ $# -ne 2 ]; then
    usage
    exit 1
fi

source .env

TYPE=$1
NUMBER=$2

[ -e statics/index.html ] && rm statics/index.html

sed "s/\$\$VEHICLE_TYPE/$TYPE/g" statics/_index.html >> statics/index.html
sed -i "s/\$\$NUMBER/$NUMBER/g" statics/index.html
sed -i "s/\$\$GOOGLE_MAPS_KEY/$GOOGLE_MAPS_KEY/g" statics/index.html

go run main.go