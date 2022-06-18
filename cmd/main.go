package main

import (
	"ca-library-app/internal/composites"
	"fmt"
)

func main() {
	// entry point
	authorComposite, err := composites.NewAuthorComposite()
	if err != nil {
	}

	bookComposite, err := composites.NewBookComposite()
	if err != nil {
	}

	// register in router
	fmt.Println(authorComposite, bookComposite)
}
