package chickenfarm

var Instance Provider

type Provider interface {
	GetPullet(weight int) (price int, err error)
}
