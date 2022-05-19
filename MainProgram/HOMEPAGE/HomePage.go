package HOMEPAGE

import (
	"fmt"
	"net/http"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Wellcome homepage!")
}
