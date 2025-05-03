#!/bin/bash

/usr/local/bin/init-replica.sh &

exec /usr/local/bin/docker-entrypoint.sh "$@"
