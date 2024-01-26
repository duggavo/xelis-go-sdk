package daemon

import (
	"context"
	"net/http"
	netUrl "net/url"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
	"github.com/xelis-project/xelis-go-sdk/transport"
)

type RPC struct {
	Client   *jrpc2.Client
	Username string
	Password string
}

func NewRPC(url, user, pass string) (*RPC, error) {
	daemonUrl, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	t := transport.NewTransport(http.DefaultTransport, user, pass)

	httpClient := http.Client{
		Transport: &t,
	}

	channel := jhttp.NewChannel(daemonUrl.String(), &jhttp.ChannelOptions{
		Client: jhttp.HTTPClient(&httpClient),
	})

	rpcClient := jrpc2.NewClient(channel, nil)
	daemon := &RPC{
		Client:   rpcClient,
		Username: user,
		Password: pass,
	}

	return daemon, nil
}

func (d *RPC) GetVersion(ctx context.Context) (version string, err error) {
	err = d.Client.CallResult(ctx, string(GetVersion), nil, &version)
	return
}
func (d *RPC) GetNetwork(ctx context.Context) (network string, err error) {
	err = d.Client.CallResult(ctx, string(GetNetwork), nil, &network)
	return
}

/*
	func dumpjson(d any) string {
		v, _ := json.Marshal(d)
		return string(v)
	}

	func (d *RPC) GetNonce(ctx context.Context, params GetNonceParams) (nonce uint64, err error) {
		fmt.Println("get nonce params are", dumpjson(params))
		err = d.Client.CallResult(ctx, string(GetNonce), params, &nonce)
		return
	}
*/
func (d *RPC) GetTopoHeight(ctx context.Context) (topoHeight uint64, err error) {
	err = d.Client.CallResult(ctx, string(GetTopoHeight), nil, &topoHeight)
	return
}

func (d *RPC) GetAddress(ctx context.Context, params GetAddressParams) (address string, err error) {
	err = d.Client.CallResult(ctx, string(GetAddress), params, &address)
	return
}

// TODO: split_address

// TODO: rescan

// TODO: get_balance

// TODO: get_tracked_assets

// TODO: get_asset_precision

// TODO: get_transaction

// TODO: build_transaction

// TODO: list_transactions

// TODO: sign_data

// TODO: estimate_fees
