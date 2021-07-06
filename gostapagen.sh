#!/bin/sh

echo "$(<LICENSE)"

# Configure your stapagen app below:
APP_PORT=8080
CONTENT_DELIMITER="b-nova-content-header"
REPO_URL="https://github.com/b-nova-openhub/jams-vanilla-content"
REPO_BRANCH="main"
REPO_CLONE_PATH="/tmp"
REPO_CONTENT_DIR="/content"

./bin/gostapagen serve --appPort=$APP_PORT --contentDelimiter=$CONTENT_DELIMITER --repoUrl=$REPO_URL --repoBranch=$REPO_BRANCH --repoClonePath=$REPO_CLONE_PATH --repoContentDir=$REPO_CONTENT_DIR
