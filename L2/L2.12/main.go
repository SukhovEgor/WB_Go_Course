/* Утилита grep
Реализовать утилиту фильтрации текстового потока (аналог команды grep).
Программа должна читать входной поток (STDIN или файл) и выводить строки, соответствующие заданному шаблону (подстроке или регулярному выражению).
Необходимо поддерживать следующие флаги:
-A N — после каждой найденной строки дополнительно вывести N строк после неё (контекст).
-B N — вывести N строк до каждой найденной строки.
-C N — вывести N строк контекста вокруг найденной строки (включает и до, и после; эквивалентно -A N -B N).
-c — выводить только то количество строк, что совпадающих с шаблоном (т.е. вместо самих строк — число).
-i — игнорировать регистр.
-v — инвертировать фильтр: выводить строки, не содержащие шаблон.
-F — воспринимать шаблон как фиксированную строку, а не регулярное выражение (т.е. выполнять точное совпадение подстроки).
-n — выводить номер строки перед каждой найденной строкой.

Программа должна поддерживать сочетания флагов (например, -C 2 -n -i – 2 строки контекста, вывод номеров, без учета регистра и т.д.).
Результат работы должен максимально соответствовать поведению команды UNIX grep.
Обязательно учесть пограничные случаи (начало/конец файла для контекста, повторяющиеся совпадения и пр.).
Код должен быть чистым, отформатированным (gofmt), работать без ситуаций гонки и успешно проходить golint. */

package main

import (
	"flag"
	"fmt"
	"os"

	"grep/grep"
)

func main() {
	config := parseConfig()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: grep [options] PATTERN [FILE]")
		os.Exit(1)
	}

	pattern := flag.Arg(0)

	var (
		data []string
		err  error
	)

	if flag.NArg() >= 2 {
		data, err = grep.ReadFromFile(flag.Arg(1))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
			os.Exit(1)
		}
	} else {
		data = grep.ReadFromStdin()
	}

	result, err := grep.Grep(data, pattern, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grep error: %v\n", err)
		os.Exit(1)
	}

	for _, line := range result {
		fmt.Println(line)
	}
}

func parseConfig() grep.GrepConfig {
	var config grep.GrepConfig

	flag.IntVar(&config.AfterContext, "A", 0, "print NUM lines of trailing context")
	flag.IntVar(&config.BeforeContext, "B", 0, "print NUM lines of leading context")
	flag.IntVar(&config.Context, "C", 0, "print NUM lines of output context")

	flag.BoolVar(&config.Count, "c", false, "print only count of matching lines")
	flag.BoolVar(&config.IgnoreCase, "i", false, "ignore case")
	flag.BoolVar(&config.InvertMatch, "v", false, "select non-matching lines")
	flag.BoolVar(&config.FixedString, "F", false, "treat pattern as fixed string")
	flag.BoolVar(&config.LineNumber, "n", false, "show line numbers")

	flag.Parse()

	config.Contextualize()

	return config
}