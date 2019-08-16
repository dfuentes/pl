package main

import (
	"bufio"
	"os"
)

type Config struct {
	Subscriptions []string
}

func Load(path string) (*Config, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	subs := []string{}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		subs = append(subs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Config{
		Subscriptions: subs,
	}, nil
}
