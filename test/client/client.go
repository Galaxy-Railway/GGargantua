package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Galaxy-Railway/GGargantua/api/protobuf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"strings"
)

func main() {
	// set logger
	atom := zap.NewAtomicLevel()
	originLogger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.Lock(os.Stdout),
		atom,
	))
	logger := originLogger.Sugar()

	// get grpc client
	ctx := context.Background()
	opts := make([]grpc.CallOption, 0)
	conn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalf("dail grpc server failed, %v", err)
	}
	client := protobuf.NewJobServiceClient(conn)

	// create job
	jobid, err := client.CreateAJob(ctx, &emptypb.Empty{}, opts...)
	if err != nil {
		logger.Fatalf("create job failed, %v", err)
	}
	getBaidu.Uuid.Uuid = jobid.Uuid

	// start job
	_, err = client.StartAJob(ctx, getBaidu, opts...)
	if err != nil {
		logger.Fatalf("start job failed, %v", err)
	}

	// get job result
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		text := scanner.Text()
		if strings.Compare("1", text) == 0 {
			result, err := client.GetJobResult(ctx, jobid, opts...)
			if err != nil {
				logger.Fatalf("get job result failed, %v", err)
			}
			pretty, err := json.MarshalIndent(result, "", "\t")
			if err != nil {
				logger.Fatalf("marshal indent failed, %v", err)
			}
			fmt.Println(string(pretty))
		} else if strings.Compare("2", text) == 0 {
			_, err := client.CancelAJob(ctx, jobid, opts...)
			if err != nil {
				logger.Fatalf("get job result failed, %v", err)
			}
			//fmt.Printf("result:/n%+v\n", result)
			logger.Infof("stoped job")
		} else {
			logger.Errorf("wrong input: %s", text)
			break
		}
	}
}

var (
	getBaidu = &protobuf.UpdateJobContent{
		Uuid: &protobuf.JobUuid{Uuid: ""},
		MainStep: &protobuf.Step{
			Type: protobuf.StepType_REQUEST,
			RequestStep: &protobuf.RequestStepType{
				Request: &protobuf.Request{
					RequestSchema: protobuf.SchemaType_HTTP,
					Times:         10,
					Concurrency:   1,
					HttpRequest: &protobuf.HttpRequest{
						Method:  "GET",
						Url:     "https://www.baidu.com",
						Headers: nil,
						Params:  nil,
						Body:    nil,
					},
					HttpsRequest: nil,
				},
			},
			ScriptStep:   nil,
			ForStep:      nil,
			IfStep:       nil,
			SequenceStep: nil,
		},
	}
)
