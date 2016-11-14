package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	dir := flag.String("dir", "/tmp/bench", "Benchmark dir")
	flag.Parse()

	err := os.MkdirAll(*dir, 0775)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if len(os.Args) == 1 {
		log.Fatalf("<name> required")
	}

	old := filepath.Join(*dir, os.Args[1])

	if len(os.Args) > 2 {
		new := filepath.Join(*dir, os.Args[2])
		cmd := exec.Command("benchcmp", old, new)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("error: %s", err)
			return
		}
		return
	}

	cmd := exec.Command("go", "test", "-run=NONE", "-bench=.", "./...")
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	f, err := os.Create(old)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if _, err := io.Copy(f, bytes.NewReader(b)); err != nil {
		log.Fatalf("error: %s", err)
	}

	if _, err := io.Copy(os.Stdout, bytes.NewReader(b)); err != nil {
		log.Fatalf("error: %s", err)
	}
}
