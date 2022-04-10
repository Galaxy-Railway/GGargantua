package client

import (
	"context"
	"encoding/json"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf/request_pb"
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/request/http"
	"testing"
)

func TestSendSingleHttp(t *testing.T) {
	httpRequest := new(http.RequestContent)
	httpRequest.Url = "http://www.baidu.com"
	httpRequest.Method = "GET"
	bytes, err := json.Marshal(httpRequest)
	if err != nil {
		t.Fatalf("marshal struct into json failed, err: %v", err)
	}

	singleRequest := new(request_pb.SingleRequest)
	singleRequest.RequestSchema = request_pb.SchemaType_HTTP
	singleRequest.RequestContent = bytes

	conn, err := GetConnection()
	if err != nil {
		t.Fatalf("get connection failed, err: %v", err)
	}

	c := request_pb.NewSingleRequestSenderClient(conn)
	r, err := c.SendSingleRequest(context.Background(), singleRequest)
	if err != nil {
		t.Errorf("failed to execute rpc, err: %v", err)
	}
	t.Logf("send rpc success")
	t.Log(string(r.ResponseContent))
}
