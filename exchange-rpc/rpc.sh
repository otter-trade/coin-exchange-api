#!/bin/bash

goctl rpc protoc pb/exchange.proto --style go_zero --go_out=. --go-grpc_out=. --zrpc_out=. --home="../common/tpl"