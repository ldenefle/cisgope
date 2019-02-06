package main

import (
	cscope "github.com/ldenefle/cisgope/cscope"
	gui "github.com/ldenefle/cisgope/gui"
	log "github.com/sirupsen/logrus"
	pin "gopkg.in/alecthomas/kingpin.v2"
)

const dbName = "cscope/cscope.out"

func main() {
	pin.Version("cisgope version 0.1.0")
	pin.CommandLine.HelpFlag.Short('h')
	pin.CommandLine.VersionFlag.Short('v')
	pin.Parse()
	log.SetLevel(log.DebugLevel)
	db, err := cscope.NewCscope(dbName)
	if err != nil {
		log.Error(err)
	}
	gui.Display(db)
	return
}
