package usecase

import (
	"testing"

	"api/domain"
	"api/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserUseCase(t *testing.T) {
	user := &mocks.UserRepository{}
	userUseCase := &UserUseCase{
		User: user,
	}
	// testing add user using t.Run
	t.Run("AddUser", func(t *testing.T) {
		id := primitive.NewObjectID()
		resultUser := &domain.User{
			ID:       id,
			Username: "test",
			Email:    "test@test.com",
			Password: "test",
		}
		user.On("Add", mock.Anything, resultUser).Return(id, nil)
		user, err := userUseCase.AddUser(resultUser)
		assert.NoError(t, err)
		assert.Equal(t, id, user)
	})

	// testing get user by id using t.Run
	t.Run("GetUserByID", func(t *testing.T) {
		id := primitive.NewObjectID()
		resultUser := &domain.User{
			ID:       id,
			Username: "test",
			Email:    "test@test.com",
			Password: "test",
		}
		user.On("Fetch", mock.Anything, id).Return(resultUser, nil)
		user, err := userUseCase.GetUserByID(id)
		assert.NoError(t, err)
		assert.Equal(t, resultUser, user)
	})

	// testing get user by username using t.Run
	t.Run("GetUserByUsername", func(t *testing.T) {
		resultUser := &domain.User{
			ID:       primitive.NewObjectID(),
			Username: "test",
			Email:    "test@test.com",
			Password: "test",
		}
		user.On("FetchByUsername", mock.Anything, "test").Return(resultUser, nil)
		user, err := userUseCase.GetUserByUsername("test")
		assert.NoError(t, err)
		assert.Equal(t, resultUser, user)
	})
}
