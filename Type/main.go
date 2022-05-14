package main

import "fmt"

type HTTPState int

const (
	StatusOK              HTTPState = 200
	StatusUnauthorized    HTTPState = 401
	StatusPaymentRequired HTTPState = 402
	StatusForbidden       HTTPState = 403
)

// HTTPStateに対してString関数を組み込む
func (s HTTPState) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	default:
		return fmt.Sprintf("HTTPStatus(%d)", s)
	}
}