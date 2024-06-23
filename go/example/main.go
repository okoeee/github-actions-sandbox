package main

import "fmt"

// build 時に ldflags でバージョンを埋め込む
var version string

func main() {
	fmt.Printf("Example %s\n", version)
}
