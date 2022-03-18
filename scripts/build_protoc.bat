protoc -I. --go_out=plugins=grpc:. ^
    .\api\proto_files\health_check.proto ^
    .\api\proto_files\single_request_sender.proto ^
    .\api\proto_files\http_content.proto