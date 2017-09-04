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

type CT_Selection struct {
	// Pane
	PaneAttr ST_Pane
	// Active Cell Location
	ActiveCellAttr *string
	// Active Cell Index
	ActiveCellIdAttr *uint32
	// Sequence of References
	SqrefAttr *ST_Sqref
}

func NewCT_Selection() *CT_Selection {
	ret := &CT_Selection{}
	ret.PaneAttr = ST_PaneTopLeft
	ret.ActiveCellIdAttr = gooxml.Uint32(0)
	return ret
}

func (m *CT_Selection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PaneAttr != ST_PaneUnset {
		attr, err := m.PaneAttr.MarshalXMLAttr(xml.Name{Local: "pane"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ActiveCellAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "activeCell"},
			Value: fmt.Sprintf("%v", *m.ActiveCellAttr)})
	}
	if m.ActiveCellIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "activeCellId"},
			Value: fmt.Sprintf("%v", *m.ActiveCellIdAttr)})
	}
	if m.SqrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sqref"},
			Value: fmt.Sprintf("%v", *m.SqrefAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Selection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PaneAttr = ST_PaneTopLeft
	m.ActiveCellIdAttr = gooxml.Uint32(0)
	for _, attr := range start.Attr {
		if attr.Name.Local == "pane" {
			m.PaneAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "activeCell" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ActiveCellAttr = &parsed
		}
		if attr.Name.Local == "activeCellId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ActiveCellIdAttr = &pt
		}
		if attr.Name.Local == "sqref" {
			parsed, err := ParseSliceST_Sqref(attr.Value)
			if err != nil {
				return err
			}
			m.SqrefAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Selection: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Selection and its children
func (m *CT_Selection) Validate() error {
	return m.ValidateWithPath("CT_Selection")
}

// ValidateWithPath validates the CT_Selection and its children, prefixing error messages with path
func (m *CT_Selection) ValidateWithPath(path string) error {
	if err := m.PaneAttr.ValidateWithPath(path + "/PaneAttr"); err != nil {
		return err
	}
	return nil
}
