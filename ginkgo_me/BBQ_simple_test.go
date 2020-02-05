package ginkgo_me

import (
	mock_beeffarm "test_me/pkgs/beef-farm/mock"
	mock_chickenfarm "test_me/pkgs/chicken-farm/mock"
	mock_coalfarm "test_me/pkgs/coal-farm/mock"
	mock_mangalstore "test_me/pkgs/mangal-store/mock"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("calc", func() {

	It("< 10", func() {
		// INIT
		mockCtrl := gomock.NewController(GinkgoT())
		mockBeefFarm := mock_beeffarm.NewMockProvider(mockCtrl)
		mockChickenFarm := mock_chickenfarm.NewMockProvider(mockCtrl)
		mockMangalStore := mock_mangalstore.NewMockProvider(mockCtrl)
		mockCoatFarm := mock_coalfarm.NewMockProvider(mockCtrl)

		subject := coolPriceCalculator{
			beefFarm:    mockBeefFarm,
			chickenFarm: mockChickenFarm,
			mangalStore: mockMangalStore,
			coalFarm:    mockCoatFarm,
		}

		// PREP
		mockBeefFarm.EXPECT().GetEntrecote(8).Return(150, nil)
		mockMangalStore.EXPECT().GetMangal().Return(10, nil)
		mockCoatFarm.EXPECT().GetCoal(1).Return(30, nil)

		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		totalPrice, err := subject.CalculatePrice(8)
		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

		Expect(err).To(BeNil())
		Expect(totalPrice).To(Equal(190))

		mockCtrl.Finish()
	})

	It("> 10", func() {
		// INIT
		mockCtrl := gomock.NewController(GinkgoT())
		mockBeefFarm := mock_beeffarm.NewMockProvider(mockCtrl)
		mockChickenFarm := mock_chickenfarm.NewMockProvider(mockCtrl)
		mockMangalStore := mock_mangalstore.NewMockProvider(mockCtrl)
		mockCoatFarm := mock_coalfarm.NewMockProvider(mockCtrl)

		subject := coolPriceCalculator{
			beefFarm:    mockBeefFarm,
			chickenFarm: mockChickenFarm,
			mangalStore: mockMangalStore,
			coalFarm:    mockCoatFarm,
		}

		// PREP
		//mockBeefFarm.EXPECT().GetEntrecote(8).Return(150, nil)
		//mockMangalStore.EXPECT().GetMangal().Return(10, nil)
		//mockCoatFarm.EXPECT().GetCoal(1).Return(0, errors.New("cant do it"))

		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		_, err := subject.CalculatePrice(11)
		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		// CALL !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("too much people"))
		//Expect(totalPrice).To(Equal(190))

		mockCtrl.Finish()
	})

})
