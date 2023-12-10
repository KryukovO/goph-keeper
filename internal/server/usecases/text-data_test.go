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

func TestAddTextData(t *testing.T) {
	data1 := entities.TextData{
		UserID:   1,
		Label:    "Label1",
		Data:     "Text",
		Metadata: "Metadata",
	}

	type args struct {
		data entities.TextData
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockTextDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct adding",
			prepare: func(mock *mocks.MockTextDataRepository) {
				mock.EXPECT().AddTextData(gomock.Any(), data1).Return(nil)
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
			prepare: func(mock *mocks.MockTextDataRepository) {
				mock.EXPECT().AddTextData(gomock.Any(), data1).Return(entities.ErrTextDataAlreadyExists)
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
		repo := mocks.NewMockTextDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewTextDataUseCase(repo, time.Minute)

		err := uc.AddTextData(context.Background(), test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestUpdateTextData(t *testing.T) {
	var (
		data1 = entities.TextData{
			UserID:   1,
			Label:    "Label1",
			Data:     "Text",
			Metadata: "Metadata",
		}
		data2 = entities.TextData{
			UserID:   1,
			Label:    "Label1",
			Data:     "New text",
			Metadata: "New metadata",
		}
	)

	type args struct {
		oldLabel string
		data     entities.TextData
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockTextDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct updating",
			prepare: func(mock *mocks.MockTextDataRepository) {
				mock.EXPECT().UpdateTextData(gomock.Any(), data1.Label, data2).Return(nil)
			},
			args: args{
				oldLabel: data1.Label,
				data:     data2,
			},
			wants: wants{
				wantErr: false,
			},
		},
		{
			name: "Data already exists",
			prepare: func(mock *mocks.MockTextDataRepository) {
				mock.EXPECT().UpdateTextData(gomock.Any(), data1.Label, data2).Return(entities.ErrTextDataAlreadyExists)
			},
			args: args{
				oldLabel: data1.Label,
				data:     data2,
			},
			wants: wants{
				wantErr: true,
			},
		},
	}

	for _, test := range tests {
		repo := mocks.NewMockTextDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewTextDataUseCase(repo, time.Minute)

		err := uc.UpdateTextData(context.Background(), test.args.oldLabel, test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestDeleteTextData(t *testing.T) {
	data1 := entities.TextData{
		UserID:   1,
		Label:    "Label1",
		Data:     "Text",
		Metadata: "Metadata",
	}

	type args struct {
		data entities.TextData
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockTextDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct deleting",
			prepare: func(mock *mocks.MockTextDataRepository) {
				mock.EXPECT().DeleteTextData(gomock.Any(), data1).Return(nil)
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
		repo := mocks.NewMockTextDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewTextDataUseCase(repo, time.Minute)

		err := uc.DeleteTextData(context.Background(), test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestTextLabelsList(t *testing.T) {
	data1 := entities.TextData{
		UserID:   1,
		Label:    "Label1",
		Data:     "Text",
		Metadata: "Metadata",
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
		prepare func(mock *mocks.MockTextDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct",
			prepare: func(mock *mocks.MockTextDataRepository) {
				mock.EXPECT().TextLabelsList(gomock.Any(), data1.UserID).Return([]string{data1.Label}, nil)
			},
			args: args{
				userID: data1.UserID,
			},
			wants: wants{
				list:    []string{data1.Label},
				wantErr: false,
			},
		},
		{
			name: "Data not found",
			prepare: func(mock *mocks.MockTextDataRepository) {
				mock.EXPECT().TextLabelsList(gomock.Any(), data1.UserID).Return([]string{}, nil)
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
		repo := mocks.NewMockTextDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewTextDataUseCase(repo, time.Minute)

		list, err := uc.TextLabelsList(context.Background(), test.args.userID)
		if test.wants.wantErr {
			assert.Error(t, err)
			assert.Empty(t, list)
		} else {
			assert.NoError(t, err)
			assert.ElementsMatch(t, test.wants.list, list)
		}
	}
}

func TestTextData(t *testing.T) {
	data1 := entities.TextData{
		UserID:   1,
		Label:    "Label1",
		Data:     "Text",
		Metadata: "Metadata",
	}

	type args struct {
		data *entities.TextData
	}

	type wants struct {
		wantErr bool
	}

	tests := []struct {
		name    string
		prepare func(mock *mocks.MockTextDataRepository)
		args    args
		wants   wants
	}{
		{
			name: "Correct",
			prepare: func(mock *mocks.MockTextDataRepository) {
				mock.EXPECT().TextData(gomock.Any(), &data1).Return(nil)
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
		repo := mocks.NewMockTextDataRepository(gomock.NewController(t))

		if test.prepare != nil {
			test.prepare(repo)
		}

		uc := NewTextDataUseCase(repo, time.Minute)

		err := uc.TextData(context.Background(), test.args.data)
		if test.wants.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
