/**
* goDHT a "Distributed Hash Table" library for the go language
*
* This file is part of goDHT.
*
* goDHT is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* goDHT is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with goDHT.  If not, see <http://www.gnu.org/licenses/>.
 */
package node

import (
	"testing"
)

func Test_IsGood_Responded(t *testing.T) {

	undertest := New(GenerateID())

	undertest.SetLastResponse()

	if !undertest.IsGood() {
		t.Fatal("Node should be good but isn't", undertest)
	}
}

func Test_IsGood_LastQueryAndOnceResponded(t *testing.T) {

	undertest := New(GenerateID())

	undertest.SetLastQuery()
	undertest.SetOnceResponded()

	if !undertest.IsGood() {
		t.Fatal("Node should be good but isn't", undertest)
	}

}

func Test_IsGood_LastQuery_FAIL(t *testing.T) {

	undertest := New(GenerateID())

	undertest.SetLastQuery()

	if undertest.IsGood() {
		t.Fatal("Node should not be good", undertest)
	}
}

func Test_IsGood_OnceResponded_FAIL(t *testing.T) {

	undertest := New(GenerateID())

	undertest.SetOnceResponded()

	if undertest.IsGood() {
		t.Fatal("Node should not be good", undertest)
	}
}

func Test_IsGood_Fail(t *testing.T) {

	undertest := New(GenerateID())

	if undertest.IsGood() {
		t.Fatal("Node should not be good", undertest)
	}
}
