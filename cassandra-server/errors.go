package main

import "errors"

var (
	ErrTableExists = errors.New("table already exists in KeySpace")
)
