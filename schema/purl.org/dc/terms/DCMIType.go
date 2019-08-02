// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package terms

import (
	"encoding/xml"
	"fmt"
)

type DCMIType struct {
}

func NewDCMIType() *DCMIType {
	ret := &DCMIType{}
	return ret
}

func (m *DCMIType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "DCMIType"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *DCMIType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing DCMIType: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the DCMIType and its children
func (m *DCMIType) Validate() error {
	return m.ValidateWithPath("DCMIType")
}

// ValidateWithPath validates the DCMIType and its children, prefixing error messages with path
func (m *DCMIType) ValidateWithPath(path string) error {
	return nil
}
