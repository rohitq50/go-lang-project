package utility

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

// GetLoggedInUserID returns current loggedIn user id
// it reads userId from request context and
// return id as string.
func GetLoggedInUserID(r *http.Request) string {
	id := context.Get(r, "userId")
	return fmt.Sprintf("%v", id)
}
