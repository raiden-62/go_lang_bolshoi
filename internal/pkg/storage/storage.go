package storage

import (
	"strconv"

	"go.uber.org/zap"
)

type Variable struct {
	type_of_val string

	str     string
	integer int
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

	if digit, err := strconv.Atoi(value); err == nil {
		r.inner[key] = Variable{integer: digit, type_of_val: "D"}
	} else {
		r.inner[key] = Variable{str: value, type_of_val: "S"}
	}

	r.logger.Info("Added key for value") //add extra info
	r.logger.Sync()
}

func (r Storage) Get(key string) interface{} {
	value_struct, ok := r.inner[key]

	r.logger.Info("Returned value", zap.String("key", key)) //add the value of what was added
	r.logger.Sync()

	if !ok {
		return nil
	}
	if value_struct.type_of_val == "D" {
		return value_struct.integer
	} else if value_struct.type_of_val == "S" {
		return value_struct.str
	}
	return nil //should this be here?
}

func (r Storage) GetKind(key string) string {
	value_struct, ok := r.inner[key]

	r.logger.Info("Returned type of value", zap.String("type", value_struct.type_of_val))
	if !ok {
		return "somethings wrong" //this is NOT the way fs
	}
	return value_struct.type_of_val

}
