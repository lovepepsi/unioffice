// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_TextSpacingPercentOrPercentString is a union type
type ST_TextSpacingPercentOrPercentString struct {
	ST_TextSpacingPercent *int32
	ST_Percentage         *string
}

func (m *ST_TextSpacingPercentOrPercentString) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextSpacingPercentOrPercentString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_TextSpacingPercent != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_TextSpacingPercent)))
	}
	if m.ST_Percentage != nil {
		e.EncodeToken(xml.CharData(*m.ST_Percentage))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_TextSpacingPercentOrPercentString) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_TextSpacingPercent != nil {
		mems = append(mems, "ST_TextSpacingPercent")
	}
	if m.ST_Percentage != nil {
		mems = append(mems, "ST_Percentage")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_TextSpacingPercentOrPercentString) String() string {
	if m.ST_TextSpacingPercent != nil {
		return fmt.Sprintf("%v", *m.ST_TextSpacingPercent)
	}
	if m.ST_Percentage != nil {
		return fmt.Sprintf("%v", *m.ST_Percentage)
	}
	return ""
}
