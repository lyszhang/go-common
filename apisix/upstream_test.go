package apisix_sdk

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/lyszhang/go-lib/apisix/proto/pb"
	"github.com/satori/go.uuid"
)

func TestUpstream_List(t *testing.T) {
	Init()
	ctx := context.Background()
	resp, err := GetApiSixClient().GetUpstream().List(ctx, 2, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Action, resp.Count, resp.Node)
}

func TestCreateUpstream(t *testing.T) {
	u1 := uuid.NewV4()
	uid := fmt.Sprintf("%s", u1)

	ctx := context.Background()
	NewApiSixClient("http://uat.apisix-api.lls.com", "VPUDfhonjj3oqRN0uFfW3FugNBxEfp")

	resp, err := GetApiSixClient().GetUpstream().Create(ctx, &pb.CreateUpstreamReq{
		Id:    uid,
		Name:  "podName-wrerwerertertret",
		Nodes: map[string]int32{"127.0.0.1:80": 1},
		Type:  pb.UpstreamRoundrobinType,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func TestUpstream_Create(t *testing.T) {
	Init()
	ctx := context.Background()
	resp, err := GetApiSixClient().GetUpstream().Create(ctx, &pb.CreateUpstreamReq{
		Name:  "测试",
		Desc:  "cesi",
		Nodes: map[string]int32{"127.0.0.1:80": 1, "127.0.0.1:81": 1},
		Type:  pb.UpstreamRoundrobinType,
		Id:    "4",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Node, resp.Action, resp.Node)
}

func TestUpstream_Detele(t *testing.T) {
	Init()
	ctx := context.Background()
	resp, err := GetApiSixClient().GetUpstream().Delete(ctx, "1633250139")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Deleted, resp.Action)
}
