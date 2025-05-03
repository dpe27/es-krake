#!/bin/bash

# Check if the replica set initialization flag is enabled
if [[ "$DO_INIT_REPSET" = true ]]; then
    echo "Waiting for MongoDB to be ready..."
    # Wait until mongosh can connect to MongoDB
    until /usr/bin/mongosh --port 27017 --quiet --eval 'db.getMongo()'; do
        sleep 1
    done

    echo "MongoDB is up. Checking replica set status..."
    /usr/bin/mongosh --port 27017 <<EOF
        rs.initiate({
            _id: "${REPSET_NAME}",
            settings: {electionTimeoutMillis: 2000},
            members: [
                {_id: 0, host: "${REPSET_NAME}-replica0:27017"},
                {_id: 1, host: "${REPSET_NAME}-replica2:27017"},
                {_id: 2, host: "${REPSET_NAME}-replica3:27017"}
            ]
        })
EOF
fi
