package entity

import (
	errors "golang-rest-api-kata/internal/errors/entity"
)

type Memory struct {
	Key   string
	Value string
}

func NewInMemory(key string, value string) (*Memory, error) {
	m := &Memory{
		Key:   key,
		Value: value,
	}

	err := m.Validate()
	if err != nil {
		return nil, errors.ErrInvalidEntity
	}
	return m, nil
}

func (m *Memory) Validate() error {
	if m.Key == "" || m.Value == "" {
		return errors.ErrInvalidEntity
	}
	return nil
}
