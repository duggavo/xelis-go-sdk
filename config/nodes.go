package config

import "fmt"

const NODE_URL = "node.xelis.io"
const TESTNET_NODE_URL = "testnet-node.xelis.io"

var NODE_RPC = fmt.Sprintf("https://%s/json_rpc", NODE_URL)
var TESTNET_NODE_RPC = fmt.Sprintf("https://%s/json_rpc", TESTNET_NODE_URL)

var NODE_WS = fmt.Sprintf("wss://%s/json_rpc", NODE_URL)
var TESTNET_NODE_WS = fmt.Sprintf("wss://%s/json_rpc", TESTNET_NODE_URL)

// to start wallet RPC for test purposes:
// start_rpc_server 127.0.0.1:4121 user pass

var WALLET_RPC = "http://127.0.0.1:4121/json_rpc"
var TESTNET_WALLET_RPC = WALLET_RPC
