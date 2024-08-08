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
	"github.com/dolthub/go-mysql-server/sql"

	"github.com/dolthub/doltgresql/server/functions/framework"
	pgtypes "github.com/dolthub/doltgresql/server/types"
	"github.com/dolthub/doltgresql/server/types/oid"
)

// initPgGetFunctionIdentityArguments registers the functions to the catalog.
func initPgGetFunctionIdentityArguments() {
	framework.RegisterFunction(pg_get_function_identity_arguments_oid)
}

// pg_get_function_identity_arguments_oid represents the PostgreSQL system information function.
var pg_get_function_identity_arguments_oid = framework.Function1{
	Name:               "pg_get_function_identity_arguments",
	Return:             pgtypes.Text,
	Parameters:         [1]pgtypes.DoltgresType{pgtypes.Oid},
	IsNonDeterministic: true,
	Strict:             true,
	Callable: func(ctx *sql.Context, _ [2]pgtypes.DoltgresType, val any) (any, error) {
		oidVal := val.(uint32)
		err := oid.RunCallback(ctx, oidVal, oid.Callbacks{
			Function: func(ctx *sql.Context, function oid.ItemFunction) (cont bool, err error) {
				// TODO: sql.Function does not have sufficient information about its arguments
				return false, nil
			},
		})
		if err != nil {
			return "", err
		}
		return "", nil
	},
}
