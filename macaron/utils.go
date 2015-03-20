// Copyright 2015 Macaron Authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"errors"
	"strings"

	"github.com/Unknwon/com"
)

// pathInsideGOPATH checks if given path is in the one of GOPATHs.
func pathInsideGOPATH(location string) error {
	insideGoPath := false
	gopaths := com.GetGOPATHs()
	for _, wg := range gopaths {
		if strings.HasPrefix(strings.ToLower(location), strings.ToLower(wg)) {
			insideGoPath = true
			break
		}
	}

	if !insideGoPath {
		return errors.New("path is not inside any of GOPATHs")
	}

	return nil
}
