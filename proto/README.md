# proto

Protobuf API Specification

## Specificatios

[spec.md](./spec.md)

## Requirements

- [Buf](https://docs.buf.build/introduction)
- [Connect](https://connect.build/docs/go/getting-started)

You can install dependencies:
```shell
make deps
```

## Write API Specification

### Package Structure

```shell
$(pwd)                                                                                                        
 ├──buf.gen.yaml # Code generator config                                                                                                                                        
 ├──buf.work.yaml # Module management                                                                                                                                                                                                                                                                                   
 ├──oidcapis # Create package per microservice                                                                                                                                            
 │  ├──buf.yaml                                                                                                                                         
 │  └──oidc # One microservice may contains multiple innter packages                                                                                                                                             
 │     └──v1                                                                                                                                            
 │        └──oidc.proto                                                                                                                                 
 └──README.md   
```

## Generate Client Code

```shell
make generate
```

## Developers Command

### Lint

```shell
make lint
```

### Format

```shell
make format
```