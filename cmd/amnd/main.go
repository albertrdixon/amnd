package main

import (
	"os"

	"github.com/albertrdixon/amnd"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	logLevels = []string{"fatal", "error", "warn", "info", "debug"}
	app       = kingpin.New("amnd", "Plex updater daemon").Version(amnd.Version())

	cf = app.Flag("config-file", "Config file path").Short('C').PlaceHolder("/path/to/config").Default("/etc/amnd.yml").OverrideDefaultFromEnvar("AMND_CONFIG").ExistingFile()
	in = app.Flag("update-interval", "Update interval").Short('t').Default("24h").OverrideDefaultFromEnvar("AMND_INTERVAL").Duration()
	tm = app.Flag("tmpdir", "Temp dir").Default("/tmp").OverrideDefaultFromEnvar("AMND_TMP").ExistingDir()
	lv = app.Flag("log-level", "Logging level. One of: debug, info, warn, error, fatal").Short('l').Default("info").OverrideDefaultFromEnvar("AMND_LOG_LEVEL").Enum(logLevels...)
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	c, e := amnd.ReadConfig(*cf)
	if e != nil {
		panic(e)
	}
	spew.Dump(c)
}
