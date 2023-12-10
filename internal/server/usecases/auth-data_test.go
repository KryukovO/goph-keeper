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

func TestAddAuthData(t *testing.T) {
	data1 := entities.AuthData{
		UserID:   1,
		Resource: "Resource",
		Login:    "User1",
		Password: "Password",
		Metadata: "Metadata",
	}

	type args struct {
		data entities.AuthData
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockAuthDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct adding",
			prepare: func(mock *mocks.MockAuthDataRepository) {
				mock.EXPECT().AddAuthData(gomock.Any(), data1).Return(nil)
			},
			args: args{
				data: data1,
			},
			wants: wants{
				wantErr: false,
			},
		},
		{
			name: "Data already exists",
			prepare: func(mock *mocks.MockAuthDataRepository) {
				mock.EXPECT().AddAuthData(gomock.Any(), data1).Return(entities.ErrAuthDataAlreadyExists)
			},
			args: args{
				data: data1,
			},
			wants: wants{
				wantErr: true,
			},
		},
	}

	for _, test := range tests {
		repo := mocks.NewMockAuthDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewAuthDataUseCase(repo, time.Minute)

		err := uc.AddAuthData(context.Background(), test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestUpdateAuthData(t *testing.T) {
	var (
		data1 = entities.AuthData{
			UserID:   1,
			Resource: "Resource",
			Login:    "User1",
			Password: "Password",
			Metadata: "Metadata",
		}
		data2 = entities.AuthData{
			UserID:   1,
			Resource: "Resource",
			Login:    "User1",
			Password: "NewPassword",
			Metadata: "New Metadata",
		}
	)

	type args struct {
		oldResource string
		oldLogin    string
		data        entities.AuthData
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockAuthDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct updating",
			prepare: func(mock *mocks.MockAuthDataRepository) {
				mock.EXPECT().UpdateAuthData(gomock.Any(), data1.Resource, data1.Login, data2).Return(nil)
			},
			args: args{
				oldResource: data1.Resource,
				oldLogin:    data1.Login,
				data:        data2,
			},
			wants: wants{
				wantErr: false,
			},
		},
		{
			name: "Data already exists",
			prepare: func(mock *mocks.MockAuthDataRepository) {
				mock.EXPECT().UpdateAuthData(
					gomock.Any(), data1.Resource, data1.Login, data2,
				).Return(entities.ErrAuthDataAlreadyExists)
			},
			args: args{
				oldResource: data1.Resource,
				oldLogin:    data1.Login,
				data:        data2,
			},
			wants: wants{
				wantErr: true,
			},
		},
	}

	for _, test := range tests {
		repo := mocks.NewMockAuthDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewAuthDataUseCase(repo, time.Minute)

		err := uc.UpdateAuthData(context.Background(), test.args.oldResource, test.args.oldLogin, test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestDeleteAuthData(t *testing.T) {
	data1 := entities.AuthData{
		UserID:   1,
		Resource: "Resource",
		Login:    "User1",
		Password: "Password",
		Metadata: "Metadata",
	}

	type args struct {
		data entities.AuthData
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockAuthDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct deleting",
			prepare: func(mock *mocks.MockAuthDataRepository) {
				mock.EXPECT().DeleteAuthData(gomock.Any(), data1).Return(nil)
			},
			args: args{
				data: data1,
			},
			wants: wants{
				wantErr: false,
			},
		},
	}

	for _, test := range tests {
		repo := mocks.NewMockAuthDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewAuthDataUseCase(repo, time.Minute)

		err := uc.DeleteAuthData(context.Background(), test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestAuthDataList(t *testing.T) {
	data1 := entities.AuthData{
		UserID:   1,
		Resource: "Resource",
		Login:    "User1",
		Password: "Password",
		Metadata: "Metadata",
	}

	type args struct {
		userID int64
	}

	type wants struct {
		list    []entities.AuthData
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockAuthDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct",
			prepare: func(mock *mocks.MockAuthDataRepository) {
				mock.EXPECT().AuthDataList(gomock.Any(), data1.UserID).Return([]entities.AuthData{data1}, nil)
			},
			args: args{
				userID: data1.UserID,
			},
			wants: wants{
				list:    []entities.AuthData{data1},
				wantErr: false,
			},
		},
		{
			name: "Data not found",
			prepare: func(mock *mocks.MockAuthDataRepository) {
				mock.EXPECT().AuthDataList(gomock.Any(), data1.UserID).Return([]entities.AuthData{}, nil)
			},
			args: args{
				userID: data1.UserID,
			},
			wants: wants{
				list:    []entities.AuthData{},
				wantErr: false,
			},
		},
	}

	for _, test := range tests {
		repo := mocks.NewMockAuthDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewAuthDataUseCase(repo, time.Minute)

		list, err := uc.AuthDataList(context.Background(), test.args.userID)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.ElementsMatch(t, test.wants.list, list)
		}
	}
}
