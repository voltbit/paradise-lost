package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func generate_token(tokenLen int) string {
	var token []rune = make([]rune, tokenLen)
	for i := 0; i < tokenLen; i++ {
		token[i] = 'a' + rune(rand.Intn(26))
	}
	return string(token)
}

func generateTokenBlockWithEndlines(tokenLen int, tokenCount int) string {
	lineSize := tokenLen + 1
	token := make([]rune, lineSize*tokenCount)
	for i := 0; i < len(token); i += lineSize {
		for j := 0; j < tokenLen; j++ {
			token[i+j] = 'a' + rune(rand.Intn(26))
		}
		token[i+tokenLen] = '\n'
	}
	return string(token[:len(token)-1])
}

func generateTokenBlock(tokenLen int, tokenCount int) string {
	token := make([]rune, tokenLen*tokenCount)
	for i := 0; i < len(token); i++ {
		token[i] = 'a' + rune(rand.Intn(26))
	}
	return string(token)
}

func write_to_file(fh *os.File, lines int, tokenLen int, wg *sync.WaitGroup) error {
	defer wg.Done()

	_, err := fh.WriteString(generateTokenBlock(tokenLen, lines))
	if err != nil {
		fmt.Printf("[ERROR] Failed to write token: %v\n", err)
		return err
	}
	return nil
}

func generateFile(lines int, tokenLen int, endlines bool) error {
	var wg sync.WaitGroup
	var err error
	total_work := lines
	batch := 0

	// workers := runtime.NumCPU()
	workers := 1
	if workers < 1 {
		fmt.Printf("[WARN] Failed to get CPU count, defaulting to 1")
		workers = 1
	}

	fmt.Printf("[G] Generating %d tokens of size %d\n", lines, tokenLen)
	fmt.Printf("[G] Using %d workers\n", workers)

	fhs := make([]*os.File, workers)
	for i := range fhs {
		fhs[i], err = os.OpenFile("data_file", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer fhs[i].Sync()
		defer fhs[i].Close()

		if endlines {
			fhs[i].Seek(int64(i*(tokenLen+1)*batch), 0)
		} else {
			fhs[i].Seek(int64(i*(tokenLen)*batch), 0)
		}

		if i < workers-1 {
			batch = int(lines / workers)
			total_work -= batch
		} else {
			batch = total_work
		}
		wg.Add(1)
		go write_to_file(fhs[i], batch, tokenLen, &wg)
	}

	// I will not treat individual routine errors as I believe it is beyond the scope of the exercise
	// to handle the errors from multiple routines I would use buffered channels
	wg.Wait()

	return nil
}
