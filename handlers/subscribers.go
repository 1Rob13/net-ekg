package subscribers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/1Rob13/net-ekg/db"
	"github.com/1Rob13/net-ekg/models"
	"github.com/1Rob13/net-ekg/restapi/operations"
)

type SubscriberHandler struct {
	Saver db.Saver
}

func New(saver db.Saver) *SubscriberHandler {

	return &SubscriberHandler{Saver: saver}
}

func (s *SubscriberHandler) HandleGet(params operations.GetSubscribersParams) middleware.Responder {

	users, err := s.Saver.RetrieveAllUsers()

	if err != nil {
		return operations.NewGetSubscribersBadRequest().WithPayload(&models.Error{Message: err.Error()})
	}

	// this can be the place to implement
	return operations.NewGetSubscribersOK().WithPayload(users)
}

func (s *SubscriberHandler) HandlePost(params operations.PostSubscribersParams) middleware.Responder {

	err := s.Saver.Save(*params.User)

	if err != nil {
		return operations.NewGetSubscribersBadRequest().WithPayload(&models.Error{Message: err.Error()})
	}

	return operations.NewPostSubscribersCreated().WithPayload(params.User)
}
