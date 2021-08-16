package keing

import (
	"fmt"
	"time"
)

func Logger() HandlerFunc{
	return func(c *Context) {
		t := time.Now()
		c.Next()
		ts := fmt.Sprintf("%d/%d/%d - %d:%d:%d",time.Now().Year(),time.Now().Month(),time.Now().Day(),time.Now().Hour(),time.Now().Minute(),time.Now().Second())
		fmt.Printf("[Keing] %-22s| %3d | %15s | %16s | %6s  \"%-s\"\n",ts,c.code,time.Since(t),c.Request.RemoteAddr,c.Request.Method,c.Request.RequestURI)
		if len(c.Errors) > 0 {
			fmt.Println(c.Errors)
		}
	}
}
