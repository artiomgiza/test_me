package gotest_me

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

	t.Run("calc", func(t *testing.T) {
		t.Run("when people counter is exceeded max value", func(t *testing.T) {
			defer initMocks(t)()
			_, err := subject.CalculatePrice(11)
			assert.NotNil(t, err)
		})
		t.Run("when people counter is not exceeded max value",func(t *testing.T) {
			defer initMocks(t)()
			t.Run("when get entrecote returns error", func(t *testing.T) {
				peopleCounter := 10
				mockBeefFarm.EXPECT().GetEntrecote(peopleCounter).Return(0, errors.New("test_err_7483")).Times(1)

				t.Run("when get chicken returns error", func(t *testing.T) {
					mockChickenFarm.EXPECT().GetPullet(peopleCounter).Return(0, errors.New("test_err_2393"))
					_, err := subject.CalculatePrice(peopleCounter)
					assert.NotNil(t, err)
				})
				t.Run("when get chicken returns no error", func(t *testing.T) {
					t.Run("when mangal return error", func(t *testing.T) {
						// In gincko test this text
						// this case is not tested
						// flow is the same as next "when get entrecote returns no error"
						// but I think we need test this case too. Because we must be sure that meetPrice was overwritten
					})
				})
			})

			t.Run("when get entrecote returns no error", func(t *testing.T) {
				defer initMocks(t)()
				peopleCounter := 10
				meetPrice := 100
				mockBeefFarm.EXPECT().GetEntrecote(peopleCounter).Return(meetPrice, nil).AnyTimes()
				t.Run("when get mangal returns error", func(t *testing.T) {
					mockMangalStore.EXPECT().GetMangal().Return(0, errors.New("test_err_9494")).Times(1)

					t.Run("should return error about no mangal available", func(t *testing.T) {
						_, err := subject.CalculatePrice(peopleCounter)
						assert.NotNil(t, err, )
						assert.Contains(t, err.Error(), "no mangal available")
						assert.Contains(t, err.Error(), "test_err_9494")
					})
				})

				t.Run("when get mangal returns no error", func(t *testing.T) {
					mangalPrice := 1000
					mockMangalStore.EXPECT().GetMangal().Return(mangalPrice, nil).Times(2)

					t.Run("when get coal returns error", func(t *testing.T) {
						mockCoatFarm.EXPECT().GetCoal(1).Return(0, errors.New("test_err_9883"))

						t.Run("should return error about no coal available", func(t *testing.T) {
							_, err := subject.CalculatePrice(peopleCounter)
							assert.NotNil(t, err, )
							assert.Contains(t, err.Error(), "no coal available")
							assert.Contains(t, err.Error(), "test_err_9883")
						})
					})

					t.Run("when get coal returns no error", func(t *testing.T) {
						coatPrice := 10
						mockCoatFarm.EXPECT().GetCoal(1).Return(coatPrice, nil)

						t.Run("should return no error and valid total price", func(t *testing.T) {
							got, err := subject.CalculatePrice(peopleCounter)
							assert.Nil(t, err)
							expectedTotalPrice := meetPrice + mangalPrice + coatPrice
							assert.Equal(t, expectedTotalPrice, got)
						})
					})
				})
			})
		})
	})
}
