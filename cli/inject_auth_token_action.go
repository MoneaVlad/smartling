package main

import (
	"github.com/mdreizin/smartling/model"
	"github.com/mdreizin/smartling/service"
	"gopkg.in/urfave/cli.v1"
)

func injectAuthTokenAction(c *cli.Context) (err error) {
	authToken := &model.AuthToken{}
	container := c.App.Metadata[containerKey].(*service.Container)
	projectConfig := c.App.Metadata[projectConfigKey].(*model.ProjectConfig)

	if projectConfig.AuthToken.AccessToken != "" {
		if authToken, err = container.AuthService.Refresh(projectConfig.AuthToken.AccessToken); err != nil {
			authToken, err = container.AuthService.Authenticate(&projectConfig.UserToken)
		}
	} else {
		authToken, err = container.AuthService.Authenticate(&projectConfig.UserToken)
	}

	c.App.Metadata[authTokenKey] = authToken

	return err
}
