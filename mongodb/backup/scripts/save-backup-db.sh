#!/bin/bash

SCRIPTS_SRC=./mongodb/backup/scripts

# parse configs
. ${SCRIPTS_SRC}/default.config

source ${SCRIPTS_SRC}/bash-utils.sh

# Test
echo ${RED}RED${GREEN}GREEN${YELLOW}YELLOW${BLUE}BLUE${PURPLE}PURPLE${CYAN}CYAN${WHITE}WHITE${RESTORE}

echo "${YELLOW}Default host${RESTORE}: ${GREEN}$DEFAULT_HOST${RESTORE}"
echo "${YELLOW}Default port${RESTORE}: ${GREEN}$DEFAULT_PORT${RESTORE}"
echo "${YELLOW}Default username${RESTORE}: ${GREEN}$DEFAULT_USER${RESTORE}"
echo "${YELLOW}Default password${RESTORE}: ${GREEN}$DEFAULT_PASS${RESTORE}"
echo "${YELLOW}Default database${RESTORE}: ${GREEN}$DEFAULT_DB${RESTORE}"

docker exec -t itv-mongo mongodump --out $MONGO_PATH/ --host ${DEFAULT_HOST} --port ${DEFAULT_PORT} --db ${DEFAULT_DB} --username ${DEFAULT_USER} --password ${DEFAULT_PASS} --authenticationDatabase admin