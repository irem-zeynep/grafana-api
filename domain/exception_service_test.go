package domain_test

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"grafana-api/domain"
	snsMock "grafana-api/mocks/infrastructure/event/sns"
	"testing"
)

type ExceptionServiceSuite struct {
	suite.Suite
	serv      domain.IExceptionService
	publisher *snsMock.IEventPublisher
	logger    *logrus.Logger
}

func (suite *ExceptionServiceSuite) SetupTest() {
	suite.logger = logrus.New()
	suite.publisher = snsMock.NewIEventPublisher(suite.T())
	suite.serv = domain.NewExceptionService(suite.publisher, suite.logger)
}

func (suite *ExceptionServiceSuite) TestSaveException() {
	//given
	ctx := context.TODO()
	msg := "Oh no an exception!"
	suite.publisher.On("SendMessage", ctx, msg).Return(nil)

	//when
	err := suite.serv.SaveException(ctx, msg)

	//then
	assert.Nil(suite.T(), err)
}

func (suite *ExceptionServiceSuite) TestSaveException_ErrorWhileSending() {
	//given
	ctx := context.TODO()
	msg := "Oh no an exception!"
	expectedErr := errors.New("expected err")
	suite.publisher.On("SendMessage", ctx, msg).Return(expectedErr)

	//when
	err := suite.serv.SaveException(ctx, msg)

	//then
	assert.Equal(suite.T(), expectedErr, err)
}

func TestExceptionServiceSuite(t *testing.T) {
	suite.Run(t, new(ExceptionServiceSuite))
}
