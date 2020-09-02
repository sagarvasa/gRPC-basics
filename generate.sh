#!/bin/bash

protoc calc_pb/calc.proto --go_out=plugins=grpc:.
protoc error_pb/err.proto --go_out=plugins=grpc:.