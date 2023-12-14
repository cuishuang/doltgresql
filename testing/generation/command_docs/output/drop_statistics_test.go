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

func TestDropStatistics(t *testing.T) {
	tests := []QueryParses{
		Unimplemented("DROP STATISTICS name"),
		Unimplemented("DROP STATISTICS IF EXISTS name"),
		Unimplemented("DROP STATISTICS name , name"),
		Unimplemented("DROP STATISTICS IF EXISTS name , name"),
		Unimplemented("DROP STATISTICS name CASCADE"),
		Unimplemented("DROP STATISTICS IF EXISTS name CASCADE"),
		Unimplemented("DROP STATISTICS name , name CASCADE"),
		Unimplemented("DROP STATISTICS IF EXISTS name , name CASCADE"),
		Unimplemented("DROP STATISTICS name RESTRICT"),
		Unimplemented("DROP STATISTICS IF EXISTS name RESTRICT"),
		Unimplemented("DROP STATISTICS name , name RESTRICT"),
		Unimplemented("DROP STATISTICS IF EXISTS name , name RESTRICT"),
	}
	RunTests(t, tests)
}
