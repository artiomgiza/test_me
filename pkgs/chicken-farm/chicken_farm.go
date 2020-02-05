package chickenfarm

var Instance Provider

// mockgen -source=pkgs/chicken-farm/chicken_farm.go -destination=pkgs/chicken-farm/mock/chicken_farm_mock.go
type Provider interface {
	GetPullet(weight int) (price int, err error)
}
