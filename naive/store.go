package naive

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type NaiveStore struct {
	kvMap map[string][]byte
}

func NewStore() *NaiveStore {
	return &NaiveStore{kvMap: map[string][]byte{"abc": []byte("def")}}
}

func (s *NaiveStore) Create(parentID []byte, value []byte) ([]byte, error) {
	ID := uuid.New().String()

	if parentID != nil {
		s.kvMap[strings.Join([]string{string(parentID), ID}, ":")] = value
	} else {
		s.kvMap[ID] = value
	}

	return []byte(ID), nil
}

func (s *NaiveStore) Read(ID []byte) ([]byte, error) {
	for k, v := range s.kvMap {
		if bytes.HasSuffix([]byte(k), ID) {
			return v, nil
		}
	}

	return nil, fmt.Errorf("block with id %s not found", ID)
}

func (s *NaiveStore) Update(ID []byte, value []byte) ([]byte, error) {
	for k, _ := range s.kvMap {
		if bytes.HasSuffix([]byte(k), ID) {
			s.kvMap[k] = value
		}
	}
	return value, nil
}

func (s *NaiveStore) Delete(ID []byte) error {
	delete(s.kvMap, string(ID))
	return nil
}

func (s *NaiveStore) Seek(parentID []byte) ([][]byte, [][]byte, error) {
	ids := [][]byte{}
	values := [][]byte{}

	if parentID != nil {
		parentID = append(parentID, []byte(":")...)
	}

	for k, v := range s.kvMap {
		id := []byte(k)
		if parentID == nil {
			if !bytes.Contains(id, []byte(":")) {
				ids = append(ids, id)
				values = append(values, v)
			}
		} else {
			if bytes.HasPrefix(id, parentID) {
				ids = append(ids, bytes.TrimPrefix(id, parentID))
				values = append(values, v)
			}
		}
	}

	return ids, values, nil
}
