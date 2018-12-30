# Micro Framework demo

## Objective: Show a demo with REST GRPCGateway; connecting to multiple backend deps; connecting via multiple material; still have to unit test

## Scenario: 
- Show a bee shop web demo connecting to a user api; which has deps to a bee vee service.
- A bee catalog service also powers via cli + api; and has deps to the user api


## HOWTO Compile Proto

```
compile the proto file example.proto:

cd /Users/leow/go/src/github.com/leowmjw/playground-golang/MICRO/beeWax
protoc --proto_path=. --go_out=. --micro_out=. proto/example/example.proto
```
