package usecases

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/KryukovO/goph-keeper/internal/entities"
	fsMocks "github.com/KryukovO/goph-keeper/internal/server/filestorage/mocks"
	repoMocks "github.com/KryukovO/goph-keeper/internal/server/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddBinaryData(t *testing.T) {
	var (
		data1 = entities.File{
			UserID:   1,
			FileName: "File",
			Data:     *bytes.NewBuffer([]byte{}),
		}
		subscriptions = map[int64]entities.Subscription{
			data1.UserID: entities.RegularSubscription,
		}
	)

	type args struct {
		data entities.File
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name        string
		prepareRepo func(mock *repoMocks.MockSubscriptionRepository)
		prepareFS   func(mock *fsMocks.MockFileStorage)
		args        args
		wants       wants
	}{
		{
			name: "Correct adding",
			prepareRepo: func(mock *repoMocks.MockSubscriptionRepository) {
				mock.EXPECT().Sunbsriptions(gomock.Any()).Return(subscriptions, nil)
			},
			prepareFS: func(mock *fsMocks.MockFileStorage) {
				mock.EXPECT().SetSubscriptions(subscriptions)
				mock.EXPECT().Save(data1).Return(nil)
			},
			args: args{
				data: data1,
			},
			wants: wants{
				wantErr: false,
			},
		},
		{
			name: "File is too big",
			prepareRepo: func(mock *repoMocks.MockSubscriptionRepository) {
				mock.EXPECT().Sunbsriptions(gomock.Any()).Return(subscriptions, nil)
			},
			prepareFS: func(mock *fsMocks.MockFileStorage) {
				mock.EXPECT().SetSubscriptions(subscriptions)
				mock.EXPECT().Save(data1).Return(entities.ErrFileIsTooBig)
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
		repo := repoMocks.NewMockSubscriptionRepository(gomock.NewController(t))

		if test.prepareRepo != nil {
			test.prepareRepo(repo)
		}

		fs := fsMocks.NewMockFileStorage(gomock.NewController(t))

		if test.prepareFS != nil {
			test.prepareFS(fs)
		}

		uc, err := NewBinaryDataUseCase(context.Background(), repo, fs, time.Minute)

		assert.NoError(t, err)

		err = uc.AddBinaryData(context.Background(), test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestDeleteBinaryData(t *testing.T) {
	var (
		data1 = entities.File{
			UserID:   1,
			FileName: "File",
			Data:     *bytes.NewBuffer([]byte{}),
		}
		subscriptions = map[int64]entities.Subscription{
			data1.UserID: entities.RegularSubscription,
		}
	)

	type args struct {
		data entities.File
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name        string
		prepareRepo func(mock *repoMocks.MockSubscriptionRepository)
		prepareFS   func(mock *fsMocks.MockFileStorage)
		args        args
		wants       wants
	}{
		{
			name: "Correct deleting",
			prepareRepo: func(mock *repoMocks.MockSubscriptionRepository) {
				mock.EXPECT().Sunbsriptions(gomock.Any()).Return(subscriptions, nil)
			},
			prepareFS: func(mock *fsMocks.MockFileStorage) {
				mock.EXPECT().SetSubscriptions(subscriptions)
				mock.EXPECT().Delete(data1).Return(nil)
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
		repo := repoMocks.NewMockSubscriptionRepository(gomock.NewController(t))

		if test.prepareRepo != nil {
			test.prepareRepo(repo)
		}

		fs := fsMocks.NewMockFileStorage(gomock.NewController(t))

		if test.prepareFS != nil {
			test.prepareFS(fs)
		}

		uc, err := NewBinaryDataUseCase(context.Background(), repo, fs, time.Minute)

		assert.NoError(t, err)

		err = uc.DeleteBinaryData(context.Background(), test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestFileNamesList(t *testing.T) {
	var (
		data1 = entities.File{
			UserID:   1,
			FileName: "File",
			Data:     *bytes.NewBuffer([]byte{}),
		}
		subscriptions = map[int64]entities.Subscription{
			data1.UserID: entities.RegularSubscription,
		}
	)

	type args struct {
		userID int64
	}

	type wants struct {
		list []string
	}

	tests := []struct {
		name        string
		prepareRepo func(mock *repoMocks.MockSubscriptionRepository)
		prepareFS   func(mock *fsMocks.MockFileStorage)
		args        args
		wants       wants
	}{
		{
			name: "Correct",
			prepareRepo: func(mock *repoMocks.MockSubscriptionRepository) {
				mock.EXPECT().Sunbsriptions(gomock.Any()).Return(subscriptions, nil)
			},
			prepareFS: func(mock *fsMocks.MockFileStorage) {
				mock.EXPECT().SetSubscriptions(subscriptions)
				mock.EXPECT().List(data1.UserID).Return([]string{data1.FileName})
			},
			args: args{
				userID: data1.UserID,
			},
			wants: wants{
				list: []string{data1.FileName},
			},
		},
	}

	for _, test := range tests {
		repo := repoMocks.NewMockSubscriptionRepository(gomock.NewController(t))

		if test.prepareRepo != nil {
			test.prepareRepo(repo)
		}

		fs := fsMocks.NewMockFileStorage(gomock.NewController(t))

		if test.prepareFS != nil {
			test.prepareFS(fs)
		}

		uc, err := NewBinaryDataUseCase(context.Background(), repo, fs, time.Minute)

		assert.NoError(t, err)

		list := uc.FileNamesList(context.Background(), test.args.userID)

		assert.ElementsMatch(t, test.wants.list, list)
	}
}

func TestBinaryData(t *testing.T) {
	var (
		data1 = entities.File{
			UserID:   1,
			FileName: "File",
			Data:     *bytes.NewBuffer([]byte{}),
		}
		subscriptions = map[int64]entities.Subscription{
			data1.UserID: entities.RegularSubscription,
		}
	)

	type args struct {
		data *entities.File
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name        string
		prepareRepo func(mock *repoMocks.MockSubscriptionRepository)
		prepareFS   func(mock *fsMocks.MockFileStorage)
		args        args
		wants       wants
	}{
		{
			name: "Correct",
			prepareRepo: func(mock *repoMocks.MockSubscriptionRepository) {
				mock.EXPECT().Sunbsriptions(gomock.Any()).Return(subscriptions, nil)
			},
			prepareFS: func(mock *fsMocks.MockFileStorage) {
				mock.EXPECT().SetSubscriptions(subscriptions)
				mock.EXPECT().Load(&data1).Return(nil)
			},
			args: args{
				data: &data1,
			},
			wants: wants{
				wantErr: false,
			},
		},
	}

	for _, test := range tests {
		repo := repoMocks.NewMockSubscriptionRepository(gomock.NewController(t))

		if test.prepareRepo != nil {
			test.prepareRepo(repo)
		}

		fs := fsMocks.NewMockFileStorage(gomock.NewController(t))

		if test.prepareFS != nil {
			test.prepareFS(fs)
		}

		uc, err := NewBinaryDataUseCase(context.Background(), repo, fs, time.Minute)

		assert.NoError(t, err)

		err = uc.BinaryData(context.Background(), test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestUpdateSubscription(t *testing.T) {
	var (
		userID        = int64(1)
		subscription  = entities.PremiumSubscription
		subscriptions = map[int64]entities.Subscription{
			userID: entities.RegularSubscription,
		}
	)

	type args struct {
		userID       int64
		subscription entities.Subscription
	}

	tests := []struct {
		name        string
		prepareRepo func(mock *repoMocks.MockSubscriptionRepository)
		prepareFS   func(mock *fsMocks.MockFileStorage)
		args        args
	}{
		{
			name: "Correct",
			prepareRepo: func(mock *repoMocks.MockSubscriptionRepository) {
				mock.EXPECT().Sunbsriptions(gomock.Any()).Return(subscriptions, nil)
			},
			prepareFS: func(mock *fsMocks.MockFileStorage) {
				mock.EXPECT().SetSubscriptions(subscriptions)
				mock.EXPECT().UpdateSubscription(userID, subscription)
			},
			args: args{
				userID:       userID,
				subscription: subscription,
			},
		},
	}

	for _, test := range tests {
		repo := repoMocks.NewMockSubscriptionRepository(gomock.NewController(t))

		if test.prepareRepo != nil {
			test.prepareRepo(repo)
		}

		fs := fsMocks.NewMockFileStorage(gomock.NewController(t))

		if test.prepareFS != nil {
			test.prepareFS(fs)
		}

		uc, err := NewBinaryDataUseCase(context.Background(), repo, fs, time.Minute)

		assert.NoError(t, err)

		uc.UpdateSubscription(context.Background(), test.args.userID, test.args.subscription)
	}
}
