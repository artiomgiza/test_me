package ginkgo_me

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	mock_beeffarm "github.com/artiomgiza/test_me/pkgs/beef-farm/mock"
	mock_chickenfarm "github.com/artiomgiza/test_me/pkgs/chicken-farm/mock"
	mock_coalfarm "github.com/artiomgiza/test_me/pkgs/coal-farm/mock"
	mock_mangalstore "github.com/artiomgiza/test_me/pkgs/mangal-store/mock"
	"github.com/golang/mock/gomock"
)

var (
	beefFarmMock    *mock_beeffarm.MockProvider
	chickenFarmMock *mock_chickenfarm.MockProvider
	mangalStoreMock *mock_mangalstore.MockProvider
	coalFarmMock    *mock_coalfarm.MockProvider

	subject Provider
)

func TestCoolPriceCalculator2(t *testing.T) {
	t.Run("#CalculatePrice", func(t *testing.T) {
		var (
			numberOfPeople int
			totalPrice     int
			beefPrice      int
			chickenPrice   int
			mangalPrice    int
			coalPrice      int
			err            error
		)

		defer setupMocks(t)(t)

		t.Run("when number of people is less then max limit", func(t *testing.T) {
			numberOfPeople = 8

			t.Run("when all farms and store respond successfully", func(t *testing.T) {
				beefPrice = 2
				mangalPrice = 3
				coalPrice = 4

				beefFarmMock.EXPECT().GetEntrecote(numberOfPeople).Return(beefPrice, nil)
				mangalStoreMock.EXPECT().GetMangal().Return(mangalPrice, nil)
				coalFarmMock.EXPECT().GetCoal(1).Return(coalPrice, nil)

				t.Run("it returns correct total price", func(t *testing.T) {
					totalPrice, err = subject.CalculatePrice(numberOfPeople)

					assert.Nil(t, err)
					assert.Equal(t, beefPrice+mangalPrice+coalPrice, totalPrice)
				})
			})

			t.Run("when beefFarm is failed", func(t *testing.T) {
				beefFarmMock.EXPECT().GetEntrecote(numberOfPeople).Times(2).Return(beefPrice, errors.New("Error"))

				t.Run("when chickenFarm is available", func(t *testing.T) {
					chickenPrice = 1
					mangalPrice = 3
					coalPrice = 4

					chickenFarmMock.EXPECT().GetPullet(numberOfPeople).Return(chickenPrice, nil)
					mangalStoreMock.EXPECT().GetMangal().Return(mangalPrice, nil)
					coalFarmMock.EXPECT().GetCoal(1).Return(coalPrice, nil)

					t.Run("it returns correct total price", func(t *testing.T) {
						totalPrice, err = subject.CalculatePrice(numberOfPeople)

						assert.Nil(t, err)
						assert.Equal(t, chickenPrice+mangalPrice+coalPrice, totalPrice)
					})
				})

				t.Run("when chickenFarm is failed", func(t *testing.T) {
					chickenFarmMock.EXPECT().GetPullet(numberOfPeople).Return(0, errors.New("Error"))

					t.Run("it returns an error", func(t *testing.T) {
						totalPrice, err = subject.CalculatePrice(numberOfPeople)

						assert.NotNil(t, err)
						assert.Equal(t, "no meet available: Error", err.Error())
					})
				})
			})

			t.Run("when mangalStore is failed", func(t *testing.T) {
				beefPrice = 2

				beefFarmMock.EXPECT().GetEntrecote(numberOfPeople).Return(beefPrice, nil)
				mangalStoreMock.EXPECT().GetMangal().Return(0, errors.New("Error"))

				t.Run("it returns an error", func(t *testing.T) {
					totalPrice, err = subject.CalculatePrice(numberOfPeople)

					assert.NotNil(t, err)
					assert.Equal(t, "no mangal available: Error", err.Error())
				})
			})

			t.Run("when coalFarm is failed", func(t *testing.T) {
				beefPrice = 2
				mangalPrice = 3

				beefFarmMock.EXPECT().GetEntrecote(numberOfPeople).Return(beefPrice, nil)
				mangalStoreMock.EXPECT().GetMangal().Return(mangalPrice, nil)
				coalFarmMock.EXPECT().GetCoal(1).Return(0, errors.New("Error"))

				t.Run("it returns an error", func(t *testing.T) {
					totalPrice, err = subject.CalculatePrice(numberOfPeople)

					assert.NotNil(t, err)
					assert.Equal(t, "no coal available: Error", err.Error())
				})
			})
		})

		t.Run("when number of people is more then max limit", func(t *testing.T) {
			numberOfPeople = 12

			t.Run("it returns an error", func(t *testing.T) {
				totalPrice, err = subject.CalculatePrice(numberOfPeople)

				assert.NotNil(t, err)
				assert.Equal(t, "too much people", err.Error())
			})
		})
	})
}

func setupMocks(t *testing.T) func(t *testing.T) {
	ctrl := gomock.NewController(t)
	beefFarmMock = mock_beeffarm.NewMockProvider(ctrl)
	chickenFarmMock = mock_chickenfarm.NewMockProvider(ctrl)
	mangalStoreMock = mock_mangalstore.NewMockProvider(ctrl)
	coalFarmMock = mock_coalfarm.NewMockProvider(ctrl)

	subject = coolPriceCalculator{
		beefFarm:    beefFarmMock,
		chickenFarm: chickenFarmMock,
		mangalStore: mangalStoreMock,
		coalFarm:    coalFarmMock,
	}

	return func(t *testing.T) {
		ctrl.Finish()
	}
}
