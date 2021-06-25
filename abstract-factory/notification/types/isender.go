package types

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}
