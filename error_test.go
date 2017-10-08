/*
Copyright 2017 John Laswell

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package conerror

import (
	"errors"
	"testing"
)

var _ ConError = New("")
var _ error = New("")
var _ ConError = NewFromError(errors.New(""))
var _ error = NewFromError(errors.New(""))

func TestErrorBasics(t *testing.T) {
	cerr := New("Something bad")
	if cerr.Error() != "Something bad" {
		t.Fatal("Error() method is not returning expected result")
	}

	dir := cerr.Get("direction")
	if dir != nil {
		t.Fatalf("Get() provided unexpected nil value")
	}

	cerr.Set("direction", "West")
	dir = cerr.Get("direction")
	if dir != "West" {
		t.Fatalf("Get() provided unexpected value for key \"direction\"")
	}
}
