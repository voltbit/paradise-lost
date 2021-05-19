package main

import (
	pg "github.com/go-pg/pg/v10"
)

type TokenProcessor struct {
	DBClient *pg.DB
}

func NewTokenProcessor(dbUser string, dbPass string, dbName string) (*TokenProcessor, error) {
	client := pg.Connect(&pg.Options{
		Addr:     "postgres:5432",
		User:     dbUser,
		Password: dbPass,
		Database: dbName,
	})
	tp := &TokenProcessor{DBClient: client}
	// runtime.SetFinalizer(tp, )
	return tp, nil
}

func (t *TokenProcessor) generateSchema() {

}

func (t *TokenProcessor) writeData() {

}

func (t *TokenProcessor) readFile() {

}

func (t *TokenProcessor) process() {
	t.generateSchema()
	t.readFile()
	t.writeData()
}
