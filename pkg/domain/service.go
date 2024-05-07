package domain

import (
	"net/http"
	"test-driven-development/pkg/common"
	"test-driven-development/pkg/sqlstore"
)

type Service struct {
	db *sqlstore.DB
}

func NewService(db *sqlstore.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateUser(writer http.ResponseWriter, request *http.Request) {

	userRequest := new(User)

	//Read request body
	err := common.ReadFromRequestBody(request, userRequest)
	if err != nil {
		common.WriteErrToResponseBody(writer, err)
		return
	}

	user, err := s.db.CreateUser(userRequest)

	if err != nil {
		common.WriteErrToResponseBody(writer, err)
		return
	}

	//Write response
	common.WriteToResponseBody(writer, user)
}
