package core

import "bytes"

type Store interface {
	Create(parentID []byte, value []byte) ([]byte, error)
	Read(ID []byte) ([]byte, error)
	Update(ID []byte, value []byte) ([]byte, error)
	Delete(ID []byte) error
	Seek(parent []byte) ([][]byte, [][]byte, error)
}

type Transport interface {
	Serve(server *BlocksServer) error
}

type Block struct {
	ID       []byte   `json:"id"`
	Value    []byte   `json:"value"`
	Children []*Block `json:"children"`
}

type BlocksServer struct {
	store     Store
	transport Transport
}

func NewServer(store Store, transport Transport) *BlocksServer {
	return &BlocksServer{store, transport}
}

func (s *BlocksServer) Serve() {
	s.transport.Serve(s)
}

func (s *BlocksServer) Get(ID []byte) (*Block, error) {
	block := &Block{}

	if ID != nil {
		value, err := s.store.Read(ID)

		if err != nil {
			return nil, err
		}

		block.ID = ID
		block.Value = value
	}

	ids, values, err := s.store.Seek(ID)

	if err != nil {
		return nil, err
	}

	for i, id := range ids {
		if bytes.Equal(id, ID) {
			continue
		}

		block.Children = append(block.Children, &Block{id, values[i], []*Block{}})
	}

	return block, nil
}

func (s *BlocksServer) Create(parentID []byte, block *Block) (*Block, error) {
	id, err := s.store.Create(parentID, block.Value)

	if err != nil {
		return nil, err
	}

	block.ID = id

	return block, nil
}

func (s *BlocksServer) Update(ID []byte, block *Block) (*Block, error) {
	if _, err := s.store.Update(ID, block.Value); err != nil {
		return nil, err
	}

	return s.Get(ID)
}
