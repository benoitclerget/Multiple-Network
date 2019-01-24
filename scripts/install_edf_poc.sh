#!/bin/bash

echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "Install demo-sc "
echo
CHANNEL_NAME="$1"
DELAY="$2"
LANGUAGE="$3"
TIMEOUT="$4"
VERBOSE="$5"
VERSION="$6"

: ${CHANNEL_NAME:="mychannel"}
: ${DELAY:="3"}
: ${LANGUAGE:="node"}
: ${TIMEOUT:="10"}
: ${VERBOSE:="false"}
: ${VERSION:="1.0"}
LANGUAGE=`echo "$LANGUAGE" | tr [:upper:] [:lower:]`
COUNTER=1
MAX_RETRY=5

CC_SRC_PATH="github.com/chaincode/edf_poc/"
if [ "$LANGUAGE" = "node" ]; then
	CC_SRC_PATH="/opt/gopath/src/github.com/chaincode/edf_poc_smartcontract/"
fi

echo "Channel name : "$CHANNEL_NAME

# import utils
. scripts/utils.sh


## Install chaincode on peer0.org1 and peer0.org2
echo "Installing chaincode on peer0.org1..."
installEdfPocChaincode 0 1 $VERSION
echo "Install chaincode on peer0.org2..."
installEdfPocChaincode 0 2 $VERSION

# Instantiate chaincode on peer0.org2
echo "Instantiating chaincode on peer0.org1..."
instantiateEdfPocChaincode 0 1 $VERSION



echo
echo " _____   _   _   ____   "
echo "| ____| | \ | | |  _ \  "
echo "|  _|   |  \| | | | | | "
echo "| |___  | |\  | | |_| | "
echo "|_____| |_| \_| |____/  "
echo

exit 0
