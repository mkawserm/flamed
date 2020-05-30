package utility

import (
	"bytes"
	"encoding/json"
	"regexp"
)

var keyMatchRegex = regexp.MustCompile(`"(\w+)":`)

type LowerCamelCaseMarshaller struct {
	Value interface{}
}

func (s LowerCamelCaseMarshaller) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(s.Value)

	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			match[1] = bytes.ToLower(match[1:2])[0]

			data := bytes.Split(match, []byte("_"))
			if len(data) != 0 {
				var newMatch []byte
				for i, datum := range data {
					if i == 0 {
						newMatch = append(newMatch, datum...)
					} else {
						newMatch = append(newMatch, bytes.Title(datum)...)
					}
				}
				return newMatch
			} else {
				return match
			}
		},
	)

	return converted, err
}

var keyMatchRegex2 = regexp.MustCompile(`"(\w+)":`)
var wordBarrierRegex2 = regexp.MustCompile(`(\w)([A-Z])`)

type LowerSnakeCaseMarshaller struct {
	Value interface{}
}

func (c LowerSnakeCaseMarshaller) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)

	converted := keyMatchRegex2.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex2.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)

	return converted, err
}
