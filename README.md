[![Build Status](https://travis-ci.org/ekomobile/consulsd.svg)](https://travis-ci.org/ekomobile/consulsd)
[![GitHub release](https://img.shields.io/github/release/ekomobile/consulsd.svg)](https://github.com/ekomobile/consulsd/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/ekomobile/consulsd)](https://goreportcard.com/report/github.com/ekomobile/consulsd)
![Downloads](https://img.shields.io/github/downloads/ekomobile/consulsd/total.svg)
[![GoDoc](https://godoc.org/github.com/ekomobile/consulsd?status.svg)](https://godoc.org/github.com/ekomobile/consulsd)

# ConsulSD
Service discovery Consul wrapper

# Usage

```go
    // init service discovery
    sd, err := consulsd.NewServiceDiscovery()
    if err != nil {
        // ...
    }
    
    // get service
    service, err := sd.Get("awesome-service")
    if err != nil {
        // ...
    }

    address := service.Address()
    port := service.Port()
```
