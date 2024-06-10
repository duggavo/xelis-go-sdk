package xswd

import (
	"fmt"
	"log"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

func setupWebSocket(t *testing.T) (xswd *XSWD) {
	xswd, err := NewXSWD(config.LOCAL_XSWD_WS)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestWSAuthError(t *testing.T) {
	xswd := setupWebSocket(t)

	go func() {
		err := <-xswd.WS.ConnectionErr
		fmt.Println(err)
	}()

	_, err := xswd.Authorize(ApplicationData{
		ID:          "ertherth",
		Name:        "Test App",
		Description: "This is a test app.",
		Permissions: make(map[string]Permission),
	})

	if err.Error() != "Invalid application ID" {
		t.Fail()
	}
}

func TestWSGetInfo(t *testing.T) {
	xswd := setupWebSocket(t)

	res, err := xswd.Authorize(ApplicationData{
		ID:          "9F86D081884C7D659A2FEAA0C55AD015A3BF4F1B2B0B822CD15D6C15B0F00A08",
		Name:        "Test App",
		Description: "This is a test app.",
		Permissions: make(map[string]Permission),
	})
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("%+v", res)

	info, err := xswd.Daemon.GetInfo()
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("%+v", info)

	version, err := xswd.Wallet.GetVersion()
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("%s", version)

	xswd.Close()
}
