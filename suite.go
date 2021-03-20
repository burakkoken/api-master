package apimaster

import "github.com/stretchr/testify/suite"

type Suite struct {
	suite.Suite
	Client *Client
}

func (suite *Suite) Initialize() {
	suite.Client = NewClientWithSuite(suite)
}
