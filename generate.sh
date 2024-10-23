#!/bin/bash

PROTO_DIR="proto"
OUT_DIR="./gen/go"
BRANCH="main"

LAST_TAG=$(git describe --tags $(git rev-list --tags --max-count=1))
if [[ -z "$LAST_TAG" ]]; then
    NEW_TAG="0.0.1"
else
    IFS='.' read -r major minor patch <<<"${LAST_TAG#v}"
    NEW_TAG="v$major.$minor.$((patch + 1))"
fi

find "$PROTO_DIR" -name "*.proto" | while read -r proto_file; do
    echo "Processing $proto_file..."
    protoc -I "$PROTO_DIR" "$proto_file" --go_out="$OUT_DIR" --go_opt=paths=source_relative --go-grpc_out="$OUT_DIR" --go-grpc_opt=paths=source_relative
    if [ $? -ne 0 ]; then
        echo "Failed to process $proto_file"
        exit 1
    fi
done

echo "All .proto files have been processed."

git add .
git commit -m "Generated proto files for $NEW_TAG"

git tag "$NEW_TAG"

git push origin "$BRANCH"
git push origin "$NEW_TAG"

echo "Changes have been committed, tagged as $NEW_TAG, and pushed to the $BRANCH branch."
