package main

import (
	"fmt"
	"io/ioutil"
	"math"

	pg "github.com/go-pg/pg/v10"
)

type Token struct {
	Token     string
	Frequency int
}

type TokenProcessor struct {
	DBClient *pg.DB
	Data     map[string]int
}

func NewTokenProcessor(dbUser string, dbPass string, dbName string) (*TokenProcessor, error) {
	client := pg.Connect(&pg.Options{
		Addr:     "localhost:9000",
		User:     dbUser,
		Password: dbPass,
		Database: dbName,
	})
	tp := &TokenProcessor{DBClient: client}
	tp.Data = make(map[string]int)
	// runtime.SetFinalizer(tp, )
	return tp, nil
}

func (t *TokenProcessor) sequentialWrite(token string, frequency int) error {
	res, err := t.DBClient.Model(&Token{
		Token:     token,
		Frequency: frequency,
	}).Insert()
	fmt.Printf("Write result: %v\n", res)

	return err
}

func dbWriter(queries <-chan func(values ...interface{}) (pg.Result, error), results chan<- func() (pg.Result, error)) {
	for q := range queries {
		res, err := q()
		results <- (func() (pg.Result, error) { return res, err })
	}
}

func (t *TokenProcessor) batchedWrite(batchSize int, workers int) error {
	queryCount := int(math.Ceil(float64(len(t.Data)) / float64(batchSize)))
	// queries := make(chan func() (values interface{} ) (error), queryCount)
	queries := make(chan func(values ...interface{}) (pg.Result, error), queryCount)
	results := make(chan func() (pg.Result, error), queryCount)
	batchCount := 0
	currentSize := 0
	tokenRows := make([]interface{}, 0)

	for i := 0; i < workers; i++ {
		go dbWriter(queries, results)
	}

	for tk, fr := range t.Data {
		tokenRows = append(tokenRows, Token{Token: tk, Frequency: fr})
		currentSize++
		if currentSize >= batchSize {
			queries <- t.DBClient.Model(tokenRows...).OnConflict("(token) DO UPDATE").Set("frequency = EXCLUDED.frequency + token.frequency").Insert
			batchCount++
			currentSize = 0
			tokenRows = make([]interface{}, 0)
		}
	}
	if currentSize > 0 {
		queries <- t.DBClient.Model(tokenRows...).OnConflict("(token) DO UPDATE").Set("frequency = EXCLUDED.frequency + token.frequency").Insert
	}
	for i := 0; i < queryCount; i++ {
		res, err := (<-results)()
		if err != nil {
			fmt.Printf("Write job failed for batch %d: %v\n", i, err)
			return err
		}
		fmt.Printf("Write of batch [%d]: %v\n", i, res)
	}
	return nil
}

func (t *TokenProcessor) bck_batchedWrite(batchSize int, workers int) error {
	batchCount := 0
	currentSize := 0
	tokenRows := make([]interface{}, 0)
	for tk, fr := range t.Data {
		tokenRows = append(tokenRows, Token{Token: tk, Frequency: fr})
		currentSize++
		if currentSize >= batchSize {
			res, err := t.DBClient.Model(tokenRows...).OnConflict("(token) DO UPDATE").Set("frequency = EXCLUDED.frequency + token.frequency").Insert()
			if err != nil {
				return err
			}
			fmt.Printf("Write of batch [%d]: %v\n", batchCount, res)
			batchCount++
			currentSize = 0
			tokenRows = make([]interface{}, 0)
		}
	}
	if currentSize > 0 {
		res, err := t.DBClient.Model(tokenRows...).OnConflict("(token) DO UPDATE").Set("frequency = EXCLUDED.frequency + token.frequency").Insert()
		if err != nil {
			return err
		}
		fmt.Printf("Write of batch [%d]: %v\n", batchCount, res)
	}
	return nil
}

func (t *TokenProcessor) mapDataWithEndlines(tokenLength int) error {
	lineSize := tokenLength + 1
	rawData, err := ioutil.ReadFile("data_file")
	for i := 0; i < len(rawData)-tokenLength; i += lineSize {
		t.Data[string(rawData[i:(i+tokenLength)])]++
	}
	if err != nil {
		fmt.Printf("Could not read the file: %v\n", err)
		return err
	}
	fmt.Printf("[P] Total tokens: %d\n", len(rawData)/lineSize)
	fmt.Printf("[P] Unique tokens: %d\n", len(t.Data))
	return nil
}

func (t *TokenProcessor) mapData(tokenLength int) error {
	rawData, err := ioutil.ReadFile("data_file")
	for i := 0; i <= len(rawData)-tokenLength; i += tokenLength {
		t.Data[string(rawData[i:(i+tokenLength)])]++
	}
	if err != nil {
		fmt.Printf("Could not read the file: %v\n", err)
		return err
	}
	fmt.Printf("[P] Total tokens: %d\n", len(rawData)/tokenLength)
	fmt.Printf("[P] Unique tokens: %d\n", len(t.Data))
	return nil
}

func (t *TokenProcessor) cleanTable() error {
	_, err := t.DBClient.Exec("truncate table tokens")
	return err
}

func (t *TokenProcessor) start(tokenSize int) error {
	if err := t.cleanTable(); err != nil {
		return err
	}
	t.mapData(tokenSize)
	if err := t.batchedWrite(200000, 6); err != nil {
		return err
	}
	// time.Sleep(time.Second * 20)
	return nil
}
