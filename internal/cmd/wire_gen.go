// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/waite-lee/sgn/internal/apiserver"
	"github.com/waite-lee/sgn/internal/common"
	"github.com/waite-lee/sgn/pkg/app"
	"github.com/waite-lee/sgn/pkg/server"
)

// Injectors from wire_inject.go:

func BuildCommands(ac *app.AppContext) AppCommands {
	testCmd := NewTestCmd()
	serverContext := server.NewServerContext()
	apiServer := apiserver.NewApiServer(serverContext, ac)
	apiServerOptions := _wireApiServerOptionsValue
	apiServerCmd := NewApiServerCmd(apiServer, apiServerOptions)
	nftPullCmd := NewNftPullCmd()
	configCmd := NewConfigCmd()
	httpClientOptions := _wireHttpClientOptionsValue
	updateCmd := NewUpdateCmd(httpClientOptions)
	appCommands := AppCommands{
		Test:      testCmd,
		ApiServer: apiServerCmd,
		NftPull:   nftPullCmd,
		Config:    configCmd,
		Update:    updateCmd,
	}
	return appCommands
}

var (
	_wireApiServerOptionsValue  = apiserver.AsOptions
	_wireHttpClientOptionsValue = common.HttpOptions
)
