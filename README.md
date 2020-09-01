# gRPC-basics
1. gRPC is free and open source framework developed by google.
2. At high level, gRPC allow us to define REQUEST & RESPONSE for RPC (Remote Procedure Calls) and handle all the rest for us.
3. On top of it, 
    a) it's modern, fast and efficient 
    b) build on a top of HTTP/2
    c) low latency, supports streaming, language independent and makes it super easy to plug in authentication, load balancing, logging & monitoring.
4. At the core of gRPC we need to define messages and services using Protocol Buffers
(One .proto file works over multiple programming language (including client & server) and allow to use a framework that scales to millions of RPC per second)

# Types of API in gRPC
1. Unary (like traditional API)
2. Server Streaming
3. Client Streaming
4. Bi Directional Streaming

# Protocol Buffers
1. protocol buffers are language agnostic (Code can be generated for pretty much any language)
2. Data is binary and efficienty serialized (small paylaod)
3. Very convenient for transporting lot of data
4. Protocol buffers are used to define messages (data, req, res) and services (service name and rpc endpoints). We then generate code from it
5. protocol buffers defines rules to make API evolve without breaking existing client which is useful in microservices

# Protocol Buffers Vs JSON
1. size of protocol buffer is lesser than JSON hence we save lot of network bandwidth
2. Parsing JSON is CPU intensive since format is human readable and not machine readable. While protobufs are in binary format making it close how machine represent data

# HTTP/2 vs HTTP/1.1
1. HTTP/1.1 opens new tcp connection to server at each request. It doesn't compress header and it's basically request/response format (no server push)

e.g: If average webpage loads 50 assets, then http/1.1 makes 50 calls with plain text header with each request.
This increase latency and network packet size

2. HTTP/2 supports multiplexing (this reduce latency). It also allows server push (Server can push stream for one request from client). It allow header compressions
HTTP/2 is binary and by default SSL enabled (Secured)

# gRPC vs REST
1. gRPC  - 
    Smaller, faster 
    HTTP/2
    Bi Directional, Async, Unary
    Stream support
    Code generation through protobufs - 1st Class citizen
    It is API oriented

2. REST -
    text based, Bigger in size, Slower
    HTTP/1.1
    Only Client to Server
    No stream support. Only request/response
    Code generation through open API/Swagger - 2nd Class citizen
    It is resource oriented (GET/POST/PUT/DELETE)

# Project Setup
1. install golang from `https://golang.org/doc/install` or using homebrew (on mac)
2. make sure go binaries are on path. Check version using `go version`
3. In order to perform code generation, we need to install `protoc` on system. Use `brew install protobuf` for mac
