#!/bin/bash

protoc calc_pb/calc.proto --go_out=plugins=grpc:.