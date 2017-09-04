// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type CT_PatternFill struct {
	// Pattern Type
	PatternTypeAttr ST_PatternType
	// Foreground Color
	FgColor *CT_Color
	// Background Color
	BgColor *CT_Color
}

func NewCT_PatternFill() *CT_PatternFill {
	ret := &CT_PatternFill{}
	ret.PatternTypeAttr = ST_PatternType(1)
	return ret
}

func (m *CT_PatternFill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PatternTypeAttr != ST_PatternTypeUnset {
		attr, err := m.PatternTypeAttr.MarshalXMLAttr(xml.Name{Local: "patternType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.FgColor != nil {
		sefgColor := xml.StartElement{Name: xml.Name{Local: "x:fgColor"}}
		e.EncodeElement(m.FgColor, sefgColor)
	}
	if m.BgColor != nil {
		sebgColor := xml.StartElement{Name: xml.Name{Local: "x:bgColor"}}
		e.EncodeElement(m.BgColor, sebgColor)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PatternFill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PatternTypeAttr = ST_PatternType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "patternType" {
			m.PatternTypeAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_PatternFill:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "fgColor":
				m.FgColor = NewCT_Color()
				if err := d.DecodeElement(m.FgColor, &el); err != nil {
					return err
				}
			case "bgColor":
				m.BgColor = NewCT_Color()
				if err := d.DecodeElement(m.BgColor, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_PatternFill %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PatternFill
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PatternFill and its children
func (m *CT_PatternFill) Validate() error {
	return m.ValidateWithPath("CT_PatternFill")
}

// ValidateWithPath validates the CT_PatternFill and its children, prefixing error messages with path
func (m *CT_PatternFill) ValidateWithPath(path string) error {
	if err := m.PatternTypeAttr.ValidateWithPath(path + "/PatternTypeAttr"); err != nil {
		return err
	}
	if m.FgColor != nil {
		if err := m.FgColor.ValidateWithPath(path + "/FgColor"); err != nil {
			return err
		}
	}
	if m.BgColor != nil {
		if err := m.BgColor.ValidateWithPath(path + "/BgColor"); err != nil {
			return err
		}
	}
	return nil
}
