#!/usr/bin/env bash

protoDir="../protos"
outDir="../protos"
protoc -I ${protoDir}/ ${protoDir}/*proto --go_out=plugins=grpc:${outDir}
