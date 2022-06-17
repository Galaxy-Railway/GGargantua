protoc -I. ^
--go_out=. ^
--go_opt=module=github.com/Galaxy-Railway/GGargantua ^
    .\api\proto_files\*.proto