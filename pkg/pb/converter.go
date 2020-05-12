package pb

import "github.com/golang/protobuf/proto"

func (m *StateSnapshot) ToStateEntry() *StateEntry {
	entry := &StateEntry{}
	if err := proto.Unmarshal(m.Data, entry); err != nil {
		return nil
	}
	return entry
}
