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

package connection

import "net"

// Receive returns a Message from the given buffer, generally generated by the client in the main read loop of a
// connection.
func Receive(buffer []byte) (Message, bool, error) {
	if len(buffer) == 0 {
		return nil, false, nil
	}
	message, ok := allMessageHeaders[buffer[0]]
	if !ok {
		return nil, false, nil
	}
	outMessage, err := ReceiveInto(buffer, message)
	return outMessage, true, err
}

// ReceiveInto writes the contents of the buffer into the given Message.
func ReceiveInto[T Message](buffer []byte, message T) (out T, err error) {
	defaultMessage := message.DefaultMessage()
	fields := defaultMessage.Copy().Fields
	if err = decode(&decodeBuffer{buffer}, []FieldGroup{fields}, 1); err != nil {
		return out, err
	}
	decodedMessage, err := message.Decode(MessageFormat{defaultMessage.Name, fields, defaultMessage.info, false})
	if err != nil {
		return out, err
	}
	return decodedMessage.(T), nil
}

// Send sends the given message over the connection.
func Send(conn net.Conn, message Message) error {
	encodedMessage, err := message.Encode()
	if err != nil {
		return err
	}
	data, err := encode(encodedMessage)
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	return err
}
