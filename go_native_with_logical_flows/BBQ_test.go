package bbq

import (
	"testing"

	"github.com/artiomgiza/test_me/go_native_with_logical_flows/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	mockBeefFarm    *mocks.MockBeefFarm
	mockChickenFarm *mocks.MockChickenFarm
	mockMangalStore *mocks.MockMangalStore
	mockCoalFarm    *mocks.MockCoalFarm
	subject         *coolPriceCalculator
)

func initMocks(t *testing.T) (finish func()) {
	mockCtrl := gomock.NewController(t)
	mockBeefFarm = mocks.NewMockBeefFarm(mockCtrl)
	mockChickenFarm = mocks.NewMockChickenFarm(mockCtrl)
	mockMangalStore = mocks.NewMockMangalStore(mockCtrl)
	mockCoalFarm = mocks.NewMockCoalFarm(mockCtrl)

	subject = NewCoolPriceCalculator(mockBeefFarm, mockChickenFarm, mockMangalStore, mockCoalFarm)

	return func() {
		mockCtrl.Finish()
	}
}

func TestCoolPriceCalculator_CalculatePrice_SuccessFlow(t *testing.T) {
	const (
		peopleCounter = 10
		meatPrice     = 100
		mangalPrice   = 1000
		coalPrice     = 10000
	)
	var errBeefFarm = errors.New("not enough beef")

	t.Run("when cook beef", func(t *testing.T) {
		defer initMocks(t)()
		mockBeefFarm.EXPECT().GetEntrecotePrice(peopleCounter).Return(meatPrice, nil)
		mockMangalStore.EXPECT().GetMangalPrice().Return(mangalPrice, nil)
		mockCoalFarm.EXPECT().GetCoalPrice(1).Return(coalPrice, nil)

		price, err := subject.CalculatePrice(peopleCounter)
		require.NoError(t, err)
		assert.Equal(t, meatPrice+mangalPrice+coalPrice, price)
	})

	t.Run("when cook chicken", func(t *testing.T) {
		defer initMocks(t)()
		mockBeefFarm.EXPECT().GetEntrecotePrice(peopleCounter).Return(0, errBeefFarm)
		mockChickenFarm.EXPECT().GetPulletPrice(peopleCounter).Return(meatPrice, nil)
		mockMangalStore.EXPECT().GetMangalPrice().Return(mangalPrice, nil)
		mockCoalFarm.EXPECT().GetCoalPrice(1).Return(coalPrice, nil)

		price, err := subject.CalculatePrice(peopleCounter)
		require.NoError(t, err)
		assert.Equal(t, meatPrice+mangalPrice+coalPrice, price)
	})
}

func TestCoolPriceCalculator_CalculatePrice_ErrorsFlow(t *testing.T) {
	const randomPrice = 984576 // means that we do not care about this value
	var (
		errBeefFarm    = errors.New("not enough beef")
		errChickenFarm = errors.New("not enough chicken")
		errMangalStore = errors.New("no mangal available")
		errCoalFarm    = errors.New("not enough coal")
	)

	t.Run("when too many people", func(t *testing.T) {
		defer initMocks(t)()

		price, err := subject.CalculatePrice(maxPeopleCounter + 1)
		require.Error(t, err)
		assert.Zero(t, price)
	})

	t.Run("when can not get meat", func(t *testing.T) {
		const peopleCounter = 10

		defer initMocks(t)()
		mockBeefFarm.EXPECT().GetEntrecotePrice(peopleCounter).Return(0, errBeefFarm)
		mockChickenFarm.EXPECT().GetPulletPrice(peopleCounter).Return(0, errChickenFarm)

		price, err := subject.CalculatePrice(peopleCounter)
		require.Equal(t, errChickenFarm, errors.Cause(err))
		assert.Zero(t, price)
	})

	t.Run("when can not get mangal", func(t *testing.T) {
		const peopleCounter = 10

		defer initMocks(t)()
		mockBeefFarm.EXPECT().GetEntrecotePrice(peopleCounter).Return(randomPrice, nil)
		mockMangalStore.EXPECT().GetMangalPrice().Return(0, errMangalStore)

		price, err := subject.CalculatePrice(peopleCounter)
		require.Equal(t, errMangalStore, errors.Cause(err))
		assert.Zero(t, price)
	})

	t.Run("when can not get coal", func(t *testing.T) {
		const peopleCounter = 10

		defer initMocks(t)()
		mockBeefFarm.EXPECT().GetEntrecotePrice(peopleCounter).Return(randomPrice, nil)
		mockMangalStore.EXPECT().GetMangalPrice().Return(randomPrice, nil)
		mockCoalFarm.EXPECT().GetCoalPrice(1).Return(0, errCoalFarm)

		price, err := subject.CalculatePrice(peopleCounter)
		require.Equal(t, errCoalFarm, errors.Cause(err))
		assert.Zero(t, price)
	})
}
