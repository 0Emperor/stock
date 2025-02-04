package parser

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func ParseFile(FileName string) error {
	file, err := os.Open(FileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		line := scaner.Text()
		if line[0] == '$' {
			continue
		}
		Spline, err := Split(line)
		if err != nil {
			return err
		}
		if len(Spline) == 2 {
			if strings.HasPrefix(line, "optimize") {
				
			}
		} else if len(Spline) == 4 {
		} else {
			return errors.New("error parsing " + line)
		}
	}
	return nil
}

func Split(line string) ([]string, error) {
	slice := []string{}
	holder := ""
	for i := 0; i < len(line); i++ {
		v := line[i]
		if v == ':' {
			if holder == "" {
				return nil, errors.New("error parsing " + line)
			}
			slice = append(slice, holder)
		} else if v == '(' {
			if holder != "" {
				return nil, errors.New("error parsing " + line)
			}
			g := SearchPair(line[i:])
			if g == -1 {
				return nil, errors.New("error parsing " + line)
			} else {
				slice = append(slice, line[i+1:i+g])
				i = g
			}
		} else {
			holder += string(v)
		}
	}
	return slice, nil
}

func SearchPair(s string) int {
	for i, v := range s {
		if v == ')' {
			return i
		}
	}
	return -1
}
