package ginkgo_me

import (
	mock_beeffarm "github.com/artiomgiza/test_me/pkgs/beef-farm/mock"
	mock_chickenfarm "github.com/artiomgiza/test_me/pkgs/chicken-farm/mock"
	mock_coalfarm "github.com/artiomgiza/test_me/pkgs/coal-farm/mock"
	mock_mangalstore "github.com/artiomgiza/test_me/pkgs/mangal-store/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"testing"
)

func TestCoolPriceCalculator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "coolPriceCalculator suit")
}

var _ = Describe("coolPriceCalculator", func() {
	var (
		ctrl            *gomock.Controller
		beefFarmMock    *mock_beeffarm.MockProvider
		chickenFarmMock *mock_chickenfarm.MockProvider
		mangalStoreMock *mock_mangalstore.MockProvider
		coalFarmMock    *mock_coalfarm.MockProvider

		subject Provider
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

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
	})

	Describe("#CalculatePrice", func() {
		var (
			numberOfPeople int
			totalPrice     int
			beefPrice      int
			chickenPrice   int
			mangalPrice    int
			coalPrice      int
			err            error
		)

		JustBeforeEach(func() {
			totalPrice, err = subject.CalculatePrice(numberOfPeople)
		})

		Context("when number of people is less then max limit", func() {
			BeforeEach(func() {
				numberOfPeople = 8
				beefPrice = 2
				mangalPrice = 3
				coalPrice = 4
			})

			Context("when all farms and store respond successfully", func() {
				BeforeEach(func() {
					beefFarmMock.EXPECT().GetEntrecote(numberOfPeople).Return(beefPrice, nil)
					mangalStoreMock.EXPECT().GetMangal().Return(mangalPrice, nil)
					coalFarmMock.EXPECT().GetCoal(1).Return(coalPrice, nil)
				})

				It("returns correct total price", func() {
					Expect(err).To(BeNil())
					Expect(totalPrice).To(Equal(beefPrice + mangalPrice + coalPrice))
				})
			})

			Context("when beefFarm is failed", func() {
				BeforeEach(func() {
					beefFarmMock.EXPECT().GetEntrecote(numberOfPeople).Return(0, errors.New("Error"))
				})

				Context("when chickenFarm is available", func() {
					BeforeEach(func() {
						chickenPrice = 1
						mangalPrice = 3
						coalPrice = 4

						chickenFarmMock.EXPECT().GetPullet(numberOfPeople).Return(chickenPrice, nil)
						mangalStoreMock.EXPECT().GetMangal().Return(mangalPrice, nil)
						coalFarmMock.EXPECT().GetCoal(1).Return(coalPrice, nil)
					})

					It("returns correct total price", func() {
						Expect(err).To(BeNil())
						Expect(totalPrice).To(Equal(chickenPrice + mangalPrice + coalPrice))
					})
				})

				Context("when chickenFarm is failed", func() {
					BeforeEach(func() {
						chickenFarmMock.EXPECT().GetPullet(numberOfPeople).Return(0, errors.New("Error"))
					})

					It("returns an error", func() {
						Expect(err).NotTo(BeNil())
						Expect(err.Error()).To(Equal("no meet available: Error"))
					})
				})
			})

			Context("when mangalStore is failed", func() {
				BeforeEach(func() {
					beefPrice = 2

					beefFarmMock.EXPECT().GetEntrecote(numberOfPeople).Return(beefPrice, nil)
					mangalStoreMock.EXPECT().GetMangal().Return(0, errors.New("Error"))
				})

				It("returns an error", func() {
					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal("no mangal available: Error"))
				})
			})

			Context("when coalFarm is failed", func() {
				BeforeEach(func() {
					beefPrice = 2
					mangalPrice = 3

					beefFarmMock.EXPECT().GetEntrecote(numberOfPeople).Return(beefPrice, nil)
					mangalStoreMock.EXPECT().GetMangal().Return(mangalPrice, nil)
					coalFarmMock.EXPECT().GetCoal(1).Return(0, errors.New("Error"))
				})

				It("returns an error", func() {
					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal("no coal available: Error"))
				})
			})
		})

		Context("when number of people is more then max limit", func() {
			BeforeEach(func() {
				numberOfPeople = 12
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("too much people"))
			})
		})
	})
})
