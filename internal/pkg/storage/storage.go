package storage

import (
	"strconv"

	"go.uber.org/zap"
)

const (
	KindInt       = "D"
	KindStr       = "S"
	KindUndefined = "UN"
)

type Kind string

type Variable struct {
	v string
	t Kind
}

type Storage struct {
	inner  map[string]Variable
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
		inner:  make(map[string]Variable),
		logger: logger,
	}, nil
}

func (r Storage) Set(key, value string) {
	defer r.logger.Sync()
	
	switch kind := getType(value); kind{
	case KindInt:
		r.inner[key] = Variable{v: value, t: kind}
	case KindStr:
		r.inner[key] = Variable{v: value, t: kind}
	case KindUndefined:
		r.logger.Error(
			"undefined value type",
			zap.String("key", key),
			zap.Any("value", value),
		)
	}
	r.logger.Info("Added value by key", zap.String("key", key), zap.Any("value", value))
	r.logger.Sync()

}

func (r Storage) Get(key string) *string {
	res, ok := r.get(key)

	if !ok {
		return nil
	}

	return &res.v
}

func (r Storage) get(key string) (Value, bool){
	res, ok := r.inner[key]

	if !ok {
		return Variable{}, false
	}

	return res, true
}

func (r Storage) GetKind(key string) Kind {
	value_kind, ok := getType(key)

	r.logger.Info("Returned type of value")
	r.logger.Sync()

	if !ok {
		return KindUndefined
	}
	return value_kind

}

func getType(value string) Kind{
	var val any

	val, err := strconv.Atoi(value)
	if err != nil {
		val = value
	}

	switch val.(type) {
	case int:
		return KindInt
	case string:
		return KindStr
	default:
		return KindUndefined
	}
}
