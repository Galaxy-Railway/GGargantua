protoc -I. ^
--go_out=plugins=grpc:. ^
    .\api\proto_files\health_check.proto ^
    .\api\proto_files\http_common.proto ^
    .\api\proto_files\http_content.proto ^
    .\api\proto_files\https_content.proto ^
    .\api\proto_files\request_step.proto ^
    .\api\proto_files\step.proto ^
    .\api\proto_files\job.proto ^
    .\api\proto_files\common.proto ^