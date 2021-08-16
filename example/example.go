package main

import keing "github.com/hxoreyer/keing"

func main() {
	k := keing.Init()
	k.GET("/hello/:name", func(c *keing.Context) {
		c.Json(200, keing.K{
			"msg": "ni hao " + c.Name("name"),
		})
	})
	k.Run(":8090")
}
