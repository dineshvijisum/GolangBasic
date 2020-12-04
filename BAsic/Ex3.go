package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
)

func main() {
	// Create and seed the generator.
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.
	r := rand.New(rand.NewSource(99))

	// The tabwriter here helps us generate aligned output.
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	show := func(name string, v1 interface{}) {
		fmt.Fprintf(w, "%s\t%v\t\n", name, v1)
	}

	// Float32 and Float64 values are in [0, 1).
	show("Float32", r.Float32())
	show("Float64", r.Float64())
}