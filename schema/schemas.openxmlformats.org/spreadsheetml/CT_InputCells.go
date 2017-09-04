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

type CT_InputCells struct {
	// Reference
	RAttr string
	// Deleted
	DeletedAttr *bool
	// Undone
	UndoneAttr *bool
	// Value
	ValAttr string
	// Number Format Id
	NumFmtIdAttr *uint32
}

func NewCT_InputCells() *CT_InputCells {
	ret := &CT_InputCells{}
	ret.DeletedAttr = gooxml.Bool(false)
	ret.UndoneAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_InputCells) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	if m.DeletedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "deleted"},
			Value: fmt.Sprintf("%d", b2i(*m.DeletedAttr))})
	}
	if m.UndoneAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "undone"},
			Value: fmt.Sprintf("%d", b2i(*m.UndoneAttr))})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	if m.NumFmtIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "numFmtId"},
			Value: fmt.Sprintf("%v", *m.NumFmtIdAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_InputCells) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.DeletedAttr = gooxml.Bool(false)
	m.UndoneAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RAttr = parsed
		}
		if attr.Name.Local == "deleted" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DeletedAttr = &parsed
		}
		if attr.Name.Local == "undone" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UndoneAttr = &parsed
		}
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
		}
		if attr.Name.Local == "numFmtId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.NumFmtIdAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_InputCells: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_InputCells and its children
func (m *CT_InputCells) Validate() error {
	return m.ValidateWithPath("CT_InputCells")
}

// ValidateWithPath validates the CT_InputCells and its children, prefixing error messages with path
func (m *CT_InputCells) ValidateWithPath(path string) error {
	return nil
}
