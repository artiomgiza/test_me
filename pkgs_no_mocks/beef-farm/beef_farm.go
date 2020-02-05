package beeffarm

var Instance Provider

type Provider interface {
	GetEntrecote(weight int) (price int, err error)
	GetTBone(weight int) (price int, err error)
}
