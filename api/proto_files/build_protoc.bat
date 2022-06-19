protoc -I . ^
 ^
--go_out=../../ ^
--go_opt=module=github.com/Galaxy-Railway/GGargantua ^
 ^
--go-grpc_out=../../ ^
--go-grpc_opt=module=github.com/Galaxy-Railway/GGargantua ^
--go-grpc_opt=require_unimplemented_servers=false ^
 ^
--grpc-gateway_out=../../ ^
--grpc-gateway_opt=module=github.com/Galaxy-Railway/GGargantua ^
--grpc-gateway_opt generate_unbound_methods=true ^
--grpc-gateway_opt logtostderr=true ^
    .\*.proto