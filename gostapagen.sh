#!/bin/sh

echo "$(<LICENSE)"

# Configure your stapagen app below:
APP_PORT=8080
CONTENT_DELIMITER="b-nova-content-header"
REPO_URL="https://github.com/b-nova-openhub/jams-vanilla-content"
REPO_BRANCH="main"
REPO_CLONE_PATH="/tmp"
REPO_CONTENT_DIR="/content"

./bin/stapagen --port=$APP_PORT --delimiter=$CONTENT_DELIMITER --repo=$REPO_URL --branch=$REPO_BRANCH --clonePath=$REPO_CLONE_PATH --contentDir=$REPO_CONTENT_DIR
