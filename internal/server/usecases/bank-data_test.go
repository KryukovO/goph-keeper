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

func TestAddBankData(t *testing.T) {
	data1 := entities.BankData{
		UserID:         1,
		Number:         "1234",
		CardholderName: "Name",
		ExpiredAt:      "01/24",
		CVV:            "123",
		Metadata:       "Metadata",
	}

	type args struct {
		data entities.BankData
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockBankDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct adding",
			prepare: func(mock *mocks.MockBankDataRepository) {
				mock.EXPECT().AddBankData(gomock.Any(), data1).Return(nil)
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
			prepare: func(mock *mocks.MockBankDataRepository) {
				mock.EXPECT().AddBankData(gomock.Any(), data1).Return(entities.ErrBankDataAlreadyExists)
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
		repo := mocks.NewMockBankDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewBankDataUseCase(repo, time.Minute)

		err := uc.AddBankData(context.Background(), test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestUpdateBankData(t *testing.T) {
	var (
		data1 = entities.BankData{
			UserID:         1,
			Number:         "1234",
			CardholderName: "Name",
			ExpiredAt:      "01/24",
			CVV:            "123",
			Metadata:       "Metadata",
		}
		data2 = entities.BankData{
			UserID:         1,
			Number:         "1234",
			CardholderName: "New name",
			ExpiredAt:      "01/24",
			CVV:            "123",
			Metadata:       "New metadata",
		}
	)

	type args struct {
		oldNumber string
		data      entities.BankData
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockBankDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct updating",
			prepare: func(mock *mocks.MockBankDataRepository) {
				mock.EXPECT().UpdateBankData(gomock.Any(), data1.Number, data2).Return(nil)
			},
			args: args{
				oldNumber: data1.Number,
				data:      data2,
			},
			wants: wants{
				wantErr: false,
			},
		},
		{
			name: "Data already exists",
			prepare: func(mock *mocks.MockBankDataRepository) {
				mock.EXPECT().UpdateBankData(gomock.Any(), data1.Number, data2).Return(entities.ErrBankDataAlreadyExists)
			},
			args: args{
				oldNumber: data1.Number,
				data:      data2,
			},
			wants: wants{
				wantErr: true,
			},
		},
	}

	for _, test := range tests {
		repo := mocks.NewMockBankDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewBankDataUseCase(repo, time.Minute)

		err := uc.UpdateBankData(context.Background(), test.args.oldNumber, test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestDeleteBankData(t *testing.T) {
	data1 := entities.BankData{
		UserID:         1,
		Number:         "1234",
		CardholderName: "Name",
		ExpiredAt:      "01/24",
		CVV:            "123",
		Metadata:       "Metadata",
	}

	type args struct {
		data entities.BankData
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockBankDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct deleting",
			prepare: func(mock *mocks.MockBankDataRepository) {
				mock.EXPECT().DeleteBankData(gomock.Any(), data1).Return(nil)
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
		repo := mocks.NewMockBankDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewBankDataUseCase(repo, time.Minute)

		err := uc.DeleteBankData(context.Background(), test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestBankCardNumbersList(t *testing.T) {
	data1 := entities.BankData{
		UserID:         1,
		Number:         "1234",
		CardholderName: "Name",
		ExpiredAt:      "01/24",
		CVV:            "123",
		Metadata:       "Metadata",
	}

	type args struct {
		userID int64
	}

	type wants struct {
		list    []string
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockBankDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct",
			prepare: func(mock *mocks.MockBankDataRepository) {
				mock.EXPECT().BankCardNumbersList(gomock.Any(), data1.UserID).Return([]string{data1.Number}, nil)
			},
			args: args{
				userID: data1.UserID,
			},
			wants: wants{
				list:    []string{data1.Number},
				wantErr: false,
			},
		},
		{
			name: "Data not found",
			prepare: func(mock *mocks.MockBankDataRepository) {
				mock.EXPECT().BankCardNumbersList(gomock.Any(), data1.UserID).Return([]string{}, nil)
			},
			args: args{
				userID: data1.UserID,
			},
			wants: wants{
				list:    []string{},
				wantErr: false,
			},
		},
	}

	for _, test := range tests {
		repo := mocks.NewMockBankDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewBankDataUseCase(repo, time.Minute)

		list, err := uc.BankCardNumbersList(context.Background(), test.args.userID)
		if test.wants.wantErr {
			assert.Error(t, err)
			assert.Empty(t, list)
		} else {
			assert.NoError(t, err)
			assert.ElementsMatch(t, test.wants.list, list)
		}
	}
}

func TestBankCard(t *testing.T) {
	data1 := entities.BankData{
		UserID:         1,
		Number:         "1234",
		CardholderName: "Name",
		ExpiredAt:      "01/24",
		CVV:            "123",
		Metadata:       "Metadata",
	}

	type args struct {
		data *entities.BankData
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockBankDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct",
			prepare: func(mock *mocks.MockBankDataRepository) {
				mock.EXPECT().BankCard(gomock.Any(), &data1).Return(nil)
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
		repo := mocks.NewMockBankDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewBankDataUseCase(repo, time.Minute)

		err := uc.BankCard(context.Background(), test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
