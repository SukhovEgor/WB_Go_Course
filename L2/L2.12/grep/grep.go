package grep

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Grep(data []string, pattern string, config GrepConfig) ([]string, error) {
	var (
		result []string
		re     *regexp.Regexp
		err    error
	)

	if !config.FixedString {
		if config.IgnoreCase {
			re, err = regexp.Compile("(?i)" + pattern)
		} else {
			re, err = regexp.Compile(pattern)
		}

		if err != nil {
			return nil, err
		}
	}

	found := make([]bool, len(data))

	for i, line := range data {
		if match(line, pattern, re, config) {
			found[i] = true
		}
	}

	if config.Count {
		return []string{
			strconv.Itoa(countMatches(found)),
		}, nil
	}

	output := make([]bool, len(data))

	for i := range found {
		if !found[i] {
			continue
		}

		start := max(0, i-config.BeforeContext)
		end := min(len(data)-1, i+config.AfterContext)

		for j := start; j <= end; j++ {
			output[j] = true
		}
	}

	for i, line := range data {
		if !output[i] {
			continue
		}

		if config.LineNumber {
			result = append(result,
				fmt.Sprintf("%d:%s", i+1, line))
		} else {
			result = append(result, line)
		}
	}

	return result, nil
}

func match(
	line string,
	pattern string,
	re *regexp.Regexp,
	config GrepConfig,
) bool {
	var matched bool

	if config.FixedString {
		if config.IgnoreCase {
			line = strings.ToLower(line)
			pattern = strings.ToLower(pattern)
		}

		matched = strings.Contains(line, pattern)
	} else {
		matched = re.MatchString(line)
	}

	if config.InvertMatch {
		matched = !matched
	}

	return matched
}

func countMatches(matches []bool) int {
	count := 0

	for _, match := range matches {
		if match {
			count++
		}
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}