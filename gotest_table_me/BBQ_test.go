package gotesttableme

import (
	mock_beeffarm "github.com/artiomgiza/test_me/pkgs/beef-farm/mock"
	mock_chickenfarm "github.com/artiomgiza/test_me/pkgs/chicken-farm/mock"
	mock_coalfarm "github.com/artiomgiza/test_me/pkgs/coal-farm/mock"
	mock_mangalstore "github.com/artiomgiza/test_me/pkgs/mangal-store/mock"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"testing"
)


func Test_coolPriceCalculator_CalculatePrice(t *testing.T) {
	type args struct {
		peopleCounter int
	}
	type providerArgs struct {
		weight int
		price  int
		err    error
	}
	tests := []struct {
		name        string
		args        args
		beefArgs    *providerArgs
		chickenArgs *providerArgs
		mangalArgs  *providerArgs
		coatArgs    *providerArgs
		want        int
		wantErr     bool
	}{
		{
			name:    "when people counter is exceeded max value",
			args:    args{peopleCounter: 11},
			wantErr: true,
		},
		{
			name: "when get chicken returns error",
			args: args{peopleCounter: 10},
			beefArgs: &providerArgs{
				weight: 10,
				err:    errors.New("test_err_7483"),
			},
			chickenArgs: &providerArgs{
				weight: 10,
				err:    errors.New("test_err_2393"),
			},
			wantErr: true,
		},
		{
			name:     "when get mangal returns error",
			args:     args{peopleCounter: 10},
			beefArgs: &providerArgs{weight: 10, price: 100},
			mangalArgs: &providerArgs{
				weight: 10,
				err:    errors.New("test_err_9494"),
			},
			wantErr: true,
		},
		{
			name:       "when get coal returns error",
			args:       args{peopleCounter: 10},
			beefArgs:   &providerArgs{weight: 10, price: 100},
			mangalArgs: &providerArgs{weight: 10, price: 1000},
			coatArgs:   &providerArgs{err: errors.New("test_err_9883"),},
			wantErr:    true,
		},
		{
			name:       "when get coal returns no error",
			args:       args{peopleCounter: 10},
			beefArgs:   &providerArgs{weight: 10, price: 100},
			mangalArgs: &providerArgs{weight: 10, price: 1000},
			coatArgs:   &providerArgs{price: 10},
			want:       1110,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockBeefFarm := mock_beeffarm.NewMockProvider(mockCtrl)
			if tt.beefArgs != nil {
				mockBeefFarm.EXPECT().GetEntrecote(tt.beefArgs.weight).Return(tt.beefArgs.price, tt.beefArgs.err)
			}
			mockChickenFarm := mock_chickenfarm.NewMockProvider(mockCtrl)
			if tt.chickenArgs != nil {
				mockChickenFarm.EXPECT().GetPullet(tt.chickenArgs.weight).Return(tt.chickenArgs.weight, tt.chickenArgs.err)
			}
			mockMangalStore := mock_mangalstore.NewMockProvider(mockCtrl)
			if tt.mangalArgs != nil {
				mockMangalStore.EXPECT().GetMangal().Return(tt.mangalArgs.price, tt.mangalArgs.err)
			}
			mockCoatFarm := mock_coalfarm.NewMockProvider(mockCtrl)
			if tt.coatArgs != nil {
				mockCoatFarm.EXPECT().GetCoal(1).Return(tt.coatArgs.price, tt.coatArgs.err)
			}
			subject := coolPriceCalculator{
				beefFarm:    mockBeefFarm,
				chickenFarm: mockChickenFarm,
				mangalStore: mockMangalStore,
				coalFarm:    mockCoatFarm,
			}
			got, err := subject.CalculatePrice(tt.args.peopleCounter)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculatePrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculatePrice() got = %v, want %v", got, tt.want)
			}
		})
	}

}