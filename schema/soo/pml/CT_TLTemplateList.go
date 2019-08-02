// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_TLTemplateList struct {
	// Template Effects
	Tmpl []*CT_TLTemplate
}

func NewCT_TLTemplateList() *CT_TLTemplateList {
	ret := &CT_TLTemplateList{}
	return ret
}

func (m *CT_TLTemplateList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Tmpl != nil {
		setmpl := xml.StartElement{Name: xml.Name{Local: "p:tmpl"}}
		for _, c := range m.Tmpl {
			e.EncodeElement(c, setmpl)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLTemplateList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TLTemplateList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "tmpl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "tmpl"}:
				tmp := NewCT_TLTemplate()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tmpl = append(m.Tmpl, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_TLTemplateList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTemplateList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLTemplateList and its children
func (m *CT_TLTemplateList) Validate() error {
	return m.ValidateWithPath("CT_TLTemplateList")
}

// ValidateWithPath validates the CT_TLTemplateList and its children, prefixing error messages with path
func (m *CT_TLTemplateList) ValidateWithPath(path string) error {
	for i, v := range m.Tmpl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tmpl[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
