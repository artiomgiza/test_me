package boris_tests

import (
	"errors"
	"testing"

	mock_beeffarm "github.com/artiomgiza/test_me/pkgs/beef-farm/mock"
	mock_chickenfarm "github.com/artiomgiza/test_me/pkgs/chicken-farm/mock"
	mock_coalfarm "github.com/artiomgiza/test_me/pkgs/coal-farm/mock"
	mock_mangalstore "github.com/artiomgiza/test_me/pkgs/mangal-store/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var ctrl *gomock.Controller
var mockBeef *mock_beeffarm.MockProvider
var mockChicken *mock_chickenfarm.MockProvider
var mockCoal *mock_coalfarm.MockProvider
var mockMangal *mock_mangalstore.MockProvider

func setup(t *testing.T) func() {
	ctrl = gomock.NewController(t)
	mockBeef = mock_beeffarm.NewMockProvider(ctrl)
	mockChicken = mock_chickenfarm.NewMockProvider(ctrl)
	mockMangal = mock_mangalstore.NewMockProvider(ctrl)
	mockCoal = mock_coalfarm.NewMockProvider(ctrl)
	Instance = coolPriceCalculator{
		beefFarm:    mockBeef,
		chickenFarm: mockChicken,
		mangalStore: mockMangal,
		coalFarm:    mockCoal,
	}
	return func() {
		ctrl.Finish()
	}
}

func TestCoolPriceCalculator_CalculatePrice(t *testing.T) {
	t.Run("WhenTooManyEaters", func(t *testing.T) {
		defer setup(t)()
		amount, err := Instance.CalculatePrice(42)
		assert.Error(t, err)
		assert.Equal(t, 0, amount)
	})
	t.Run("WhenTooFewEaters", func(t *testing.T) {
		defer setup(t)()
		amount, err := Instance.CalculatePrice(-42)
		assert.Error(t, err)
		assert.Equal(t, 0, amount)
	})
	t.Run("WhenAllGood", func(t *testing.T) {
		const (
			counter     = 8
			price       = 10
			mangalPrice = 42
			coalPrice   = 12
		)
		t.Run("WithEntrecote", func(t *testing.T) {
			mockBeef.EXPECT().GetEntrecote(counter).Times(1).Return(counter*price, nil)
			// The next line is not needed but it is expressive and can help to understand the flow
			mockChicken.EXPECT().GetPullet(gomock.Any()).Times(0)
			// Next 2 rows can be moved to a separate function because they are the same in all cases of WhenAllGood part
			// I'd do this if it is more complex than 2 rows, now the duplication is minimal and I keep it here
			mockMangal.EXPECT().GetMangal().Times(1).Return(mangalPrice, nil)
			mockCoal.EXPECT().GetCoal(1).Return(coalPrice, nil)
			amount, err := Instance.CalculatePrice(counter)
			assert.Nil(t, err)
			assert.Equal(t, 134, amount)
		})
		t.Run("WithChicken", func(t *testing.T) {
			mockBeef.EXPECT().GetEntrecote(counter).Times(1).Return(0, errors.New("No entrecote available"))
			mockChicken.EXPECT().GetPullet(gomock.Any()).Times(1).Return(counter*price, nil)
			mockMangal.EXPECT().GetMangal().Times(1).Return(mangalPrice, nil)
			mockCoal.EXPECT().GetCoal(1).Return(coalPrice, nil)
			amount, err := Instance.CalculatePrice(counter)
			assert.Nil(t, err)
			assert.Equal(t, 134, amount)
		})
	})
}
