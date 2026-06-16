package middleware

// import (
// 	"sync"
// 	"time"

// 	"github.com/danielgtaylor/huma/v2"
// )

// type Client struct {
// 	Count     int
// 	ExpiresAt time.Time
// }

// var (
// 	clients = make(map[string]*Client)
// 	mu      sync.Mutex
// )

// func RateLimit(maxRequests int, window time.Duration) func(huma.Context, func(huma.Context)) {

// 	return func(ctx huma.Context, next func(huma.Context)) {

// 		ip := ctx.RemoteAddr()

// 		mu.Lock()
// 		defer mu.Unlock()

// 		client, exists := clients[ip]

// 		if !exists || time.Now().After(client.ExpiresAt) {

// 			clients[ip] = &Client{
// 				Count:     1,
// 				ExpiresAt: time.Now().Add(window),
// 			}

// 			next(ctx)
// 			return
// 		}

// 		if client.Count >= maxRequests {

// 			huma.WriteErr(
// 				// api
// 				// ctx
// 				429,
// 				"Too many requests",
// 			)

// 			return
// 		}

// 		client.Count++

// 		next(ctx)
// 	}
// }