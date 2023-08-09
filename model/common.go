package model

type BasicResponse interface {
	GetCode() int
	GetMsg() string
}
