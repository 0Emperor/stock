package parser

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseFile(filename string) (*Stock_exchange, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := &Stock_exchange{
		Stock:     make(map[string]int),
		Tasks: []Task{},
		To_Optimize:  []string{},
	}

	scanner := bufio.NewScanner(file)
	stockRegex := regexp.MustCompile(`^([a-zA-Z0-9_]+):(\d+)$`)
	processRegex := regexp.MustCompile(`^([a-zA-Z0-9_]+):\((.*?)\):\((.*?)\):(\d+)$`)
	optimizeRegex := regexp.MustCompile(`^optimize:\((.*?)\)$`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if match := stockRegex.FindStringSubmatch(line); match != nil {
			qty, _ := strconv.Atoi(match[2])
			data.Stock[match[1]] = qty
			continue
		}

		if match := processRegex.FindStringSubmatch(line); match != nil {
			reqs := parseKeyValuePairs(match[2])
			prods := parseKeyValuePairs(match[3])
			cycles, _ := strconv.Atoi(match[4])
			data.Tasks = append(data.Tasks, Task{
				Name:         match[1],
				Requirements: reqs,
				Products:     prods,
				NbCycles:     cycles,
			})
			continue
		}

		if match := optimizeRegex.FindStringSubmatch(line); match != nil {
			data.To_Optimize = strings.Split(match[1], ";")
			continue
		}
		return nil,errors.New("error parsing: " + line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func parseKeyValuePairs(input string) map[string]int {
	result := make(map[string]int)
	if input == "" {
		return result
	}
	pairs := strings.Split(input, ";")
	for _, pair := range pairs {
		parts := strings.Split(pair, ":")
		if len(parts) == 2 {
			qty, _ := strconv.Atoi(parts[1])
			result[parts[0]] = qty
		}
	}
	return result
}
