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

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/dml"
)

type CT_PivotFmt struct {
	Idx    *CT_UnsignedInt
	SpPr   *dml.CT_ShapeProperties
	TxPr   *dml.CT_TextBody
	Marker *CT_Marker
	DLbl   *CT_DLbl
	ExtLst *CT_ExtensionList
}

func NewCT_PivotFmt() *CT_PivotFmt {
	ret := &CT_PivotFmt{}
	ret.Idx = NewCT_UnsignedInt()
	return ret
}

func (m *CT_PivotFmt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seidx := xml.StartElement{Name: xml.Name{Local: "c:idx"}}
	e.EncodeElement(m.Idx, seidx)
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "c:spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "c:txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	if m.Marker != nil {
		semarker := xml.StartElement{Name: xml.Name{Local: "c:marker"}}
		e.EncodeElement(m.Marker, semarker)
	}
	if m.DLbl != nil {
		sedLbl := xml.StartElement{Name: xml.Name{Local: "c:dLbl"}}
		e.EncodeElement(m.DLbl, sedLbl)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotFmt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Idx = NewCT_UnsignedInt()
lCT_PivotFmt:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "idx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "idx"}:
				if err := d.DecodeElement(m.Idx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "spPr"}:
				m.SpPr = dml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "txPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "txPr"}:
				m.TxPr = dml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "marker"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "marker"}:
				m.Marker = NewCT_Marker()
				if err := d.DecodeElement(m.Marker, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dLbl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dLbl"}:
				m.DLbl = NewCT_DLbl()
				if err := d.DecodeElement(m.DLbl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_PivotFmt %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotFmt
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotFmt and its children
func (m *CT_PivotFmt) Validate() error {
	return m.ValidateWithPath("CT_PivotFmt")
}

// ValidateWithPath validates the CT_PivotFmt and its children, prefixing error messages with path
func (m *CT_PivotFmt) ValidateWithPath(path string) error {
	if err := m.Idx.ValidateWithPath(path + "/Idx"); err != nil {
		return err
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	if m.Marker != nil {
		if err := m.Marker.ValidateWithPath(path + "/Marker"); err != nil {
			return err
		}
	}
	if m.DLbl != nil {
		if err := m.DLbl.ValidateWithPath(path + "/DLbl"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
