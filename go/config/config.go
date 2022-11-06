package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type HostPort struct {
	Host string
	Port string
}

type Config struct {
	N       int
	IR      int
	CT      int
	R       int
	Address map[string]HostPort
}

func ReadConfig(file string) (Config, error) {
	cfgObj := Config{}
	cfgObj.Address = make(map[string]HostPort)
	readFile, err := os.Open(file)
	if err != nil {
		return cfgObj, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	i := 1
	for fileScanner.Scan() {
		if !strings.HasPrefix(fileScanner.Text(), "#") && len(strings.Split(strings.ReplaceAll(fileScanner.Text(), " ", ""), "\n")[0]) > 0 {
			s := strings.Split(fileScanner.Text(), " ")
			if i == 1 {
				cfgObj.N, err = strconv.Atoi(s[0])
				if err != nil {
					return cfgObj, err
				}
				cfgObj.IR, err = strconv.Atoi(s[1])
				if err != nil {
					return cfgObj, err
				}
				cfgObj.CT, err = strconv.Atoi(s[2])
				if err != nil {
					return cfgObj, err
				}
				cfgObj.R, err = strconv.Atoi(s[3])
				if err != nil {
					return cfgObj, err
				}
				i = 2
			} else {
				hp := HostPort{
					Host: s[1],
					Port: s[2],
				}
				cfgObj.Address[s[0]] = hp
			}
		}
	}

	return cfgObj, nil
}
