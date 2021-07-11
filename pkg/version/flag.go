package version

import (
	"flag"
	"fmt"
	"os"
)

var (
	showVersion       = flag.Bool("version", false, "show version")
	showVersionDetail = flag.Bool("dversion", false, "show detail version")
)

func CmdParseVersion() {
	if *showVersion {
		fmt.Println(GetVersion())
		os.Exit(0)
	}
	if *showVersionDetail {
		fmt.Println(DetailVersion())
		os.Exit(0)
	}
}
