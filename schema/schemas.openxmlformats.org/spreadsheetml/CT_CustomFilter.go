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
)

type CT_CustomFilter struct {
	// Filter Comparison Operator
	OperatorAttr ST_FilterOperator
	// Top or Bottom Value
	ValAttr *string
}

func NewCT_CustomFilter() *CT_CustomFilter {
	ret := &CT_CustomFilter{}
	ret.OperatorAttr = ST_FilterOperatorEqual
	return ret
}

func (m *CT_CustomFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.OperatorAttr != ST_FilterOperatorUnset {
		attr, err := m.OperatorAttr.MarshalXMLAttr(xml.Name{Local: "operator"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OperatorAttr = ST_FilterOperatorEqual
	for _, attr := range start.Attr {
		if attr.Name.Local == "operator" {
			m.OperatorAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
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
			return fmt.Errorf("parsing CT_CustomFilter: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CustomFilter and its children
func (m *CT_CustomFilter) Validate() error {
	return m.ValidateWithPath("CT_CustomFilter")
}

// ValidateWithPath validates the CT_CustomFilter and its children, prefixing error messages with path
func (m *CT_CustomFilter) ValidateWithPath(path string) error {
	if err := m.OperatorAttr.ValidateWithPath(path + "/OperatorAttr"); err != nil {
		return err
	}
	return nil
}
