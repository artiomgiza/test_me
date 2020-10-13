package gotesttableme

import (
	beeffarm "github.com/artiomgiza/test_me/pkgs/beef-farm"
	chickenfarm "github.com/artiomgiza/test_me/pkgs/chicken-farm"
	coalfarm "github.com/artiomgiza/test_me/pkgs/coal-farm"
	mangalstore "github.com/artiomgiza/test_me/pkgs/mangal-store"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Provider interface {
	CalculatePrice(peopleCounter int) (int, error)
}

var Instance Provider = coolPriceCalculator{
	// ... inject real fields ...
}

//////////////////////////////////////////////////////////
// implementation ////////////////////////////////////////

type coolPriceCalculator struct {
	beefFarm    beeffarm.Provider
	chickenFarm chickenfarm.Provider
	mangalStore mangalstore.Provider
	coalFarm    coalfarm.Provider
}

const (
	maxPeopleCounter = 10
)

func (b coolPriceCalculator) CalculatePrice(peopleCounter int) (int, error) {
	if peopleCounter > maxPeopleCounter {
		return 0, errors.New("too much people")
	}

	meetPrice, err := b.beefFarm.GetEntrecote(peopleCounter)
	if err != nil {
		logrus.WithError(err).Error("could not get entrecote, fallback to chicken")

		meetPrice, err = b.chickenFarm.GetPullet(peopleCounter)
		if err != nil {
			return 0, errors.Wrap(err, "no meet available")
		}
	}

	mangalPrice, err := b.mangalStore.GetMangal()
	if err != nil {
		return 0, errors.Wrap(err, "no mangal available")
	}

	coalPrice, err := b.coalFarm.GetCoal(1)
	if err != nil {
		return 0, errors.Wrap(err, "no coal available")
	}

	totalPrice := meetPrice + mangalPrice + coalPrice
	return totalPrice, nil
}
