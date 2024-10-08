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
	defer r.logger.Sync()
	if digit, err := strconv.Atoi(value); err == nil {
		r.inner[key] = Variable{integer: digit, type_of_val: "D"}
		r.logger.Info("Added value", zap.String("key", key), zap.String("value", strconv.Itoa(digit)), zap.String("type", "D"))
	} else {
		r.inner[key] = Variable{str: value, type_of_val: "S"}
		r.logger.Info("Added value", zap.String("key", key), zap.String("value", value), zap.String("type", "S"))
	}

}

func (r Storage) Get(key string) interface{} { //return pointer to the struct?
	value_struct, ok := r.inner[key]

	defer r.logger.Sync()

	if !ok {
		return nil
	}
	if value_struct.type_of_val == "D" {
		r.logger.Info("Returned value", zap.String("key", key), zap.String("value", strconv.Itoa(value_struct.integer)))
		return value_struct.integer
	} else if value_struct.type_of_val == "S" {
		r.logger.Info("Returned value", zap.String("key", key), zap.String("value", value_struct.str))
		return value_struct.str
	} else {
		return nil //should this be here?
	}
}

func (r Storage) GetKind(key string) string {
	value_struct, ok := r.inner[key]

	r.logger.Info("Returned type of value", zap.String("type", value_struct.type_of_val))
	r.logger.Sync()

	if !ok {
		return "key not found" //this is probably not how it's done
	}
	return value_struct.type_of_val

}
