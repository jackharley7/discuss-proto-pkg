VERSION = "unset"
DATE=$(shell date -u +%Y-%m-%d-%H:%M:%S-%Z)
DOCKER_IMAGE ?=	jackharley7/api-gateway
IMAGE_TAG ?= latest

# .PHONY: codegen
# codegen:
# 	GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
# 	GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
# 	go install $(shell go list -f '{{ .Dir }}' -m github.com/golang/protobuf)/protoc-gen-go
# 	protoc \
# 		--proto_path=${GOPATH}/src \
# 		--proto_path=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
# 		--grpc-gateway_out=logtostderr=true:. \
# 		--swagger_out=logtostderr=true:. \
# 		--go_out=plugins=grpc:. -I . proto/conversation.proto proto/link.proto proto/invitation.proto \
# 	gofmt -w proto

copyprotos:
	cp ../conversationservice/proto/*.proto /

codegen:
	GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install $(shell go list -f '{{ .Dir }}' -m github.com/golang/protobuf)/protoc-gen-go
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		proto/conversation.proto proto/link.proto proto/invitation.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		proto/notification.proto proto/counts.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		proto/profile.proto proto/user.proto proto/userTemp.proto proto/education.proto proto/workexperience.proto

