package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"arhat.dev/template-application-go/cmd/template-application-go/pkg"
	"arhat.dev/template-application-go/pkg/version"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cmd := pkg.NewTemplateApplicationGoCmd()
	cmd.AddCommand(version.NewVersionCmd())

	err := cmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to run template-application-go %v: %v\n", os.Args, err)
		os.Exit(1)
	}
}
