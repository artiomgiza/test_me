package bbq

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewCoolPriceCalculator(bf BeefFarm, chf ChickenFarm, ms MangalStore, cf CoalFarm) *coolPriceCalculator {
	return &coolPriceCalculator{
		beefFarm:    bf,
		chickenFarm: chf,
		mangalStore: ms,
		coalFarm:    cf,
	}
}

// A package should have interfaces and mocks for external dependencies within itself.
type (
	BeefFarm interface {
		GetEntrecotePrice(portionsCounter int) (price int, err error)
	}

	ChickenFarm interface {
		GetPulletPrice(portionsCounter int) (price int, err error)
	}

	MangalStore interface {
		GetMangalPrice() (price int, err error)
	}

	CoalFarm interface {
		GetCoalPrice(mangalsCounter int) (price int, err error)
	}
)

//go:generate mockgen -destination=./mocks/beef_farm_mock.go -package=mocks github.com/artiomgiza/test_me/go_native_with_logical_flows BeefFarm
//go:generate mockgen -destination=./mocks/chicken_farm_mock.go -package=mocks github.com/artiomgiza/test_me/go_native_with_logical_flows ChickenFarm
//go:generate mockgen -destination=./mocks/mangal_store_mock.go -package=mocks github.com/artiomgiza/test_me/go_native_with_logical_flows MangalStore
//go:generate mockgen -destination=./mocks/coal_farm_mock.go -package=mocks github.com/artiomgiza/test_me/go_native_with_logical_flows CoalFarm

type coolPriceCalculator struct {
	beefFarm    BeefFarm
	chickenFarm ChickenFarm
	mangalStore MangalStore
	coalFarm    CoalFarm
}

const (
	maxPeopleCounter = 10
)

func (b coolPriceCalculator) CalculatePrice(peopleCounter int) (int, error) {
	//time.Sleep(time.Second) // try uncommenting this to see the advantage of parallel tests

	if peopleCounter > maxPeopleCounter {
		return 0, errors.New("too much people")
	}

	meetPrice, err := b.beefFarm.GetEntrecotePrice(peopleCounter)
	if err != nil {
		logrus.WithError(err).Error("could not get entrecote, fallback to chicken")

		meetPrice, err = b.chickenFarm.GetPulletPrice(peopleCounter)
		if err != nil {
			return 0, errors.Wrap(err, "no meet available")
		}
	}

	mangalPrice, err := b.mangalStore.GetMangalPrice()
	if err != nil {
		return 0, errors.Wrap(err, "no mangal available")
	}

	coalPrice, err := b.coalFarm.GetCoalPrice(1)
	if err != nil {
		return 0, errors.Wrap(err, "no coal available")
	}

	totalPrice := meetPrice + mangalPrice + coalPrice
	return totalPrice, nil
}
