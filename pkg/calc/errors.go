package calc

import "errors"

var (
	ErrDivByZero       = errors.New("На нуль делить нельзя!")
	ErrInvalidBracket  = errors.New("Ошибка в скобках!")
	ErrInvalidOperands = errors.New("Проверьте количество операндов(+,-,/,*), их порядок и проверьте что нет букв!")
)
