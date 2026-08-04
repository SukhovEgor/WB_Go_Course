/* Реализовать утилиту, которая считывает входные данные (STDIN) и разбивает каждую строку
 по заданному разделителю,
после чего выводит определённые поля (колонки).
Аналог команды cut с поддержкой флагов:
-f "fields" — указание номеров полей (колонок), которые нужно вывести. Номера через запятую, можно диапазоны.
Например: «-f 1,3-5» — вывести 1-й и с 3-го по 5-й столбцы.
-d "delimiter" — использовать другой разделитель (символ). По умолчанию разделитель — табуляция ('\t').
-s – (separated) только строки, содержащие разделитель. Если флаг указан, то строки без разделителя игнорируются (не выводятся).
Программа должна корректно парсить аргументы, поддерживать различные комбинации 
(например, несколько отдельных полей и диапазонов), учитывать, что номера полей
могут выходить за границы (в таком случае эти поля просто игнорируются).
Стоит обратить внимание на эффективность при обработке больших файлов.
Все стандартные требования по качеству кода и тестам также применимы. */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	fields    map[int]bool
	delimiter string
	separated bool
}

func main() {
	cfg, err := parseConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		processLine(scanner.Text(), cfg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parseConfig() (*Config, error) {
	fieldsSpec := flag.String("f", "", "fields to print")
	delimiter := flag.String("d", "\t", "field delimiter")
	separated := flag.Bool("s", false, "only lines containing delimiter")

	flag.Parse()

	if *fieldsSpec == "" {
		return nil, fmt.Errorf("-f is required")
	}

	if len(*delimiter) != 1 {
		return nil, fmt.Errorf("delimiter must be exactly one character")
	}

	fields, err := parseFields(*fieldsSpec)
	if err != nil {
		return nil, err
	}

	return &Config{
		fields:    fields,
		delimiter: *delimiter,
		separated: *separated,
	}, nil
}

func parseFields(spec string) (map[int]bool, error) {
	result := make(map[int]bool)

	parts := strings.Split(spec, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)

		if strings.Contains(part, "-") {
			bounds := strings.Split(part, "-")

			if len(bounds) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			start, err := strconv.Atoi(bounds[0])
			if err != nil || start < 1 {
				return nil, fmt.Errorf("invalid field: %s", bounds[0])
			}

			end, err := strconv.Atoi(bounds[1])
			if err != nil || end < start {
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			for i := start; i <= end; i++ {
				result[i-1] = true
			}

			continue
		}

		field, err := strconv.Atoi(part)
		if err != nil || field < 1 {
			return nil, fmt.Errorf("invalid field: %s", part)
		}

		result[field-1] = true
	}

	return result, nil
}

func processLine(line string, cfg *Config) {
	if cfg.separated && !strings.Contains(line, cfg.delimiter) {
		return
	}

	parts := strings.Split(line, cfg.delimiter)

	var output []string

	for idx, value := range parts {
		if cfg.fields[idx] {
			output = append(output, value)
		}
	}

	if len(output) > 0 {
		fmt.Println(strings.Join(output, cfg.delimiter))
	}
}