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
	//rewrite to switch case
	if digit, err := strconv.Atoi(value); err == nil {
		r.inner[key] = Variable{v: strconv.Itoa(digit), t: KindInt}
		r.logger.Info("Added value", zap.String("key", key), zap.String("value", strconv.Itoa(digit)), zap.String("type", "D"))
	} else {
		r.inner[key] = Variable{v: value, t: KindStr}
		r.logger.Info("Added value", zap.String("key", key), zap.String("value", value), zap.String("type", "S"))
	}

}

func (r Storage) Get(key string) interface{} { //return pointer to string
	//also, why should it be Get(key string) *string
	//do I return integer values as *string too?????
	value_struct, ok := r.inner[key]

	defer r.logger.Sync()

	if !ok {
		return nil
	}
	if value_struct.t == KindInt {
		r.logger.Info("Returned value", zap.String("key", key), zap.String("value", value_struct.v))
		return value_struct.v
	} else if value_struct.t == KindStr {
		r.logger.Info("Returned value", zap.String("key", key), zap.String("value", value_struct.v))
		return value_struct.v
	} else {
		return nil //should this be here?
	}
}

//add private get()

func (r Storage) GetKind(key string) Kind {
	value_struct, ok := r.inner[key]

	r.logger.Info("Returned type of value") //,zap.String("type", value_struct.t))
	r.logger.Sync()

	if !ok {
		return "key not found" //this is probably not how it's done
	}
	return value_struct.t

}
