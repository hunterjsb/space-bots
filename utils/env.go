package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ReadEnv reads a .env file and returns a map of the key-value pairs.
func ReadEnv(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	envMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue // skip empty lines and comments
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // skip lines that don't have key and value
		}

		envMap[parts[0]] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envMap, nil
}

func initEnv(filePath string) map[string]string {
	env, err := ReadEnv(filePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return env
}

var Env map[string]string = initEnv(".env")
