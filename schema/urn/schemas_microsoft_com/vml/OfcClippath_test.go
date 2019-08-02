// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcClippathConstructor(t *testing.T) {
	v := vml.NewOfcClippath()
	if v == nil {
		t.Errorf("vml.NewOfcClippath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcClippath should validate: %s", err)
	}
}

func TestOfcClippathMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcClippath()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcClippath()
	xml.Unmarshal(buf, v2)
}
