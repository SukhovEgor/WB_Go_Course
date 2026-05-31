//Необходимо реализовать собственный простейший Unix shell.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Mini Shell")

	for {
		fmt.Print("$ ")

		if !scanner.Scan() {
			fmt.Println()
			return
		}

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if line == "exit" {
			return
		}

		if strings.Contains(line, "|") {
			if err := runPipeline(line); err != nil {
				fmt.Println("error:", err)
			}
			continue
		}

		if err := runCommand(strings.Fields(line)); err != nil {
			fmt.Println("error:", err)
		}
	}
}

func runCommand(args []string) error {
	if len(args) == 0 {
		return nil
	}

	switch args[0] {

	case "cd":
		return builtinCD(args)

	case "pwd":
		return builtinPWD()

	case "echo":
		return builtinEcho(args)

	case "kill":
		return builtinKill(args)

	case "ps":
		return builtinPS()
	}

	return runExternal(args)
}

func builtinCD(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("cd: path required")
	}

	return os.Chdir(args[1])
}

func builtinPWD() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println(dir)
	return nil
}

func builtinEcho(args []string) error {
	fmt.Println(strings.Join(args[1:], " "))
	return nil
}

func builtinKill(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("kill: pid required")
	}

	pid, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid pid")
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return proc.Signal(syscall.SIGTERM)
}

func builtinPS() error {
	cmd := exec.Command("tasklist")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func runExternal(args []string) error {
	cmd := exec.Command(
		"cmd",
		"/C",
		strings.Join(args, " "),
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func runPipeline(line string) error {
	stages := strings.Split(line, "|")

	var cmds []*exec.Cmd

	for _, stage := range stages {
		stage = strings.TrimSpace(stage)

		if stage == "" {
			continue
		}

		cmds = append(cmds,
			exec.Command(
				"cmd",
				"/C",
				stage,
			))
	}

	if len(cmds) == 0 {
		return nil
	}

	for i := 0; i < len(cmds)-1; i++ {
		reader, writer := io.Pipe()

		cmds[i].Stdout = writer
		cmds[i+1].Stdin = reader
	}

	last := len(cmds) - 1

	cmds[last].Stdout = os.Stdout
	cmds[last].Stderr = os.Stderr

	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return err
		}
	}

	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			return err
		}
	}

	return nil
}