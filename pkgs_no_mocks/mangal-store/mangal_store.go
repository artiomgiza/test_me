package mangalstore

var Instance Provider

type Provider interface {
	GetMangal() (price int, err error)
}
