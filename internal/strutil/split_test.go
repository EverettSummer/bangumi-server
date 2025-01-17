// Copyright (c) 2022 Trim21 <trim21.me@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package strutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bangumi/server/internal/strutil"
)

func TestPartition(t *testing.T) {
	t.Parallel()

	key, value := strutil.Partition("a=b", '=')
	assert.Equal(t, "a", key)
	assert.Equal(t, "b", value)

	key, value = strutil.Partition("a=", '=')
	assert.Equal(t, "a", key)
	assert.Equal(t, "", value)

	key, value = strutil.Partition("=", '=')
	assert.Equal(t, "", key)
	assert.Equal(t, "", value)

	key, value = strutil.Partition("ab", '=')
	assert.Equal(t, "ab", key)
	assert.Equal(t, "", value)
}

func TestSplit(t *testing.T) {
	t.Parallel()

	s := strutil.Split("a=b", "=")
	assert.EqualValues(t, []string{"a", "b"}, s)

	s = strutil.Split("a==b", "=")
	assert.EqualValues(t, []string{"a", "b"}, s)

	s = strutil.Split("", "=")
	assert.EqualValues(t, []string{}, s)
}
