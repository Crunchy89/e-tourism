package usecase

import (
	"api/domain"
	"api/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func generateFasilitas() *domain.Fasilitas {
	data := &domain.Fasilitas{
		ID:            primitive.NewObjectID(),
		ForeignID:     primitive.NewObjectID(),
		NamaFasilitas: "",
		Jumlah:        0,
		IsDelete:      false,
		IsActive:      false,
		Log: &domain.Log{
			Created: 0,
			Updated: 0,
		},
	}
	return data
}

func TestAddFasilitas(t *testing.T) {
	fasilitasRepo := &mocks.FasilitasRepository{}
	fasilitasUseCase := &FasilitasUseCase{
		Fasilitas: fasilitasRepo,
	}
	t.Run("AddFasilitas should success and return id", func(t *testing.T) {
		fasilitasData := generateFasilitas()
		fasilitasRepo.Mock.On("Add", mock.Anything, fasilitasData).Return(fasilitasData.ID, nil)
		res, err := fasilitasUseCase.AddFasilitas(fasilitasData)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

}
