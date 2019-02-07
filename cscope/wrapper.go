package cscope

import (
	"bytes"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
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

// Systemwide cscope installation check
func cscopeCheck() error {
	binary, lookErr := exec.LookPath("cscope")
	if lookErr != nil {
		return errors.New("Cscope couldn't be detected on the system")
	} else {
		log.Infof("Cscope detected in %s", binary)
	}
	return nil
}

// Check db is sane
func dbCheck(path string) error {
	cmd := exec.Command("cscope", "-d", "-f", path, "-L")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	slurp, _ := ioutil.ReadAll(stderr)

	if err := cmd.Wait(); err != nil {
		return err
	}
	if string(slurp) != "" {
		return errors.New(string(slurp))
	}
	return nil
}

func NewCscope(path string) (Cscope, error) {
	if err := cscopeCheck(); err != nil {
		return Cscope{}, err
	}
	if err := dbCheck(path); err != nil {
		return Cscope{}, err
	}
	return Cscope{
		dbPath: path,
	}, nil
}
