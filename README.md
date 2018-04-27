# sd
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
