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
	"fmt"
	"math"

	"github.com/dolthub/go-mysql-server/sql"

	"github.com/dolthub/doltgresql/server/functions/framework"
	pgtypes "github.com/dolthub/doltgresql/server/types"
)

// initAcos registers the functions to the catalog.
func initAcos() {
	framework.RegisterFunction(acos_float64)
}

// acos_float64 represents the PostgreSQL function of the same name, taking the same parameters.
var acos_float64 = framework.Function1{
	Name:       "acos",
	Return:     pgtypes.Float64,
	Parameters: []pgtypes.DoltgresType{pgtypes.Float64},
	Callable: func(ctx *sql.Context, val1 any) (any, error) {
		r := math.Acos(val1.(float64))
		if math.IsNaN(r) {
			return nil, fmt.Errorf("input is out of range")
		}
		return r, nil
	},
	Strict: true,
}
