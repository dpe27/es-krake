#!/bin/bash

# Wait until mongosh can return a connection
echo "Waiting for MongoDB to be ready..."
until /usr/bin/mongosh --quiet --eval 'db.getMongo()'; do
    sleep 1
done
echo "MongoDB is ready."

# Check if SHARD_LIST is set
if [[ -z "$SHARD_LIST" ]]; then
    echo "Error: SHARD_LIST is not set."
    exit 1
fi

# Split set of shard URLs text by ';' separator
IFS=';' read -r -a array <<<"$SHARD_LIST"

# Add each shard definition to the cluster
echo "Adding shards to the cluster..."
for shard in "${array[@]}"; do
    echo "Adding shard: ${shard}"
    /usr/bin/mongosh --port 27017 --quiet <<EOF
        sh.addShard("${shard}")
EOF
done

echo "All shards added successfully."
