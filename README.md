# Go Map Access Panic

This repository demonstrates a common error in Go programs: panicking due to accessing non-existent keys in maps.

The `bug.go` file shows a simple Go program that triggers a runtime panic by accessing a map key without checking if it exists. The `bugSolution.go` file provides a solution to prevent this error. 

**Common Error:** When a program attempts to access a key in a map that is not present, Go throws a runtime panic. 

**Solution:** Before accessing a key, ensure the key exists in the map using `value, ok := myMap[key]` and handle the case where `ok` is false.