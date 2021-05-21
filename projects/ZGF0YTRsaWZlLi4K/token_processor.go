package main

import (
	pg "github.com/go-pg/pg/v10"
)

type Token struct {
	Token string
	Freq  int
}

type TokenProcessor struct {
	DBClient *pg.DB
}

func NewTokenProcessor(dbUser string, dbPass string, dbName string) (*TokenProcessor, error) {
	client := pg.Connect(&pg.Options{
		Addr:     "localhost:9000",
		User:     dbUser,
		Password: dbPass,
		Database: dbName,
	})
	tp := &TokenProcessor{DBClient: client}
	// runtime.SetFinalizer(tp, )
	return tp, nil
}

func (t *TokenProcessor) writeData() error {

	return nil
}

func (t *TokenProcessor) readFile() error {
	return nil
}

func (t *TokenProcessor) start() error {
	if err := t.readFile(); err != nil {
		return err
	}
	if err := t.writeData(); err != nil {
		return err
	}
	return nil
}
