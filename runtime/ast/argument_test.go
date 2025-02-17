/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArgument_MarshalJSON(t *testing.T) {

	t.Parallel()

	t.Run("without label", func(t *testing.T) {

		t.Parallel()

		argument := Argument{
			Expression: &BoolExpression{
				Value: false,
				Range: Range{
					StartPos: Position{Offset: 1, Line: 2, Column: 3},
					EndPos:   Position{Offset: 4, Line: 5, Column: 6},
				},
			},
		}

		actual, err := json.Marshal(argument)
		require.NoError(t, err)

		assert.JSONEq(t,
			`
            {
                "Expression": {
                    "Type": "BoolExpression",
                    "Value": false,
                    "StartPos": {"Offset": 1, "Line": 2, "Column": 3},
                    "EndPos": {"Offset": 4, "Line": 5, "Column": 6}
                },
                "StartPos": {"Offset": 1, "Line": 2, "Column": 3},
                "EndPos": {"Offset": 4, "Line": 5, "Column": 6}
            }
            `,
			string(actual),
		)
	})

	t.Run("with label", func(t *testing.T) {

		t.Parallel()

		argument := Argument{
			Label:         "ok",
			LabelStartPos: &Position{Offset: 7, Line: 8, Column: 9},
			LabelEndPos:   &Position{Offset: 10, Line: 11, Column: 12},
			Expression: &BoolExpression{
				Value: false,
				Range: Range{
					StartPos: Position{Offset: 1, Line: 2, Column: 3},
					EndPos:   Position{Offset: 4, Line: 5, Column: 6},
				},
			},
		}

		actual, err := json.Marshal(argument)
		require.NoError(t, err)

		assert.JSONEq(t,
			`
            {
                "Label": "ok",
                "LabelStartPos": {"Offset": 7, "Line": 8, "Column": 9},
                "LabelEndPos": {"Offset": 10, "Line": 11, "Column": 12},
                "Expression": {
                    "Type": "BoolExpression",
                    "Value": false,
                    "StartPos": {"Offset": 1, "Line": 2, "Column": 3},
                    "EndPos": {"Offset": 4, "Line": 5, "Column": 6}
                },
                "StartPos": {"Offset": 7, "Line": 8, "Column": 9},
                "EndPos": {"Offset": 4, "Line": 5, "Column": 6}
            }
            `,
			string(actual),
		)
	})
}
