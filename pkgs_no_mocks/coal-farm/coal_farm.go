package coalfarm

var Instance Provider

type Provider interface {
	GetCoal(mangalsCounter int) (price int, err error)
}
