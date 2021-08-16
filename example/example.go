package main

import "github.com/hxoreyer/keing"

func main() {
	k := keing.Init()
	k.GET("/hello", func(c *keing.Context) {
		c.String(200, "hello go")
	})
	k.Run(":8090")
}
