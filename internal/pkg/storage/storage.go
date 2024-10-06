package storage

import (
	"go.uber.org/zap"
)

type Storage struct {
	inner  map[string]string //change to value
	logger *zap.Logger
}

func NewStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}

	defer logger.Sync()
	logger.Info("Created new storage")

	return Storage{
		inner:  make(map[string]string),
		logger: logger,
	}, nil
}

func (r Storage) Set(key, value string) {
	r.inner[key] = value
	r.logger.Info("Added key for value")
	r.logger.Sync()
}

func (r Storage) Get(key string) *string {
	res, ok := r.inner[key]
	r.logger.Info("Returned value", zap.String("key", key), zap.String("value", res))
	r.logger.Sync()
	if !ok {
		return nil
	}
	return &res
}
