package json

import "encoding/json"

func (m *JSONPayload) ToJSONMap() map[string]interface{} {
	jsonMap := make(map[string]interface{})

	if err := json.Unmarshal(m.Payload, &jsonMap); err != nil {
		return nil
	}

	return jsonMap
}
