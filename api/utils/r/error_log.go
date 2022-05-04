package r

import (
	"errors"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Ex interface {
	IsDatabaseError() bool
	IsRepositoryError() bool
	IsUseCaseError() bool
	IsDataNotFound() bool
	Error() string
}

type errs struct {
	isDatabaseError   bool
	isRepositoryError bool
	isUseCaseError    bool
	isDataNotFound    bool
	Message           string
}

func (e *errs) IsDatabaseError() bool {
	return e.isDatabaseError
}

func (e *errs) IsRepositoryError() bool {
	return e.isRepositoryError
}

func (e *errs) IsUseCaseError() bool {
	return e.isUseCaseError
}

func (e *errs) IsDataNotFound() bool {
	return e.isDataNotFound
}

func newDBError(message string) errs {
	return errs{
		isDatabaseError:   true,
		isRepositoryError: false,
		isUseCaseError:    false,
		isDataNotFound:    false,
		Message:           message,
	}
}

func newRPError(message string) errs {
	return errs{
		isDatabaseError:   false,
		isRepositoryError: true,
		isUseCaseError:    false,
		isDataNotFound:    false,
		Message:           message,
	}
}

func newUCError(message string) errs {
	return errs{
		isDatabaseError:   false,
		isRepositoryError: false,
		isUseCaseError:    true,
		isDataNotFound:    false,
		Message:           message,
	}
}

func newNFError(message string) errs {
	return errs{
		isDatabaseError:   false,
		isRepositoryError: false,
		isUseCaseError:    false,
		isDataNotFound:    true,
		Message:           message,
	}
}

func newOError(message string) errs {
	return errs{
		isDatabaseError:   false,
		isRepositoryError: false,
		isUseCaseError:    false,
		isDataNotFound:    false,
		Message:           message,
	}
}

func NewErrorDataNotFound(coll string, err error) Ex {
	return &ErrorRepository{
		errs:       newNFError(err.Error()),
		Collection: coll,
	}
}

func NewErrorMongo(coll string, err error) Ex {
	if err == mongo.ErrNoDocuments {
		return NewErrorDataNotFound(coll, err)
	}
	return &ErrorRepository{
		errs:       newRPError(err.Error()),
		Collection: coll,
	}
}

func NewErrorRepository(coll string, err error) Ex {
	return &ErrorRepository{
		errs:       newRPError(err.Error()),
		Collection: coll,
	}
}

type ErrorRepository struct {
	errs
	Collection string
}

func (e *ErrorRepository) Error() string {
	return fmt.Sprintf("[ERROR]RP Collection (%s) : %s", e.Collection, e.Message)
}

func NewErrorUseCase(err error) Ex {
	return &ErrorUseCase{newUCError(err.Error())}
}

type ErrorUseCase struct {
	errs
}

func (e *ErrorUseCase) Error() string {
	return fmt.Sprintf("[ERROR]USECASE : %s", e.Message)
}

func NewErrorDatabase(err error) Ex {
	return &ErrorDatabase{newDBError(err.Error())}
}

type ErrorDatabase struct {
	errs
}

func (e *ErrorDatabase) Error() string {
	return fmt.Sprintf("[ERROR]DB : %s", e.Message)
}

func NewBodyParseError(err error) Ex {
	return &ErrorResult{
		errs:   newOError(err.Error()),
		Status: http.StatusUnprocessableEntity,
	}
}

func NewErr(err string) Ex {
	return NewErrorUseCase(errors.New(err))
}

type ErrorResult struct {
	errs
	Status int32
}

func (e *ErrorResult) Error() string {
	return fmt.Sprintf("[ERROR]Result (%d) : %s", e.Status, e.Message)
}
