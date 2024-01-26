package daemon

import "encoding/json"

type RPCRequest struct {
	ID      int64                  `json:"id"`
	JSONRPC string                 `json:"jsonrpc"`
	Method  RPCMethod              `json:"method"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

type RPCResponse struct {
	ID     int64           `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetNonceParams struct {
	Address    string `json:"address"`
	Topoheight uint64 `json:"topoheight,omitempty"` // optional
}

type GetAddressParams struct {
	IntegratedData string `json:"integrated_data"`
}

type RPCMethod string

const (
	GetVersion    RPCMethod = "get_version"
	GetNetwork    RPCMethod = "get_network"
	GetNonce      RPCMethod = "get_nonce"
	GetTopoHeight RPCMethod = "get_topoheight"
	GetAddress    RPCMethod = "get_address"
)
