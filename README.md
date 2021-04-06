# Go API client for aarch64.com

Example:

```go
package main

import (
	"log"

	"github.com/natesales/aarch64-client-go"
)

func main() {
	client := aarch64.Client{APIKey: "cfea63484ccfea63484bfe78ed72d2cbfe78eea63484bfe78e"}

	resp, err := client.Projects()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", resp)
}
```
