#!/bin/bash

/usr/local/bin/mongos-start.sh &

exec /usr/local/bin/docker-entrypoint.sh "$@"
