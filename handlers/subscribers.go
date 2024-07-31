package subscribers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/1Rob13/net-ekg/db"
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

		return operations.NewGetSubscribersOK()
	}

	// this can be the place to implement
	return operations.NewGetSubscribersOK().WithPayload(users)
}

func (s *SubscriberHandler) HandlePost(params operations.PostSubscribersParams) middleware.Responder {

	panic("not implemented")

	//read the request

}
