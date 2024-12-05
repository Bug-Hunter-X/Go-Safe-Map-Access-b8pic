```go
package main

import (

    "fmt"
)

func main() {
    m := make(map[string]int)
    m["a"] = 1
    value, ok := m["b"] // Check for key existence
    if ok {
        fmt.Println(value) 
    } else {
        fmt.Println("Key not found") // Handle the case where the key is missing
    }
}
```