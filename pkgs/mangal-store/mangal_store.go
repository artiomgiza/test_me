package mangalstore

// mockgen -source=pkgs/mangal-store/mangal_store.go -destination=pkgs/mangal-store/mock/mangal_store_mock.go
type Provider interface {
	GetMangal() (price int, err error)
}
