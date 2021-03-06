// Copyright 2021 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package validator // import "miniflux.app/validator"

import "testing"

func TestIsValidURL(t *testing.T) {
	scenarios := map[string]bool{
		"https://www.example.org": true,
		"http://www.example.org/": true,
		"www.example.org":         false,
	}

	for link, expected := range scenarios {
		result := isValidURL(link)
		if result != expected {
			t.Errorf(`Unexpected result, got %v instead of %v`, result, expected)
		}
	}
}

func TestValidateRange(t *testing.T) {
	if err := ValidateRange(-1, 0); err == nil {
		t.Error(`An invalid offset should generate a error`)
	}

	if err := ValidateRange(0, -1); err == nil {
		t.Error(`An invalid limit should generate a error`)
	}

	if err := ValidateRange(42, 42); err != nil {
		t.Error(`A valid offset and limit should not generate any error`)
	}
}

func TestValidateDirection(t *testing.T) {
	for _, status := range []string{"asc", "desc"} {
		if err := ValidateDirection(status); err != nil {
			t.Error(`A valid direction should not generate any error`)
		}
	}

	if err := ValidateDirection("invalid"); err == nil {
		t.Error(`An invalid direction should generate a error`)
	}
}
