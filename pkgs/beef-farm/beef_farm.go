package beeffarm

// mockgen -source=pkgs/beef-farm/beef_farm.go -destination=pkgs/beef-farm/mock/beef_farm_mock.go
type Provider interface {
	GetEntrecote(weight int) (price int, err error)
	GetTBone(weight int) (price int, err error)
}
