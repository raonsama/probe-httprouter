// Copyright 2024 Raymond Ma'ruf <raonsama@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.ListenAndServe(":8000", mux)
}
