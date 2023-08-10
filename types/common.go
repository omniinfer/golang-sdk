package types

type BasicResponse interface {
	GetCode() int
	GetMsg() string
}
