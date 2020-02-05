package coalfarm

var Instance Provider

// mockgen -source=pkgs/coal-farm/coal_farm.go -destination=pkgs/coal-farm/mock/coal_farm_mock.go
type Provider interface {
	GetCoal(mangalsCounter int) (price int, err error)
}
