package apisix_sdk

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/lyszhang/go-lib/apisix/proto/pb"
)

func Init() {
	host := "http://10.20.111.110:80"
	key := "edd1c9f034335f136f87ad84b625c8f1"
	NewApiSixClient(host, key)
}

func TestApisixClient_ListRoute(t *testing.T) {
	Init()
	ctx := context.Background()
	resp, err := GetApiSixClient().GetRoute().List(ctx, 2, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Node, resp.Count)
}

func TestRoute_CreateRoute(t *testing.T) {
	Init()
	ctx := context.Background()
	req := &pb.CreateRouteReq{
		Id:          "9999",
		Name:        "测试wqwqw",
		Desc:        "测试wqwqw",
		Uris:        []string{"/api/v1", "/test"},
		Hosts:       []string{"customer.chtw.com", "*.bar.com"},
		RemoteAddrs: []string{"127.0.0.0", "192.168.26.25"},
		UpstreamId:  "4",
	}
	resp, err := GetApiSixClient().GetRoute().Create(ctx, req, nil, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Node, resp.Action, resp.ErrorMsg)
}

func TestRoute_DeleteRoute(t *testing.T) {
	Init()
	ctx := context.Background()
	resp, err := GetApiSixClient().GetRoute().Delete(ctx, "1633250139")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Deleted, resp.Action, resp.Key)
}
