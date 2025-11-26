package parser

import "io"


type Service interface {
	ParseCSV(file io.Reader) ([]Transaction, error)
}

type service struct {}

func NewService() Service {
	return &service{}
}