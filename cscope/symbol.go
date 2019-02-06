package cscope

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type Symbol struct {
	File        string
	Name        string
	LineNumber  int
	LineContent string
}

func (s Symbol) String() string {
	return s.File + ":" + strconv.Itoa(s.LineNumber) + ":  " + s.Name
}

func parseSymbols(reader io.Reader) ([]Symbol, error) {
	symbols := make([]Symbol, 0)
	scanner := bufio.NewScanner(reader)
	var err error
	var lineNumber int
	for scanner.Scan() {
		s := strings.SplitN(scanner.Text(), " ", 4)
		if len(s) < 4 {
			return nil, errors.New("Couldn't parse")
		}
		if lineNumber, err = strconv.Atoi(s[2]); err != nil {
			return nil, err
		}
		symbols = append(symbols, Symbol{
			File:        s[0],
			Name:        s[1],
			LineNumber:  lineNumber,
			LineContent: s[3],
		})
	}
	if len(symbols) == 0 {
		return nil, errors.New("Couldn't find symbol")
	}
	return symbols, nil
}
