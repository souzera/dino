package handler

import "fmt"

func errorParamRequired(param string) error {
	return fmt.Errorf("parametro %s é obrigatório", param)
}

func errorUniqueViolation(param string) error {
	return fmt.Errorf("parametro %s já existe", param)
}