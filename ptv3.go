package main

import (
	"ptv3/internal/route"
)

// entry point
func main() {
	r := route.Route()
	r.Run(":8080")
}
