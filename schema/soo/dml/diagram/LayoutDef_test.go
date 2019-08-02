// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/diagram"
)

func TestLayoutDefConstructor(t *testing.T) {
	v := diagram.NewLayoutDef()
	if v == nil {
		t.Errorf("diagram.NewLayoutDef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.LayoutDef should validate: %s", err)
	}
}

func TestLayoutDefMarshalUnmarshal(t *testing.T) {
	v := diagram.NewLayoutDef()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewLayoutDef()
	xml.Unmarshal(buf, v2)
}
