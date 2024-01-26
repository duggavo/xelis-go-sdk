package daemon

import (
	"context"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

func setupRPC(t *testing.T) (daemon *RPC, ctx context.Context) {
	ctx = context.Background()
	daemon, err := NewRPC(config.TESTNET_WALLET_RPC, "user", "pass")
	if err != nil {
		t.Fatal(err)
	}

	return
}
func TestRPCUnknownMethod(t *testing.T) {
	wallet, ctx := setupRPC(t)
	res, err := wallet.Client.Call(ctx, "UnknownMethod", nil)
	if err == nil {
		t.Fatal("Expected an error")
	}

	t.Log(res)
}
func TestRPCGetVersion(t *testing.T) {
	wallet, ctx := setupRPC(t)
	res, err := wallet.GetVersion(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
func TestRPCGetNetwork(t *testing.T) {
	wallet, ctx := setupRPC(t)
	res, err := wallet.GetNetwork(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}

var addr string

func TestRPCGetAddress(t *testing.T) {
	wallet, ctx := setupRPC(t)
	res, err := wallet.GetAddress(ctx, GetAddressParams{})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
	addr = res
}

/*
TODO: this somewhat does not work?
func TestRPCGetNonce(t *testing.T) {
	wallet, ctx := setupRPC(t)
	res, err := wallet.GetNonce(ctx, GetNonceParams{
		Address: addr,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}*/
