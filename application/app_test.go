package application

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/masashi1014/gomock_sample/repository"
	mocks "github.com/masashi1014/gomock_sample/repository/mocks"
)

func Test_FindOrCreateUser(t *testing.T) {
	tests := []struct {
		name string
		id int64
		err error

		app func(ctrl *gomock.Controller) *App
	}{
		{
			name:"IsRegistered = true",
			id: 1,
			err: nil,
			app: func(ctrl *gomock.Controller) *App {
				user := &repository.User{
					ID: 1,
					Name: "DummyUser",
					IsRegistered: true,
				}
				urMock := mocks.NewMockUserRepository(ctrl)
				urMock.EXPECT().GetUserByID(int64(1)).Return(user)

				return &App{
					UserRepository:urMock,
				}
			},
		},
		{
			name:"IsRegistered = false",
			id: 1,
			err: nil,
			app: func(ctrl *gomock.Controller) *App {
				user := &repository.User{
					ID: 1,
					Name: "DummyUser",
					IsRegistered: false,
				}
				urMock := mocks.NewMockUserRepository(ctrl)
				urMock.EXPECT().GetUserByID(int64(1)).Return(user)

				urMock.EXPECT().Create().Return(user,nil)

				return &App{
					UserRepository:urMock,
				}
			},
		},
	}

	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			app := tt.app(ctrl)
			err := app.FindOrCreateUser(tt.id)
			if err != tt.err {
				fmt.Printf("expected = %v, actual = %v", tt.err, err)
			}
		})
	}
}

