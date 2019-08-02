// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_RotX struct {
	ValAttr *int8
}

func NewCT_RotX() *CT_RotX {
	ret := &CT_RotX{}
	return ret
}

func (m *CT_RotX) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RotX) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := int8(parsed)
			m.ValAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RotX: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_RotX and its children
func (m *CT_RotX) Validate() error {
	return m.ValidateWithPath("CT_RotX")
}

// ValidateWithPath validates the CT_RotX and its children, prefixing error messages with path
func (m *CT_RotX) ValidateWithPath(path string) error {
	if m.ValAttr != nil {
		if *m.ValAttr < -90 {
			return fmt.Errorf("%s/m.ValAttr must be >= -90 (have %v)", path, *m.ValAttr)
		}
		if *m.ValAttr > 90 {
			return fmt.Errorf("%s/m.ValAttr must be <= 90 (have %v)", path, *m.ValAttr)
		}
	}
	return nil
}
