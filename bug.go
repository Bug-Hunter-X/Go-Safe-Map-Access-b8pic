```go
func main() {
    m := make(map[string]int)
    m["a"] = 1
    fmt.Println(m["b"]) // Accessing a non-existent key
}
```