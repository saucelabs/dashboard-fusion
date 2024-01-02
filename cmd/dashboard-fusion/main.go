// Copyright 2023 Sauce Labs Inc., all rights reserved.

package main

import (
	"encoding/json"
	"log"
	"os"

	fusion "github.com/saucelabs/dashboard-fusion"
	"github.com/spf13/pflag"
)

var args = struct {
	dash   *string
	panels *[]string
	out    *string
}{
	dash:   pflag.String("dash", "", "Location of base dashboard [required]"),
	panels: pflag.StringSlice("panels", []string{}, "Location of panel(s) to be merged into base dashboard [required]"),
	out:    pflag.String("out", "", "Location of updated dashboard, defaults to stdout"),
}

func main() {
	if !pflag.Parsed() {
		pflag.Parse()
	}
	if *args.dash == "" || len(*args.panels) == 0 {
		pflag.Usage()
		return
	}

	d, err := readFromFile[fusion.Dashboard](*args.dash)
	if err != nil {
		log.Fatal("reading dashboard ", err)
	}

	ps := d.Panels()
	for i := range *args.panels {
		ps2, err := readFromFile[[]fusion.Panel]((*args.panels)[i])
		if err != nil {
			dd, err2 := readFromFile[fusion.Dashboard]((*args.panels)[i])
			if err2 != nil {
				log.Fatal("reading panels ", err, err2)
			}
			ps2 = dd.Panels()
		}

		ps = fusion.MergePanels(ps, ps2)
	}

	d["panels"], err = json.Marshal(ps)
	if err != nil {
		log.Fatal("marshalling merged panels ", err)
	}

	var out *os.File
	if *args.out != "" {
		var err error
		out, err = os.Create(*args.out)
		if err != nil {
			log.Fatal("creating output dashboard ", err)
		}
		defer func() {
			if err := out.Close(); err != nil {
				panic(err)
			}
		}()
	} else {
		out = os.Stdout
	}

	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	if err := enc.Encode(d); err != nil {
		log.Println("encoding output dashboard ", err)
	}
}

func readFromFile[T any](filename string) (T, error) {
	var obj T

	f, err := os.Open(filename)
	if err != nil {
		return obj, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	if err := json.NewDecoder(f).Decode(&obj); err != nil {
		return obj, err
	}

	return obj, nil
}
