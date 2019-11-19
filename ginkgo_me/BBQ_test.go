package ginkgo_me

import (
	"errors"
	mock_beeffarm "test_me/pkgs/beef-farm/mock"
	mock_chickenfarm "test_me/pkgs/chicken-farm/mock"
	mock_coalfarm "test_me/pkgs/coal-farm/mock"
	mock_mangalstore "test_me/pkgs/mangal-store/mock"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBBQ(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BBQ suit")
}

var _ = Describe("calc", func() {

	var (
		mockCtrl        *gomock.Controller
		mockBeefFarm    *mock_beeffarm.MockProvider
		mockChickenFarm *mock_chickenfarm.MockProvider
		mockMangalStore *mock_mangalstore.MockProvider
		mockCoatFarm    *mock_coalfarm.MockProvider

		subject coolPriceCalculator
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
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
	})

	Describe(".Prepare", func() {
		var (
			// inputs
			peopleCount int
			// outputs
			totalPrice int
			err        error
		)

		JustBeforeEach(func() {
			totalPrice, err = subject.CalculatePrice(peopleCount)
		})

		Context("when peopleCounter is exceeded max value", func() {
			BeforeEach(func() {
				peopleCount = 11
			})

			It("should return error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("too much people"))
			})
		})

		Context("when peopleCounter is not exceeded max value", func() {
			BeforeEach(func() {
				peopleCount = 8
			})

			Context("when GetEntrecote returns error", func() {
				BeforeEach(func() {
					mockBeefFarm.EXPECT().GetEntrecote(peopleCount).Return(0, errors.New("test_err_7483"))
				})

				Context("when GetChicken returns error", func() {
					BeforeEach(func() {
						mockChickenFarm.EXPECT().GetPullet(peopleCount).Return(0, errors.New("test_err_2393"))
					})

					It("should return error", func() {
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(ContainSubstring("no meet available"))
						Expect(err.Error()).To(ContainSubstring("test_err_2393"))
					})
				})
			})

			Context("when GetEntrecote returns no error", func() {
				var meetPrice int

				BeforeEach(func() {
					meetPrice = 100
					mockBeefFarm.EXPECT().GetEntrecote(peopleCount).Return(meetPrice, nil)
				})

				Context("when GetMangal returns error", func() {
					BeforeEach(func() {
						mockMangalStore.EXPECT().GetMangal().Return(0, errors.New("test_err_9494"))
					})

					It("should return error", func() {
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(ContainSubstring("no mangal available"))
						Expect(err.Error()).To(ContainSubstring("test_err_9494"))
					})
				})

				Context("when GetMangal returns no error", func() {
					var mangalPrice int

					BeforeEach(func() {
						mangalPrice = 1000
						mockMangalStore.EXPECT().GetMangal().Return(mangalPrice, nil)
					})

					Context("when GetCoal returns error", func() {
						BeforeEach(func() {
							mockCoatFarm.EXPECT().GetCoal(1).Return(0, errors.New("test_err_9883"))
						})

						It("should return error", func() {
							Expect(err).To(HaveOccurred())
							Expect(err.Error()).To(ContainSubstring("no coal available"))
							Expect(err.Error()).To(ContainSubstring("test_err_9883"))
						})
					})

					Context("when GetCoal returns no error", func() {
						var coatPrice int

						BeforeEach(func() {
							coatPrice = 10
							mockCoatFarm.EXPECT().GetCoal(1).Return(coatPrice, nil)
						})

						It("should return no error", func() {
							Expect(err).To(Succeed())
							Expect(totalPrice).To(Equal(1110))
						})
					})
				})
			})
		})
	})
})
