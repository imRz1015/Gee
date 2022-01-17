package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	handlerFunc := func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))

	}
	return handlerFunc
}
