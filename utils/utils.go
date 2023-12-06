package utils

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ReadLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Couldn't open file '%s' due to error: %s\n", filePath, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("Couldn't parse contents of file '%s' due to error: %s\n", filePath, scanner.Err())
	}
	return lines
}

func FirstOrDefault[T any](slice []T, defaultValue T) T {
	if len(slice) == 0 {
		return defaultValue
	}

	return slice[0]
}

func LinesFromShell(shellCommand string) []string {
	var stdout bytes.Buffer
	cmd := exec.Command("bash", "-c", shellCommand)

	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Couldn't run command '%s' due to error: %s\n", shellCommand, err)
	}

	split := strings.Split(stdout.String(), "\n")
	return split[:len(split)-1]
}

func ToInt(str string) int {
	converted, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Couldn't convert '%s' to int due to error: %s\n", str, err)
	}
	return converted
}
