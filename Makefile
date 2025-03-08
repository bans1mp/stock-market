PROTOC = protoc
GRPC_CPP_PLUGIN = grpc_cpp_plugin

compile_protos: 
	make compile_go_protos
	make compile_cpp_protos

compile_go_protos:
	cd backend && protoc --proto_path=../proto --go_out=. --go-grpc_out=. ../proto/trade.proto

compile_cpp_protos:
	cd matching-engine && protoc --proto_path=../proto --cpp_out=./proto --grpc_out=./proto --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` ../proto/trade.proto
