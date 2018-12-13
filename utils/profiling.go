package utils

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func ProfileCPU() func() {
	fmt.Println("Starting profiling")
	f, err := os.Create("profile.cpu")
	if err != nil {
		panic(err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}

	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}
}
