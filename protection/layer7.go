```go
package protection

import (
	"net/http"
)

type Layer7Protection struct {
	protectionEnabled bool
}

func NewLayer7Protection(config map[string]interface{}) *Layer7Protection {
	return &Layer7Protection{
		protectionEnabled: config["layer7Protection"].(bool),
	}
}

func (l *Layer7Protection) ProtectLayer7(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if l.protectionEnabled {
			// Implement Layer 7 protection logic here
		}
		next.ServeHTTP(w, r)
	})
}
```