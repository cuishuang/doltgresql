// Copyright 2023 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package messages

func init() {
	initializeDefaultMessage(BindComplete{})
	addMessageHeader(BindComplete{})
}

// BindComplete represents a PostgreSQL message.
type BindComplete struct{}

var bindCompleteDefault = Message{
	Name: "BindComplete",
	Fields: []*Field{
		{
			Name: "Header",
			Type: Byte1,
			Tags: Header,
			Data: int32('2'),
		},
		{
			Name: "MessageLength",
			Type: Int32,
			Tags: MessageLengthInclusive,
			Data: int32(4),
		},
	},
}

var _ MessageType = BindComplete{}

// encode implements the interface MessageType.
func (m BindComplete) encode() (Message, error) {
	return m.defaultMessage().Copy(), nil
}

// decode implements the interface MessageType.
func (m BindComplete) decode(s Message) (MessageType, error) {
	if err := s.MatchesStructure(*m.defaultMessage()); err != nil {
		return nil, err
	}
	return BindComplete{}, nil
}

// defaultMessage implements the interface MessageType.
func (m BindComplete) defaultMessage() *Message {
	return &bindCompleteDefault
}
