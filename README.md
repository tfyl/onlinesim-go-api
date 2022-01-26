# Onlinesim GO API

## (Note: I would not recommend them as they rate limit very early on and the only solution is using proxies)

Wrapper for automatic reception of SMS-messages by onlinesim.ru

## Installation

Require this package in your `package.json` or install it by running:
```bash
go get github.com/tfyl/onlinesim-go-api
```

### Example
```go
package main

import (
    "github.com/tfyl/onlinesim-go-api"
)

func main() {
    client := onlinesim.NewClient("", "en", "").Numbers()
    
    error, data := client.Get("vkcom", 7)
    if error != nil {
        panic(error)
    }

    println("end")
    println(fmt.Sprintf("%+v\n", data))
}
```

## Documentation

All documentation is in the wiki of this project - **[Documentation](https://github.com/tfyl/onlinesim-go-api/wiki)**

## Bugs

If you have any problems, please create Issues [here](https://github.com/tfyl/onlinesim-go-api/issues)   
