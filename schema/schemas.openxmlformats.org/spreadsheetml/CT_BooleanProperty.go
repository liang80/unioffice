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

type CT_BooleanProperty struct {
	// Value
	ValAttr *bool
}

func NewCT_BooleanProperty() *CT_BooleanProperty {
	ret := &CT_BooleanProperty{}
	ret.ValAttr = gooxml.Bool(true)
	return ret
}

func (m *CT_BooleanProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%d", b2i(*m.ValAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BooleanProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = gooxml.Bool(true)
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_BooleanProperty: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_BooleanProperty and its children
func (m *CT_BooleanProperty) Validate() error {
	return m.ValidateWithPath("CT_BooleanProperty")
}

// ValidateWithPath validates the CT_BooleanProperty and its children, prefixing error messages with path
func (m *CT_BooleanProperty) ValidateWithPath(path string) error {
	return nil
}
