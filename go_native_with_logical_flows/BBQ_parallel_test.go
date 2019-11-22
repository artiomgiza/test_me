package bbq

import (
	"testing"

	"github.com/artiomgiza/test_me/go_native_with_logical_flows/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This allows you to have a separate state for each test, in order to avoid a race condition.
func newParallelBBQTestSuite(t *testing.T) *bbqTestSuite {
	t.Parallel()
	res := &bbqTestSuite{mockCtrl: gomock.NewController(t)}
	res.mockBeefFarm = mocks.NewMockBeefFarm(res.mockCtrl)
	res.mockChickenFarm = mocks.NewMockChickenFarm(res.mockCtrl)
	res.mockMangalStore = mocks.NewMockMangalStore(res.mockCtrl)
	res.mockCoalFarm = mocks.NewMockCoalFarm(res.mockCtrl)
	res.subject = NewCoolPriceCalculator(res.mockBeefFarm, res.mockChickenFarm, res.mockMangalStore, res.mockCoalFarm)
	return res
}

type bbqTestSuite struct {
	mockCtrl        *gomock.Controller
	mockBeefFarm    *mocks.MockBeefFarm
	mockChickenFarm *mocks.MockChickenFarm
	mockMangalStore *mocks.MockMangalStore
	mockCoalFarm    *mocks.MockCoalFarm
	subject         *coolPriceCalculator
}

func (s *bbqTestSuite) finish() {
	s.mockCtrl.Finish()
}

func TestCoolPriceCalculator_Parallel_CalculatePrice_SuccessFlow(t *testing.T) {
	t.Parallel()
	const (
		peopleCounter = 10
		meatPrice     = 100
		mangalPrice   = 1000
		coalPrice     = 10000
	)
	var errBeefFarm = errors.New("not enough beef")

	t.Run("when cook beef", func(t *testing.T) {
		s := newParallelBBQTestSuite(t)
		defer s.finish()
		s.mockBeefFarm.EXPECT().GetEntrecotePrice(peopleCounter).Return(meatPrice, nil)
		s.mockMangalStore.EXPECT().GetMangalPrice().Return(mangalPrice, nil)
		s.mockCoalFarm.EXPECT().GetCoalPrice(1).Return(coalPrice, nil)

		price, err := s.subject.CalculatePrice(peopleCounter)
		require.NoError(t, err)
		assert.Equal(t, meatPrice+mangalPrice+coalPrice, price)
	})

	t.Run("when cook chicken", func(t *testing.T) {
		s := newParallelBBQTestSuite(t)
		defer s.finish()
		s.mockBeefFarm.EXPECT().GetEntrecotePrice(peopleCounter).Return(0, errBeefFarm)
		s.mockChickenFarm.EXPECT().GetPulletPrice(peopleCounter).Return(meatPrice, nil)
		s.mockMangalStore.EXPECT().GetMangalPrice().Return(mangalPrice, nil)
		s.mockCoalFarm.EXPECT().GetCoalPrice(1).Return(coalPrice, nil)

		price, err := s.subject.CalculatePrice(peopleCounter)
		require.NoError(t, err)
		assert.Equal(t, meatPrice+mangalPrice+coalPrice, price)
	})
}

func TestCoolPriceCalculator_Parallel_CalculatePrice_ErrorsFlow(t *testing.T) {
	t.Parallel()
	const randomPrice = 984576 // means that we do not care about this value
	var (
		errBeefFarm    = errors.New("not enough beef")
		errChickenFarm = errors.New("not enough chicken")
		errMangalStore = errors.New("no mangal available")
		errCoalFarm    = errors.New("not enough coal")
	)

	t.Run("when too many people", func(t *testing.T) {
		s := newParallelBBQTestSuite(t)
		defer s.finish()

		price, err := s.subject.CalculatePrice(maxPeopleCounter + 1)
		require.Error(t, err)
		assert.Zero(t, price)
	})

	t.Run("when can not get meat", func(t *testing.T) {
		const peopleCounter = 10

		s := newParallelBBQTestSuite(t)
		defer s.finish()
		s.mockBeefFarm.EXPECT().GetEntrecotePrice(peopleCounter).Return(0, errBeefFarm)
		s.mockChickenFarm.EXPECT().GetPulletPrice(peopleCounter).Return(0, errChickenFarm)

		price, err := s.subject.CalculatePrice(peopleCounter)
		require.Equal(t, errChickenFarm, errors.Cause(err))
		assert.Zero(t, price)
	})

	t.Run("when can not get mangal", func(t *testing.T) {
		const peopleCounter = 10

		s := newParallelBBQTestSuite(t)
		defer s.finish()
		s.mockBeefFarm.EXPECT().GetEntrecotePrice(peopleCounter).Return(randomPrice, nil)
		s.mockMangalStore.EXPECT().GetMangalPrice().Return(0, errMangalStore)

		price, err := s.subject.CalculatePrice(peopleCounter)
		require.Equal(t, errMangalStore, errors.Cause(err))
		assert.Zero(t, price)
	})

	t.Run("when can not get coal", func(t *testing.T) {
		const peopleCounter = 10

		s := newParallelBBQTestSuite(t)
		defer s.finish()
		s.mockBeefFarm.EXPECT().GetEntrecotePrice(peopleCounter).Return(randomPrice, nil)
		s.mockMangalStore.EXPECT().GetMangalPrice().Return(randomPrice, nil)
		s.mockCoalFarm.EXPECT().GetCoalPrice(1).Return(0, errCoalFarm)

		price, err := s.subject.CalculatePrice(peopleCounter)
		require.Equal(t, errCoalFarm, errors.Cause(err))
		assert.Zero(t, price)
	})
}
