package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
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

	contents, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	subs := []string{}

	scanner := bufio.NewScanner(bytes.NewReader(contents))
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
