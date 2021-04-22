#! /bin/bash

# Usage: sh notify.sh

# Stop script on first error
set -e

REPO=$(echo $GITHUB_CONTEXT | jq -r '.repository')
TAGVERSION=$(echo $GITHUB_CONTEXT | jq -r '.event.release.tag_name')
TAGURL=$(echo $GITHUB_CONTEXT | jq -r '.event.release.html_url')
BODY=$(echo $GITHUB_CONTEXT | jq -r '.event.release.body' | sed -E ':a;N;$!ba;s/\r{0,1}\n/\\n/g')
echo "{\"username\":\"Cloud Bot Notify\",\"icon_url\":\"https://www.mattermost.org/wp-content/uploads/2016/04/icon.png\",\"text\":\"# **New Release for $REPO** - Release [$TAGVERSION]($TAGURL)\n '$BODY'\"}" > mattermost.json
cat mattermost.json