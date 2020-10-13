package ginkgo_me

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

func Test_coolPriceCalculator_CalculatePrice(t *testing.T) {

	t.Run("calc", func(t *testing.T) {
		t.Run("when people counter is exceeded max value", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			b := coolPriceCalculator{}
			_, err := b.CalculatePrice(11)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(),"too much people", "should return error about people exceeded max value")
		})
		t.Run("when people counter is not exceeded max value",func(t *testing.T) {
			t.Run("when get entrecote returns error", func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				peopleCounter := 10
				mockBeefFarm := mock_beeffarm.NewMockProvider(ctrl)
				mockBeefFarm.EXPECT().GetEntrecote(peopleCounter).Return(0, errors.New("test_err_7483")).Times(1)
				t.Run("when get chicken returns error", func(t *testing.T) {
					mockChickenFarm := mock_chickenfarm.NewMockProvider(ctrl)
					b := coolPriceCalculator{
						beefFarm:    mockBeefFarm,
						chickenFarm: mockChickenFarm,
					}
					mockChickenFarm.EXPECT().GetPullet(peopleCounter).Return(0, errors.New("test_err_2393"))
					_, err := b.CalculatePrice(peopleCounter)
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
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				peopleCounter := 10
				meetPrice := 100
				mockBeefFarm := mock_beeffarm.NewMockProvider(ctrl)
				mockBeefFarm.EXPECT().GetEntrecote(peopleCounter).Return(meetPrice, nil).AnyTimes()
				t.Run("when get mangal returns error", func(t *testing.T) {
					mockMangalStore := mock_mangalstore.NewMockProvider(ctrl)
					mockMangalStore.EXPECT().GetMangal().Return(0, errors.New("test_err_9494")).AnyTimes()

					t.Run("should return error about no mangal available", func(t *testing.T) {
						b := coolPriceCalculator{
							beefFarm:    mockBeefFarm,
							mangalStore: mockMangalStore,
						}
						_, err := b.CalculatePrice(peopleCounter)
						assert.NotNil(t, err, )
						assert.Contains(t, err.Error(), "no mangal available")
						assert.Contains(t, err.Error(), "test_err_9494")
					})
				})

				t.Run("when get mangal returns no error", func(t *testing.T) {
					mangalPrice := 1000
					mockMangalStore := mock_mangalstore.NewMockProvider(ctrl)
					mockMangalStore.EXPECT().GetMangal().Return(mangalPrice, nil).AnyTimes()

					t.Run("when get coal returns error", func(t *testing.T) {
						mockCoatFarm := mock_coalfarm.NewMockProvider(ctrl)
						mockCoatFarm.EXPECT().GetCoal(1).Return(0, errors.New("test_err_9883"))

						t.Run("should return error about no coal available", func(t *testing.T) {
							b := coolPriceCalculator{
								beefFarm:    mockBeefFarm,
								mangalStore: mockMangalStore,
								coalFarm:    mockCoatFarm,
							}

							_, err := b.CalculatePrice(peopleCounter)
							assert.NotNil(t, err, )
							assert.Contains(t, err.Error(), "no coal available")
							assert.Contains(t, err.Error(), "test_err_9883")
						})
					})

					t.Run("when get coal returns no error", func(t *testing.T) {
						coatPrice := 10
						mockCoatFarm := mock_coalfarm.NewMockProvider(ctrl)
						mockCoatFarm.EXPECT().GetCoal(1).Return(coatPrice, nil)

						t.Run("should return no error and valid total price", func(t *testing.T) {
							b := coolPriceCalculator{
								beefFarm:    mockBeefFarm,
								mangalStore: mockMangalStore,
								coalFarm:    mockCoatFarm,
							}

							got, err := b.CalculatePrice(peopleCounter)
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
