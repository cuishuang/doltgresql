// Copyright 2024 Dolthub, Inc.
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

package functions

import (
	"math"

	"github.com/dolthub/doltgresql/server/functions/framework"

	pgtypes "github.com/dolthub/doltgresql/server/types"
)

// init registers the functions to the catalog.
func init() {
	framework.RegisterFunction(asind_float64)
}

// asind_float64 represents the PostgreSQL function of the same name, taking the same parameters.
var asind_float64 = framework.Function1{
	Name:       "asind",
	Return:     pgtypes.Float64,
	Parameters: []pgtypes.DoltgresType{pgtypes.Float64},
	Callable: func(ctx framework.Context, val1 any) (any, error) {
		if val1 == nil {
			return nil, nil
		}
		return toDegrees(math.Asin(val1.(float64))), nil
	},
}
