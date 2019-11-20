package gotestplainme

import (
	mock_beeffarm "github.com/artiomgiza/test_me/pkgs/beef-farm/mock"
	mock_chickenfarm "github.com/artiomgiza/test_me/pkgs/chicken-farm/mock"
	mock_coalfarm "github.com/artiomgiza/test_me/pkgs/coal-farm/mock"
	mock_mangalstore "github.com/artiomgiza/test_me/pkgs/mangal-store/mock"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mockBeefFarm    *mock_beeffarm.MockProvider
	mockChickenFarm *mock_chickenfarm.MockProvider
	mockMangalStore *mock_mangalstore.MockProvider
	mockCoatFarm    *mock_coalfarm.MockProvider

	subject coolPriceCalculator
)

func initMocks(t *testing.T) (finish func()) {
	mockCtrl := gomock.NewController(t)
	mockBeefFarm = mock_beeffarm.NewMockProvider(mockCtrl)
	mockChickenFarm = mock_chickenfarm.NewMockProvider(mockCtrl)
	mockMangalStore = mock_mangalstore.NewMockProvider(mockCtrl)
	mockCoatFarm = mock_coalfarm.NewMockProvider(mockCtrl)

	subject = coolPriceCalculator{
		beefFarm:    mockBeefFarm,
		chickenFarm: mockChickenFarm,
		mangalStore: mockMangalStore,
		coalFarm:    mockCoatFarm,
	}
	return func() {
		mockCtrl.Finish()
	}
}

func Test_coolPriceCalculator_CalculatePrice_Gotest(t *testing.T) {
	t.Run("when people counter is exceeded max value", func(t *testing.T) {
		_, err := subject.CalculatePrice(11)
		assert.NotNil(t, err)
	})

	t.Run("when get chicken returns error", func(t *testing.T) {
		defer initMocks(t)()
		peopleCounter := 10
		mockBeefFarm.EXPECT().GetEntrecote(peopleCounter).Return(0, errors.New("test_err_7483"))
		mockChickenFarm.EXPECT().GetPullet(peopleCounter).Return(0, errors.New("test_err_2393"))
		_, err := subject.CalculatePrice(peopleCounter)
		assert.NotNil(t, err)
	})

	t.Run("when get mangal returns error", func(t *testing.T) {
		defer initMocks(t)()
		peopleCounter := 10
		meetPrice := 100
		mockBeefFarm.EXPECT().GetEntrecote(peopleCounter).Return(meetPrice, nil)

		mockMangalStore.EXPECT().GetMangal().Return(0, errors.New("test_err_9494"))
		_, err := subject.CalculatePrice(peopleCounter)
		assert.NotNil(t, err, )
		assert.Contains(t, err.Error(), "no mangal available", "should return error about no mangal available")
		assert.Contains(t, err.Error(), "test_err_9494")
	})

	t.Run("when get coal returns error", func(t *testing.T) {
		defer initMocks(t)()
		peopleCounter := 10
		meetPrice := 100
		mangalPrice := 1000
		mockBeefFarm.EXPECT().GetEntrecote(peopleCounter).Return(meetPrice, nil)
		mockMangalStore.EXPECT().GetMangal().Return(mangalPrice, nil)
		mockCoatFarm.EXPECT().GetCoal(1).Return(0, errors.New("test_err_9883"))
		_, err := subject.CalculatePrice(peopleCounter)
		assert.NotNil(t, err, )
		assert.Contains(t, err.Error(), "no coal available", "should return error about no coal available")
		assert.Contains(t, err.Error(), "test_err_9883")
	})

	t.Run("when get coal returns no error", func(t *testing.T) {
		defer initMocks(t)()
		coatPrice := 10
		peopleCounter := 10
		meetPrice := 100
		mangalPrice := 1000

		mockBeefFarm.EXPECT().GetEntrecote(peopleCounter).Return(meetPrice, nil)
		mockMangalStore.EXPECT().GetMangal().Return(mangalPrice, nil)
		mockCoatFarm.EXPECT().GetCoal(1).Return(coatPrice, nil)

		got, err := subject.CalculatePrice(peopleCounter)
		assert.Nil(t, err)
		expectedTotalPrice := meetPrice + mangalPrice + coatPrice
		assert.Equal(t, expectedTotalPrice, got)
	})
}
