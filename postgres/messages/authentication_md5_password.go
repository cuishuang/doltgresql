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

import "github.com/dolthub/doltgresql/postgres/connection"

func init() {
	connection.InitializeDefaultMessage(AuthenticationMD5Password{})
}

// AuthenticationMD5Password represents a PostgreSQL message.
type AuthenticationMD5Password struct {
	Salt int32
}

var authenticationMD5PasswordDefault = connection.MessageFormat{
	Name: "AuthenticationMD5Password",
	Fields: connection.FieldGroup{
		{
			Name:  "Header",
			Type:  connection.Byte1,
			Flags: connection.Header,
			Data:  int32('R'),
		},
		{
			Name:  "MessageLength",
			Type:  connection.Int32,
			Flags: connection.MessageLengthInclusive,
			Data:  int32(12),
		},
		{
			Name: "Status",
			Type: connection.Int32,
			Data: int32(5),
		},
		{
			Name: "Salt",
			Type: connection.Byte4,
			Data: int32(0),
		},
	},
}

var _ connection.Message = AuthenticationMD5Password{}

// Encode implements the interface connection.Message.
func (m AuthenticationMD5Password) Encode() (connection.MessageFormat, error) {
	outputMessage := m.DefaultMessage().Copy()
	outputMessage.Field("Salt").MustWrite(m.Salt)
	return outputMessage, nil
}

// Decode implements the interface connection.Message.
func (m AuthenticationMD5Password) Decode(s connection.MessageFormat) (connection.Message, error) {
	if err := s.MatchesStructure(*m.DefaultMessage()); err != nil {
		return nil, err
	}
	return AuthenticationMD5Password{
		Salt: s.Field("Salt").MustGet().(int32),
	}, nil
}

// DefaultMessage implements the interface connection.Message.
func (m AuthenticationMD5Password) DefaultMessage() *connection.MessageFormat {
	return &authenticationMD5PasswordDefault
}
