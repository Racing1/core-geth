// Copyright 2019 The multi-geth Authors
// This file is part of the multi-geth library.
//
// The multi-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The multi-geth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the multi-geth library. If not, see <http://www.gnu.org/licenses/>.

package tests

// writeDifficultyTestsReferencePairs defines pairs for test generation.
// TODO: move to more accessible api?
var writeDifficultyTestsReferencePairs = map[string]string{
	"Byzantium":      "ETC_Atlantis",
	"Constantinople": "ETC_Agharta",
	"EIP2384":        "ETC_Phoenix",
}
