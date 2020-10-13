package ginkgo_me

import (
	sunnyfarm "farm"
	somepkg "google"

	"github.com/Sirupsen/logrus"

	"github.com/pkg/errors"
)

type CoolPriceCalculator struct{}

const (
	maxPeopleCounter = 10
)

func (b CoolPriceCalculator) CalculatePrice(peopleCounter int) (int, error) {
	if peopleCounter > maxPeopleCounter {
		return 0, errors.New("too much people")
	}

	///////////////////////////
	// Get meet ///////////////
	///////////////////////////

	farms, err := somepkg.GetNearbyFarms()
	if err != nil {
		return 0, errors.Wrap(err, "could not get farms")
	}

	farm, err := chooseFarm(farms)
	if err != nil {
		return 0, errors.Wrap(err, "could not choose farms")
	}

	farmPhone, err := somepkg.GetFarmPhone(farm)
	if err != nil {
		return 0, errors.Wrap(err, "could not get farm phone farms")
	}

	ok, price, err := somepkg.ZoozCall(farmPhone, 123)
	if err != nil {
		logrus.WithError(err).Error("could not call to farm")

		fallbackPhone := getFallbackPhone(farm, farmPhone)
		ok, price, err := somepkg.ZoozCall(fallbackPhone, 123)
		if err != nil {
			return 0, errors.Wrap(err, "could not call to farm")
		}
		if !ok {
			err := retryCall(3, fallbackPhone, farmPhone, farm, &price)
			if err != nil {
				return 0, errors.Wrap(err, "could not call to farm")
			}
		}
	}

	if !ok {
		err := retryCall(3, farmPhone, farmPhone, farm, &price)
		if err != nil {
			return 0, errors.Wrap(err, "could not call to farm")
		}
	}

	meetPrice, err := sunnyfarm.Farm(farm).GetMeet("Entrecote")
	if err != nil {

		// No entrecote, get Pullet

		ok, _, err := somepkg.ZoozCall(farmPhone, 666)
		if err != nil {
			logrus.WithError(err).Error("could not call to farm")

			fallbackPhone := getFallbackPhone(farm, farmPhone)
			ok, price, err := somepkg.ZoozCall(fallbackPhone, 666)
			if err != nil {
				return 0, errors.Wrap(err, "could not call to farm")
			}
			if !ok {
				err := retryCall(3, fallbackPhone, farmPhone, farm, &price)
				if err != nil {
					return 0, errors.Wrap(err, "could not call to farm")
				}
			}
		}

		if !ok {
			err := retryCall(3, farmPhone, farmPhone, farm, &meetPrice)
			if err != nil {
				return 0, errors.Wrap(err, "could not call to farm")
			}
		}

		pricePullet, err := sunnyfarm.Farm(farm).GetMeet("Pullet")
		if err != nil {
			return 0, errors.Wrap(err, "could not get pullet")
		}
		meetPrice = pricePullet
	}

	///////////////////////////
	// Get mangal /////////////
	///////////////////////////

	stores, err := somepkg.GetNearbyMakolet()
	if err != nil {
		return 0, errors.Wrap(err, "could not get stores")
	}

	store, err := chooseMakolet(stores)
	if err != nil {
		return 0, errors.Wrap(err, "could not choose store")
	}

	storePhone, err := somepkg.GetPhone(store)
	if err != nil {
		return 0, errors.Wrap(err, "could not get farm phone store")
	}

	okStore, price, err := somepkg.ZoozCall(storePhone, 555)
	if err != nil {
		logrus.WithError(err).Error("could not call to store")

		fallbackPhone := getFallbackPhone(store, storePhone)
		ok, price, err := somepkg.ZoozCall(fallbackPhone, 123)
		if err != nil {
			return 0, errors.Wrap(err, "could not call")
		}
		if !ok {
			err := retryCall(3, fallbackPhone, farmPhone, farm, &price)
			if err != nil {
				return 0, errors.Wrap(err, "could not call to store")
			}
		}
	}

	if !okStore {
		err := retryCall(3, farmPhone, farmPhone, farm, &price)
		if err != nil {
			return 0, errors.Wrap(err, "could not call to store")
		}
	}

	mangalPrice, err := sunnyfarm.Farm(farm).GetMeet("Entrecote")
	if err != nil {

		// No entrecote, get Pullet

		ok, _, err := somepkg.ZoozCall(storePhone, 666)
		if err != nil {
			logrus.WithError(err).Error("could not call to store")

			fallbackPhone := getFallbackPhone(farm, farmPhone)
			ok, price, err := somepkg.ZoozCall(fallbackPhone, 666)
			if err != nil {
				return 0, errors.Wrap(err, "could not call to store")
			}
			if !ok {
				err := retryCall(3, fallbackPhone, storePhone, store, &price)
				if err != nil {
					return 0, errors.Wrap(err, "could not call to store")
				}
			}
		}

		if !ok {
			return 0, errors.Wrap(err, "could not call to farm")
		}
	}

	///////////////////////////
	// Get coal ///////////////
	///////////////////////////

	//...
	coalPrice := 100

	return meetPrice + mangalPrice + coalPrice, nil
}

func retryCall(i int, s string, s2 string, s3 string, dd *int) error {
	return nil
}

func getFallbackPhone(s string, s2 string) string {
	return ""
}

func chooseFarm(farms []string) (string, error) {
	return "", nil
}

func chooseMakolet(farms []string) (string, error) {
	return "", nil
}
