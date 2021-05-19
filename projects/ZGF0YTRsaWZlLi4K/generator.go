package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
)

func token_generator(token_len int) string {
	// should I seed the rand?
	var token []rune = make([]rune, token_len)
	for i := 0; i < token_len; i++ {
		token[i] = 'a' + rune(rand.Intn(26))
	}
	return string(token)
}

func write_to_file(fh *os.File, lines int, token_len int, wg *sync.WaitGroup) error {
	defer fh.Sync()
	defer wg.Done()

	for i := 0; i < lines; i++ {
		_, err := fh.WriteString(token_generator(token_len) + "\n")
		if err != nil {
			fmt.Printf("[ERROR] Failed to write token: %v\n", err)
			return err
		}
	}
	return nil
}

func generate_file(lines int, token_len int) error {
	var wg sync.WaitGroup
	var err error
	total_work := lines
	batch := 0

	workers := runtime.NumCPU()
	if workers < 1 {
		fmt.Printf("[WARN] Failed to get CPU count, defaulting to 1")
		workers = 1
	}
	fhs := make([]*os.File, workers)
	for i := range fhs {
		fhs[i], err = os.OpenFile("data_tokens", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer fhs[i].Close()

		fhs[i].Seek(int64(i*(token_len+1)*batch), 0)

		if i < workers-1 {
			batch = int(lines / workers)
			total_work -= batch
		} else {
			batch = total_work
		}
		wg.Add(1)
		go write_to_file(fhs[i], batch, token_len, &wg)
	}

	// I will not treat individual routine errors as I believe it is beyond the scope of the exercise
	// to handle the errors from multiple routines I would use buffered channels
	wg.Wait()

	return nil
}
