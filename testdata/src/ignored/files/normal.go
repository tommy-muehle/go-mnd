package files

import (
	"net/http"
)

func example3() {
	http.StatusText(200) // want "Magic number: 200"
}
