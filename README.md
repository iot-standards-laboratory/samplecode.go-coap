# samplecode.go-coap

## Prerequirement

- go 실행환경 구축

## commands for execution

- 실행 명령어 

  ```bash
  go run coap.go
  ```

- Options

  - `-server`

    - Server를 실행함

      ex) `go run coap.go -server`

  - `-ip`

    - IP를 설정함. 

      ex) `go run coap.go -ip {ip_address}`

    - Default IP : 127.0.0.1

  - `-port`

    - PORT를 설정함. 

      ex) `go run coap.go -port {port_number}`

    - Default port : 5683 

# ref
- https://github.com/go-ocf/go-coap