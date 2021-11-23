// Package config parse and validates command flags.
package config

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

// Config represents the current source and target config.
// Source are root folder of encrypted files.
// Target are folder path where decrypted file is writen.
type Config struct {
	Source      string
	Target      string
	Concurrency int
	Quiet       bool
}

// exit will exit and print the usage.
// Used in case of errors during flags parse/validate.
func exit(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	os.Exit(1)
}

// validate makes sure from and to are Redis URIs or file paths,
// and generates the final Config.
func validate(from string, to string, concurrency int, quiet bool) (Config, error) {
	cfg := Config{
		Source:      from,
		Target:      to,
		Concurrency: concurrency,
		Quiet:       quiet,
	}

	switch {
	case cfg.Source == "":
		return cfg, fmt.Errorf("from is required")
	case cfg.Target == "":
		return cfg, fmt.Errorf("to is required")
	case cfg.Concurrency < 1:
		cfg.Concurrency = runtime.NumCPU()
	}

	return cfg, nil
}

// Parse parses the command line flags and returns a Config.
func Parse() Config {
	from := flag.String("f", "", "path to encrypted folder")
	to := flag.String("t", "", "path where decrypted file will be writen to")
	concurrency := flag.Int("n", 0, "number of worker, default is number of cpu cores")
	quiet := flag.Bool("q", false, "quiet mode, default is false")

	flag.Parse()

	cfg, err := validate(*from, *to, *concurrency, *quiet)
	if err != nil {
		// we exit here instead of returning so that we can show
		// the usage examples in case of an error.
		exit(err)
	}

	return cfg
}
