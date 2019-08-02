// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package math

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_D struct {
	DPr *CT_DPr
	E   []*CT_OMathArg
}

func NewCT_D() *CT_D {
	ret := &CT_D{}
	return ret
}

func (m *CT_D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.DPr != nil {
		sedPr := xml.StartElement{Name: xml.Name{Local: "m:dPr"}}
		e.EncodeElement(m.DPr, sedPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	for _, c := range m.E {
		e.EncodeElement(c, see)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "dPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "dPr"}:
				m.DPr = NewCT_DPr()
				if err := d.DecodeElement(m.DPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "e"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "e"}:
				tmp := NewCT_OMathArg()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.E = append(m.E, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_D %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_D
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_D and its children
func (m *CT_D) Validate() error {
	return m.ValidateWithPath("CT_D")
}

// ValidateWithPath validates the CT_D and its children, prefixing error messages with path
func (m *CT_D) ValidateWithPath(path string) error {
	if m.DPr != nil {
		if err := m.DPr.ValidateWithPath(path + "/DPr"); err != nil {
			return err
		}
	}
	for i, v := range m.E {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/E[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
