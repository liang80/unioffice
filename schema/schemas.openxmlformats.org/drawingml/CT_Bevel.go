// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Bevel struct {
	WAttr    *int64
	HAttr    *int64
	PrstAttr ST_BevelPresetType
}

func NewCT_Bevel() *CT_Bevel {
	ret := &CT_Bevel{}
	ret.PrstAttr = ST_BevelPresetTypeCircle
	return ret
}

func (m *CT_Bevel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.WAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w"},
			Value: fmt.Sprintf("%v", *m.WAttr)})
	}
	if m.HAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "h"},
			Value: fmt.Sprintf("%v", *m.HAttr)})
	}
	if m.PrstAttr != ST_BevelPresetTypeUnset {
		attr, err := m.PrstAttr.MarshalXMLAttr(xml.Name{Local: "prst"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Bevel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PrstAttr = ST_BevelPresetTypeCircle
	for _, attr := range start.Attr {
		if attr.Name.Local == "w" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.WAttr = &parsed
		}
		if attr.Name.Local == "h" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.HAttr = &parsed
		}
		if attr.Name.Local == "prst" {
			m.PrstAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Bevel: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Bevel and its children
func (m *CT_Bevel) Validate() error {
	return m.ValidateWithPath("CT_Bevel")
}

// ValidateWithPath validates the CT_Bevel and its children, prefixing error messages with path
func (m *CT_Bevel) ValidateWithPath(path string) error {
	if m.WAttr != nil {
		if *m.WAttr < 0 {
			return fmt.Errorf("%s/m.WAttr must be >= 0 (have %v)", path, *m.WAttr)
		}
		if *m.WAttr > 27273042316900 {
			return fmt.Errorf("%s/m.WAttr must be <= 27273042316900 (have %v)", path, *m.WAttr)
		}
	}
	if m.HAttr != nil {
		if *m.HAttr < 0 {
			return fmt.Errorf("%s/m.HAttr must be >= 0 (have %v)", path, *m.HAttr)
		}
		if *m.HAttr > 27273042316900 {
			return fmt.Errorf("%s/m.HAttr must be <= 27273042316900 (have %v)", path, *m.HAttr)
		}
	}
	if err := m.PrstAttr.ValidateWithPath(path + "/PrstAttr"); err != nil {
		return err
	}
	return nil
}
