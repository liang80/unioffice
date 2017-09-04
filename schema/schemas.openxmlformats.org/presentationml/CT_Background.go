// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Background struct {
	// Black and White Mode
	BwModeAttr drawingml.ST_BlackWhiteMode
	// Background Properties
	BgPr *CT_BackgroundProperties
	// Background Style Reference
	BgRef *drawingml.CT_StyleMatrixReference
}

func NewCT_Background() *CT_Background {
	ret := &CT_Background{}
	ret.BwModeAttr = drawingml.ST_BlackWhiteModeWhite
	return ret
}

func (m *CT_Background) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BwModeAttr != drawingml.ST_BlackWhiteModeUnset {
		attr, err := m.BwModeAttr.MarshalXMLAttr(xml.Name{Local: "bwMode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.BgPr != nil {
		sebgPr := xml.StartElement{Name: xml.Name{Local: "p:bgPr"}}
		e.EncodeElement(m.BgPr, sebgPr)
	}
	if m.BgRef != nil {
		sebgRef := xml.StartElement{Name: xml.Name{Local: "p:bgRef"}}
		e.EncodeElement(m.BgRef, sebgRef)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Background) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.BwModeAttr = drawingml.ST_BlackWhiteModeWhite
	for _, attr := range start.Attr {
		if attr.Name.Local == "bwMode" {
			m.BwModeAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_Background:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "bgPr":
				m.BgPr = NewCT_BackgroundProperties()
				if err := d.DecodeElement(m.BgPr, &el); err != nil {
					return err
				}
			case "bgRef":
				m.BgRef = drawingml.NewCT_StyleMatrixReference()
				if err := d.DecodeElement(m.BgRef, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Background %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Background
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Background and its children
func (m *CT_Background) Validate() error {
	return m.ValidateWithPath("CT_Background")
}

// ValidateWithPath validates the CT_Background and its children, prefixing error messages with path
func (m *CT_Background) ValidateWithPath(path string) error {
	if err := m.BwModeAttr.ValidateWithPath(path + "/BwModeAttr"); err != nil {
		return err
	}
	if m.BgPr != nil {
		if err := m.BgPr.ValidateWithPath(path + "/BgPr"); err != nil {
			return err
		}
	}
	if m.BgRef != nil {
		if err := m.BgRef.ValidateWithPath(path + "/BgRef"); err != nil {
			return err
		}
	}
	return nil
}
