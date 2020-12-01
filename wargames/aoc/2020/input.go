package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func GetIntList(name string) ([]int, error) {
	var result []int
	fh, err := os.Open(name)
	defer fh.Close()
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReader(fh)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				return result, nil
			}
			return nil, err
		}
		num, err := strconv.Atoi(string(line))
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
}
