MICRO_REGISTRY:=mdns
protoSrc = $(wildcard user/proto/example/*.proto) \
					$(wildcard beeCatalog/proto/example/example.proto) \
					$(wildcard beeVee/proto/example/*.proto) \
					$(wildcard beeWax/proto/example/*.proto)
SRCS = $(wildcard */proto/*/*.proto)
PBS = $(SRCS:.proto=.pb.go)

default: proto

proto: 
	@cd beeWax/proto/example/ && protoc --proto_path=. --go_out=. --micro_out=. *.proto

$(PBS): $(SRCS)
	@echo "Change in pb.go?"
	@protoc --proto_path=. --go_out=. --micro_out=. $@

test:
	@echo "Run tests .."

build: 
	@echo "Run build .. $(SRCS)"

deps:
	@echo "Get deps .."

tools:
	@echo "Acquire tools .."

list: 
	@echo "Listing services registered!"
	@MICRO_REGISTRY=$(MICRO_REGISTRY) micro list services
	@# Above for local; below for consul in a cluster?
	@# @micro list services

.PHONY: deps tools test build proto
