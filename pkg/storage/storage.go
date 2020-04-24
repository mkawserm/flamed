package storage

type Storage struct {
}

func (s *Storage) Create(namespace []byte, key []byte, value []byte) (bool, error) {
	return true, nil
}

func (s *Storage) Update(namespace []byte, key []byte, value []byte) (bool, error) {
	return true, nil
}

func (s *Storage) Delete(namespace []byte, key []byte) (bool, error) {
	return true, nil
}

func (s *Storage) Read(namespace []byte, key []byte) ([]byte, error) {
	return nil, nil
}

func (s *Storage) CreateMeta(namespace []byte, data []byte) (bool, error) {
	return false, nil
}

func (s *Storage) UpdateMeta(namespace []byte, data []byte) (bool, error) {
	return false, nil
}

func (s *Storage) DeleteMeta(namespace []byte) (bool, error) {
	return false, nil
}

func (s *Storage) ReadMeta(namespace []byte) ([]byte, error) {
	return nil, nil
}
