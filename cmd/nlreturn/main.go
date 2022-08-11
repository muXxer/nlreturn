package main

import (
	"github.com/muXxer/nlreturn/v2"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(nlreturn.NewAnalyzer())
}
