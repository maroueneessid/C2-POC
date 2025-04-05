all:
	make clean
	make flush_redis
	make proto_build
	make cert_gen
	make c2s
	make c2s_manager
	make asset_agent


proto_build:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative ./proto_defs/common/common.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative ./proto_defs/manager/manager.proto

flush_redis:
	redis-cli --scan --pattern '*' | xargs redis-cli DEL

asset_agent:
	go build -o asset/bin/asset.elf asset/*.go
	GOOS=windows GOARCH=amd64 go build -o asset/bin/asset.exe  asset/*.go

c2s:
	go build -o server/bin/server.elf server/*.go

c2s_manager:
	go build -o manager/bin/manager.elf manager/*.go

cert_gen :
	chmod +x ./utils/cert/gen.sh 
	./utils/cert/gen.sh

.IGNORE clean :
	rm */bin/*
	rm -rf ~/.customC2/*
	rm ./utils/cert/server/*
