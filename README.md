# Go API client for aarch64.com

Example:

```go
package main

import (
	"fmt"
	"github.com/natesales/aarch64-client-go/pkg/aarch64"
	"log"
)

func main() {
	client := aarch64.Client{APIKey: "cfea63484ccfea63484bfe78ed72d2cbfe78eea63484bfe78e"}

	resp, err := client.Projects()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", resp)
}

```