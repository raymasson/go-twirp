build:
	protoc --proto_path=./proto --twirp_out=./users --go_out=./users ./proto/users.proto