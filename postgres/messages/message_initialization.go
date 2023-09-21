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

import "fmt"

// allMessageHeaders contains any message headers that should be read within the main read loop of a connection.
var allMessageHeaders = make(map[byte]MessageType)

// allMessageNames contains the names of all messages, as they should all be unique.
var allMessageNames = make(map[string]struct{})

// allMessageDefaults contains all of the default message pointers, to make sure that they're not accidentally being reused.
var allMessageDefaults = make(map[*Message]struct{})

// addMessageHeader adds the given MessageType's header. This also ensures that each header is unique. This should be
// called in an init() function.
func addMessageHeader(message MessageType) {
	for _, field := range message.defaultMessage().Fields {
		if field.Tags&Header > 0 {
			header := byte(field.Data.(int32))
			if _, ok := allMessageHeaders[header]; ok {
				panic(fmt.Errorf("Header already taken.\nMessage:\n\n%s", message.defaultMessage().String()))
			}
			allMessageHeaders[header] = message
			return
		}
	}
	panic(fmt.Errorf("Header does not exist.\nMessage:\n\n%s", message.defaultMessage().String()))
}

// initializeDefaultMessage creates the internal structure of the default message, while ensuring that the structure of
// the message is correct. This should be called in an init() function.
func initializeDefaultMessage(messageType MessageType) {
	message := messageType.defaultMessage()
	if _, ok := allMessageDefaults[message]; ok {
		panic(fmt.Errorf("Message default was used in another message.\nMessage:\n\n%s", message.String()))
	}
	allMessageDefaults[message] = struct{}{}
	if message.info != nil {
		panic(fmt.Errorf("Message has already been initialized.\nMessage:\n\n%s", message.String()))
	}
	if _, ok := allMessageNames[message.Name]; ok {
		panic(fmt.Errorf("Message has already been initialized with the same name.\nName: %s", message.Name))
	}
	allMessageNames[message.Name] = struct{}{}
	message.info = &messageInfo{make(map[string]messageFieldInfo), false, message}
	message.isDefault = true

	allFieldNames := make(map[string]struct{}) // Verify that all field names are unique
	headerFound := false                       // Only one header may exist in the message
	messageLengthFound := false                // Only one message length may exist in the message
	endingByteNFound := false                  // If a ByteN has been found that does not have a preceding ByteCount
	repeatedFoundHeight := 0                   // The depth that a Repeated type has been found
	type FieldTraversal struct {
		Index  int
		Fields []*Field
	}

	ftStack := NewStack[FieldTraversal]()
	ftStack.Push(FieldTraversal{0, message.Fields})
	for !ftStack.Empty() {
		// If we're at the end of the loop for this stacked entry, then we pop it and move to the next
		if ftStack.Peek().Index >= len(ftStack.Peek().Fields) {
			_ = ftStack.Pop()
			continue
		}
		// Check if we've found a ByteN that is not preceded by a ByteCount-tagged field, as it should be the last
		// field, and we're now looking at a field after it.
		if endingByteNFound {
			panic(fmt.Errorf("ByteN found that was not preceded by a field with the ByteCount tag.\nMessage:\n\n%s", message.String()))
		}
		// If the stack is larger than Repeated's height, then we're probably in Repeated's children.
		// Otherwise, there are more non-child fields after the Repeated type.
		if ftStack.Len() <= repeatedFoundHeight {
			panic(fmt.Errorf("Repeated is not on the last field at its level\nMessage:\n\n%s", message.String()))
		}
		// Grab the field.
		field := ftStack.Peek().Fields[ftStack.Peek().Index]
		// Verify uniqueness and correctness of tags (if any)
		if field.Tags&Header > 0 {
			if headerFound {
				panic(fmt.Errorf("Multiple headers in message.\nMessage:\n\n%s", message.String()))
			}
			headerFound = true
		}
		if field.Tags&(MessageLengthInclusive|MessageLengthExclusive) > 0 {
			if messageLengthFound {
				panic(fmt.Errorf("Multiple message lengths in message.\nMessage:\n\n%s", message.String()))
			}
			switch field.Type {
			case Byte1, Int8, Int16, Int32:
			default:
				panic(fmt.Errorf("Message length tags are only allowed on integer types.\nField: %s\nMessage:\n\n%s", field.Name, message.String()))
			}
			messageLengthFound = true
		}
		if field.Tags&ByteCount > 0 {
			switch field.Type {
			case Byte1, Int8, Int16, Int32:
			default:
				panic(fmt.Errorf("ByteCount tag is only allowed on integer types.\nField: %s\nMessage:\n\n%s", field.Name, message.String()))
			}
		}
		if field.Tags&ExcludeTerminator > 0 && field.Type != String {
			panic(fmt.Errorf("ExcludeTerminator tag is only allowed on String fields.\nField: %s\nMessage:\n\n%s", field.Name, message.String()))
		}
		// Verify uniqueness of names (case-sensitive for maximum flexibility)
		if len(field.Name) == 0 {
			panic(fmt.Errorf("All fields must have a name.\nMessage:\n\n%s", message.String()))
		}
		if _, ok := allFieldNames[field.Name]; ok {
			panic(fmt.Errorf("Multiple fields with the same name.\nMessage:\n\n%s", message.String()))
		}
		allFieldNames[field.Name] = struct{}{}
		// Verify that ByteN is the last field, or is preceded by a field with the ByteCount tag
		usesByteCount := false
		if field.Type == ByteN {
			// If the preceding field has the ByteCount tag, then ByteN does not have the ending-field-only restriction
			if ftStack.Peek().Index > 0 && (ftStack.Peek().Fields[ftStack.Peek().Index-1].Tags&ByteCount > 0) {
				usesByteCount = true
			} else {
				endingByteNFound = true
			}
		}
		// Verify the type for each default value
		switch field.Type {
		case Byte1, Int8, Int16, Int32, Repeated:
			if _, ok := field.Data.(int32); !ok {
				panic(fmt.Errorf("Integer field types must set their Data to an int32 value.\nField: %s\nMessage:\n\n%s", field.Name, message.String()))
			}
		case ByteN:
			if _, ok := field.Data.([]byte); !ok {
				panic(fmt.Errorf("ByteN fields must set their Data to a []byte value.\nField: %s\nMessage:\n\n%s", field.Name, message.String()))
			}
		case String:
			if _, ok := field.Data.(string); !ok {
				panic(fmt.Errorf("String fields must set their Data to a string value.\nField: %s\nMessage:\n\n%s", field.Name, message.String()))
			}
		default:
			panic("message type has not been defined")
		}
		// Verify that, for fields with children, the default count matches the default child count
		if len(field.Children) > 0 {
			count := int32(0)
			switch field.Type {
			case Byte1, Int8, Int16, Int32, Repeated:
				count = field.Data.(int32)
			default:
				panic(fmt.Errorf("Only integer types may have children, as they determine the count.\nField: %s\nMessage:\n\n%s", field.Name, message.String()))
			}
			// A value of zero means that the child is only used as a prototype. A value of one means that the child is
			// actually used as a default value. We do not allow declaring children with multiple default values.
			if count != 0 && count != 1 {
				panic(fmt.Errorf("Only integer types may have children, as they determine the count.\nField: %s\nMessage:\n\n%s", field.Name, message.String()))
			}
			if len(field.Children) > 1 {
				panic(fmt.Errorf("Only a single child may be declared in the default message.\nField: %s\nMessage:\n\n%s", field.Name, message.String()))
			}
		}
		// Repeated may only be on a single field. Children of a Repeated field cannot also have Repeated children.
		if field.Type == Repeated {
			if repeatedFoundHeight != 0 {
				panic(fmt.Errorf("Multiple Repeated types declared.\nField: %s\nMessage:\n\n%s", field.Name, message.String()))
			}
			repeatedFoundHeight = ftStack.Len()
		}
		// RepeatedTerminator is only allowed on Repeated types, and therefore follows all of its restrictions automatically.
		if field.Tags&RepeatedTerminator > 0 {
			if field.Type != Repeated {
				panic(fmt.Errorf("RepeatedTerminator may only be used on a Repeated type.\nMessage:\n\n%s", message.String()))
			}
			message.info.appendNullByte = true
		}

		// Write the field info into our message
		parentName := ""
		if ftStack.Len() > 1 {
			parentName = ftStack.PeekDepth(1).Fields[ftStack.PeekDepth(1).Index-1].Name
		}
		message.info.fieldInfo[field.Name] = messageFieldInfo{
			RelativeIndex: ftStack.Peek().Index,
			Parent:        parentName,
			UsesByteCount: usesByteCount,
		}

		// Increment the index
		ftStack.PeekReference().Index++
		// If there are any children, then we throw them onto the stack
		if len(field.Children) == 1 {
			ftStack.Push(FieldTraversal{0, field.Children[0]})
		}
	}
}
