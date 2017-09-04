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

type CT_RevisionConflict struct {
	// Sheet Id
	SheetIdAttr *uint32
	RIdAttr     *uint32
	UaAttr      *bool
	RaAttr      *bool
}

func NewCT_RevisionConflict() *CT_RevisionConflict {
	ret := &CT_RevisionConflict{}
	ret.UaAttr = gooxml.Bool(false)
	ret.RaAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_RevisionConflict) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SheetIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheetId"},
			Value: fmt.Sprintf("%v", *m.SheetIdAttr)})
	}
	if m.RIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rId"},
			Value: fmt.Sprintf("%v", *m.RIdAttr)})
	}
	if m.UaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ua"},
			Value: fmt.Sprintf("%d", b2i(*m.UaAttr))})
	}
	if m.RaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ra"},
			Value: fmt.Sprintf("%d", b2i(*m.RaAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RevisionConflict) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.UaAttr = gooxml.Bool(false)
	m.RaAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SheetIdAttr = &pt
		}
		if attr.Name.Local == "rId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RIdAttr = &pt
		}
		if attr.Name.Local == "ua" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UaAttr = &parsed
		}
		if attr.Name.Local == "ra" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RaAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RevisionConflict: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_RevisionConflict and its children
func (m *CT_RevisionConflict) Validate() error {
	return m.ValidateWithPath("CT_RevisionConflict")
}

// ValidateWithPath validates the CT_RevisionConflict and its children, prefixing error messages with path
func (m *CT_RevisionConflict) ValidateWithPath(path string) error {
	return nil
}
