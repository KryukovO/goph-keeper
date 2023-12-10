package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/KryukovO/goph-keeper/internal/entities"
	"github.com/KryukovO/goph-keeper/internal/server/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegistration(t *testing.T) {
	var (
		user1 = entities.User{
			Login:        "user1",
			Password:     "1234",
			Subscription: entities.RegularSubscription,
		}
		secret   = []byte("secret")
		tokenTTL = 30 * time.Minute
	)

	type args struct {
		user entities.User
	}

	type wants struct {
		userID  int64
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockUserRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct registration",
			prepare: func(mock *mocks.MockUserRepository) {
				mock.EXPECT().CreateUser(gomock.Any(), user1).Return(int64(1), nil)
			},
			args: args{
				user: user1,
			},
			wants: wants{
				userID:  int64(1),
				wantErr: false,
			},
		},
		{
			name: "User already exists",
			prepare: func(mock *mocks.MockUserRepository) {
				mock.EXPECT().CreateUser(gomock.Any(), user1).Return(int64(0), entities.ErrUserAlreadyExists)
			},
			args: args{
				user: user1,
			},
			wants: wants{
				userID:  int64(0),
				wantErr: true,
			},
		},
	}

	for _, test := range tests {
		repo := mocks.NewMockUserRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		user := NewUserUseCase(repo, time.Minute, secret, tokenTTL)

		userID, token, err := user.Registration(context.Background(), test.args.user)
		if test.wants.wantErr {
			assert.Error(t, err)
			assert.Empty(t, userID)
			assert.Empty(t, token)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.wants.userID, userID)
			assert.NotEmpty(t, token)
		}
	}
}

func TestLogin(t *testing.T) {
	var (
		user1 = entities.User{
			ID:       1,
			Login:    "user1",
			Password: "1234",
		}
		secret   = []byte("secret")
		tokenTTL = 30 * time.Minute
	)

	type args struct {
		user entities.User
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockUserRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct login",
			prepare: func(mock *mocks.MockUserRepository) {
				mock.EXPECT().User(gomock.Any(), &user1).Return(nil)
			},
			args: args{
				user: user1,
			},
			wants: wants{
				wantErr: false,
			},
		},
		{
			name: "Invalid login/password",
			prepare: func(mock *mocks.MockUserRepository) {
				mock.EXPECT().User(gomock.Any(), &user1).Return(entities.ErrInvalidLoginPassword)
			},
			args: args{
				user: user1,
			},
			wants: wants{
				wantErr: true,
			},
		},
	}

	for _, test := range tests {
		repo := mocks.NewMockUserRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		user := NewUserUseCase(repo, time.Minute, secret, tokenTTL)

		token, err := user.Authorization(context.Background(), test.args.user)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, token)
		}
	}
}
