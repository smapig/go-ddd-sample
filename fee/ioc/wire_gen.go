// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package ioc

import (
	"github.com/smapig/go-ddd-sample/core/domain"
	"github.com/smapig/go-ddd-sample/core/infrastructure/config"
	"github.com/smapig/go-ddd-sample/core/infrastructure/log"
	"github.com/smapig/go-ddd-sample/core/infrastructure/orm"
	"github.com/smapig/go-ddd-sample/core/service/fee"
	fee2 "github.com/smapig/go-ddd-sample/fee"
	"github.com/smapig/go-ddd-sample/fee/controller"
)

// Injectors from wire.go:

func InitializeConfig(confPath string) (config.AppConfig, error) {
	appConfig := config.NewConfigProvider(confPath)
	return appConfig, nil
}

func InitializeLogger(cfg config.AppConfig) (log.Logger, error) {
	logger := log.NewLogger(cfg)
	return logger, nil
}

func InitializeDbContext(logger log.Logger, conf config.AppConfig) (orm.DbContext, error) {
	dbContext, err := orm.NewDBContext(logger, conf)
	if err != nil {
		return nil, err
	}
	return dbContext, nil
}

func InitializeGenericRepository(dbContext orm.DbContext, logger log.Logger) (orm.UnitOfWorkRepository, error) {
	unitOfWorkRepository := domain.NewGenericRepository(dbContext, logger)
	return unitOfWorkRepository, nil
}

func InitializeFeeService(conf config.AppConfig, logger log.Logger, repo orm.UnitOfWorkRepository) (fee.FeeService, error) {
	feeService := fee.NewFeeService(conf, logger, repo)
	return feeService, nil
}

func InitializeController(conf config.AppConfig, logger log.Logger, feeService fee.FeeService) (controller.Controller, error) {
	controllerController := controller.NewController(conf, logger, feeService)
	return controllerController, nil
}

func InitializeServer(confPath string) (fee2.Server, error) {
	appConfig, err := InitializeConfig(confPath)
	if err != nil {
		return nil, err
	}
	logger, err := InitializeLogger(appConfig)
	if err != nil {
		return nil, err
	}
	dbContext, err := InitializeDbContext(logger, appConfig)
	if err != nil {
		return nil, err
	}
	unitOfWorkRepository, err := InitializeGenericRepository(dbContext, logger)
	if err != nil {
		return nil, err
	}
	feeService, err := InitializeFeeService(appConfig, logger, unitOfWorkRepository)
	if err != nil {
		return nil, err
	}
	controllerController, err := InitializeController(appConfig, logger, feeService)
	if err != nil {
		return nil, err
	}
	server := fee2.NewServer(logger, appConfig, controllerController)
	return server, nil
}
