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

	"baliance.com/gooxml"
)

type CT_TextField struct {
	// Field Type
	TypeAttr ST_ExternalConnectionType
	// Position
	PositionAttr *uint32
}

func NewCT_TextField() *CT_TextField {
	ret := &CT_TextField{}
	ret.TypeAttr = ST_ExternalConnectionTypeGeneral
	ret.PositionAttr = gooxml.Uint32(0)
	return ret
}

func (m *CT_TextField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_ExternalConnectionTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.PositionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "position"},
			Value: fmt.Sprintf("%v", *m.PositionAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextField) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_ExternalConnectionTypeGeneral
	m.PositionAttr = gooxml.Uint32(0)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "position" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.PositionAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextField: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextField and its children
func (m *CT_TextField) Validate() error {
	return m.ValidateWithPath("CT_TextField")
}

// ValidateWithPath validates the CT_TextField and its children, prefixing error messages with path
func (m *CT_TextField) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
