/* Утилита sort
Реализовать упрощённый аналог UNIX-утилиты sort (сортировка строк).
Программа должна читать строки (из файла или STDIN) и выводить их отсортированными.

Обязательные флаги (как в GNU sort):
-k N — сортировать по столбцу (колонке) №N (разделитель — табуляция по умолчанию).
Например, «sort -k 2» отсортирует строки по второму столбцу каждой строки.
-n — сортировать по числовому значению (строки интерпретируются как числа).
-r — сортировать в обратном порядке (reverse).
-u — не выводить повторяющиеся строки (только уникальные). */



package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Config struct {
	Column      int
	Numeric     bool
	Reverse     bool
	Unique      bool
	Month       bool
	IgnoreBlanks bool
	Check       bool
	Human       bool
}

var months = map[string]int{
	"jan": 1,
	"feb": 2,
	"mar": 3,
	"apr": 4,
	"may": 5,
	"jun": 6,
	"jul": 7,
	"aug": 8,
	"sep": 9,
	"oct": 10,
	"nov": 11,
	"dec": 12,
}

func main() {
	cfg := parseFlags()

	lines, err := readInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if cfg.Unique {
		lines = unique(lines)
	}

	if cfg.Check {
		if isSorted(lines, cfg) {
			fmt.Println("data is sorted")
		} else {
			fmt.Println("data is not sorted")
		}
		return
	}

	sort.Slice(lines, func(i, j int) bool {
		return less(lines[i], lines[j], cfg)
	})

	if cfg.Reverse {
		reverse(lines)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func parseFlags() Config {
	var cfg Config

	flag.IntVar(&cfg.Column, "k", 0, "column number")
	flag.BoolVar(&cfg.Numeric, "n", false, "numeric sort")
	flag.BoolVar(&cfg.Reverse, "r", false, "reverse sort")
	flag.BoolVar(&cfg.Unique, "u", false, "unique lines")
	flag.BoolVar(&cfg.Month, "M", false, "month sort")
	flag.BoolVar(&cfg.IgnoreBlanks, "b", false, "ignore trailing blanks")
	flag.BoolVar(&cfg.Check, "c", false, "check if sorted")
	flag.BoolVar(&cfg.Human, "h", false, "human readable numbers")

	flag.Parse()

	return cfg
}

func readInput() ([]string, error) {
	var scanner *bufio.Scanner

	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func less(a, b string, cfg Config) bool {
	keyA := extractKey(a, cfg)
	keyB := extractKey(b, cfg)

	if cfg.Month {
		return monthValue(keyA) < monthValue(keyB)
	}

	if cfg.Human {
		return humanValue(keyA) < humanValue(keyB)
	}

	if cfg.Numeric {
		numA, _ := strconv.ParseFloat(keyA, 64)
		numB, _ := strconv.ParseFloat(keyB, 64)
		return numA < numB
	}

	return keyA < keyB
}

func extractKey(line string, cfg Config) string {
	if cfg.IgnoreBlanks {
		line = strings.TrimRight(line, " ")
	}

	if cfg.Column <= 0 {
		return line
	}

	parts := strings.Split(line, "\t")

	if cfg.Column > len(parts) {
		return ""
	}

	return parts[cfg.Column-1]
}

func monthValue(s string) int {
	return months[strings.ToLower(strings.TrimSpace(s))]
}

func humanValue(s string) float64 {
	s = strings.TrimSpace(strings.ToUpper(s))

	multiplier := 1.0

	switch {
	case strings.HasSuffix(s, "K"):
		multiplier = 1024
		s = strings.TrimSuffix(s, "K")

	case strings.HasSuffix(s, "M"):
		multiplier = 1024 * 1024
		s = strings.TrimSuffix(s, "M")

	case strings.HasSuffix(s, "G"):
		multiplier = 1024 * 1024 * 1024
		s = strings.TrimSuffix(s, "G")
	}

	v, _ := strconv.ParseFloat(s, 64)

	return v * multiplier
}

func unique(lines []string) []string {
	seen := make(map[string]struct{})

	var result []string

	for _, line := range lines {
		if _, ok := seen[line]; ok {
			continue
		}

		seen[line] = struct{}{}
		result = append(result, line)
	}

	return result
}

func reverse(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func isSorted(lines []string, cfg Config) bool {
	for i := 1; i < len(lines); i++ {
		if less(lines[i], lines[i-1], cfg) {
			return false
		}
	}
	return true
}