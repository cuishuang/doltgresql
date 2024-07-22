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

	"github.com/dolthub/go-mysql-server/sql"

	"github.com/dolthub/doltgresql/server/functions/framework"
	pgtypes "github.com/dolthub/doltgresql/server/types"
)

// initPi registers the functions to the catalog.
func initPi() {
	framework.RegisterFunction(pi)
}

// pi represents the PostgreSQL function of the same name, taking the same parameters.
var pi = framework.Function0{
	Name:   "pi",
	Return: pgtypes.Float64,
	Strict: true,
	Callable: func(ctx *sql.Context) (any, error) {
		return float64(math.Pi), nil
	},
}
