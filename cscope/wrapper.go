package cscope

import (
	"bytes"
	"errors"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

var dbPath string

var cmds = map[int]string{
	0: "-L0",
	2: "-L2",
	3: "-L3",
	4: "-L4",
}

type Cscope struct {
	dbPath string
}

func (cscope Cscope) Cmd(command int, symbol string) ([]Symbol, error) {
	cmd := cmds[command]
	if cmd == "" {
		return nil, errors.New("Unsupported command")
	}
	out, err := exec.Command("cscope", "-d", "-f", cscope.dbPath, cmd, symbol).Output()
	if err != nil {
		return nil, err
	}
	return parseSymbols(bytes.NewReader(out))
}

func NewCscope(path string) (Cscope, error) {
	binary, lookErr := exec.LookPath("cscope")
	if lookErr != nil {
		return Cscope{}, errors.New("Cscope couldn't be detected on the system")
	} else {
		log.Infof("Cscope detected in %s", binary)
	}
	return Cscope{
		dbPath: path,
	}, nil
}
