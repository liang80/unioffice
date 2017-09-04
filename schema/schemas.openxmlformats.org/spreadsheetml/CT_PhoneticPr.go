// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_PhoneticPr struct {
	// Font Id
	FontIdAttr uint32
	// Character Type
	TypeAttr ST_PhoneticType
	// Alignment
	AlignmentAttr ST_PhoneticAlignment
}

func NewCT_PhoneticPr() *CT_PhoneticPr {
	ret := &CT_PhoneticPr{}
	ret.TypeAttr = ST_PhoneticTypeFullwidthKatakana
	ret.AlignmentAttr = ST_PhoneticAlignmentLeft
	return ret
}

func (m *CT_PhoneticPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fontId"},
		Value: fmt.Sprintf("%v", m.FontIdAttr)})
	if m.TypeAttr != ST_PhoneticTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AlignmentAttr != ST_PhoneticAlignmentUnset {
		attr, err := m.AlignmentAttr.MarshalXMLAttr(xml.Name{Local: "alignment"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PhoneticPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_PhoneticTypeFullwidthKatakana
	m.AlignmentAttr = ST_PhoneticAlignmentLeft
	for _, attr := range start.Attr {
		if attr.Name.Local == "fontId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.FontIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "alignment" {
			m.AlignmentAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PhoneticPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PhoneticPr and its children
func (m *CT_PhoneticPr) Validate() error {
	return m.ValidateWithPath("CT_PhoneticPr")
}

// ValidateWithPath validates the CT_PhoneticPr and its children, prefixing error messages with path
func (m *CT_PhoneticPr) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.AlignmentAttr.ValidateWithPath(path + "/AlignmentAttr"); err != nil {
		return err
	}
	return nil
}
