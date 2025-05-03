#!/bin/bash

until mongosh -u "$MONGO_INITDB_ROOT_USERNAME" -p "$MONGO_INITDB_ROOT_PASSWORD" --eval "db.adminCommand('ping')" >/dev/null 2>&1; do
    echo "Waiting for MongoDB to be ready..."
    sleep 2
done

echo "MongoDB is ready. Creating user..."

MONGO_USER="krake"
MONGO_PASSWORD="123456789"

mongosh -u "$MONGO_INITDB_ROOT_USERNAME" -p "$MONGO_INITDB_ROOT_PASSWORD" admin <<EOF
    use es_krake;
    db.createUser({
        user: "$MONGO_USER",
        pwd: "$MONGO_PASSWORD",
        roles: ["readWrite"]
    });
EOF
