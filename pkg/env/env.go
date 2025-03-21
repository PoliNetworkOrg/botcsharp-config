package env

import (
	"fmt"
	"os"
	"strconv"
)


type Env struct {
}

func (e *Env) GetStringPanic(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("The required env variable %s was not set", key))
	}

	return value
}

func (e *Env) GetString(key string, fallback *string) *string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return &value
}

func (e *Env) GetBool(key string, defaultValue *bool) *bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return &boolValue
}

func (e *Env) GetInt(key string, defaultValue *int64) *int64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}

	return &intValue
}
