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

package output

import "testing"

func TestCreateTextSearchDictionary(t *testing.T) {
	tests := []QueryParses{
		Unimplemented("CREATE TEXT SEARCH DICTIONARY name ( TEMPLATE = template )"),
		Unimplemented("CREATE TEXT SEARCH DICTIONARY name ( TEMPLATE = template , option = value )"),
		Unimplemented("CREATE TEXT SEARCH DICTIONARY name ( TEMPLATE = template , option = value , option = value )"),
	}
	RunTests(t, tests)
}
