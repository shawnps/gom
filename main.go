package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf(`Usage of %s:
 Tasks:
   gom build   [options]   : Build with vendor packages
   gom install [options]   : Install bundled packages into vendor directory
   gom test    [options]   : Run tests with bundles
   gom run     [options]   : Run go file with bundles
   gom doc     [options]   : Run godoc for bundles
   gom exec    [arguments] : Execute command with bundle environment
   gom gen travis-yml      : Generate .travis.yml which uses "gom test"
   gom gen gomfile         : Scan packages from current directory as root
                              recursively, and generate Gomfile
`, os.Args[0])
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
	}
	handleSignal()

	var err error
	subArgs := flag.Args()[1:]
	switch flag.Arg(0) {
	case "install", "i":
		err = install(subArgs)
	case "build", "b":
		err = run(append([]string{"go", "build"}, subArgs...), None)
	case "test", "t":
		err = run(append([]string{"go", "test"}, subArgs...), None)
	case "run", "r":
		err = run(append([]string{"go", "run"}, subArgs...), None)
	case "doc", "d":
		err = run(append([]string{"godoc"}, subArgs...), None)
	case "exec", "e":
		err = run(subArgs, None)
	case "gen", "g":
		switch flag.Arg(1) {
		case "travis-yml":
			err = genTravisYml()
		case "gomfile":
			err = genGomfile()
		default:
			usage()
		}
	default:
		usage()
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "gom: ", err)
		os.Exit(1)
	}
}
