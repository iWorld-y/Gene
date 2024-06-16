package Gene

import (
	"log"
	"time"
)

func Logger(ctx *Context) {
	start := time.Now()
	ctx.Next()
	log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(start))
}
