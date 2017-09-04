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

type CT_CellFormula struct {
	TAttr    ST_CellFormulaType
	AcaAttr  *bool
	RefAttr  *string
	Dt2DAttr *bool
	DtrAttr  *bool
	Del1Attr *bool
	Del2Attr *bool
	R1Attr   *string
	R2Attr   *string
	CaAttr   *bool
	SiAttr   *uint32
	BxAttr   *bool
	Content  string
}

func NewCT_CellFormula() *CT_CellFormula {
	ret := &CT_CellFormula{}
	ret.TAttr = ST_CellFormulaType(1)
	return ret
}

func (m *CT_CellFormula) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TAttr != ST_CellFormulaTypeUnset {
		attr, err := m.TAttr.MarshalXMLAttr(xml.Name{Local: "t"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AcaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "aca"},
			Value: fmt.Sprintf("%d", b2i(*m.AcaAttr))})
	}
	if m.RefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
			Value: fmt.Sprintf("%v", *m.RefAttr)})
	}
	if m.Dt2DAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dt2D"},
			Value: fmt.Sprintf("%d", b2i(*m.Dt2DAttr))})
	}
	if m.DtrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dtr"},
			Value: fmt.Sprintf("%d", b2i(*m.DtrAttr))})
	}
	if m.Del1Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "del1"},
			Value: fmt.Sprintf("%d", b2i(*m.Del1Attr))})
	}
	if m.Del2Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "del2"},
			Value: fmt.Sprintf("%d", b2i(*m.Del2Attr))})
	}
	if m.R1Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r1"},
			Value: fmt.Sprintf("%v", *m.R1Attr)})
	}
	if m.R2Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r2"},
			Value: fmt.Sprintf("%v", *m.R2Attr)})
	}
	if m.CaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ca"},
			Value: fmt.Sprintf("%d", b2i(*m.CaAttr))})
	}
	if m.SiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "si"},
			Value: fmt.Sprintf("%v", *m.SiAttr)})
	}
	if m.BxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bx"},
			Value: fmt.Sprintf("%d", b2i(*m.BxAttr))})
	}
	e.EncodeElement(m.Content, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellFormula) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TAttr = ST_CellFormulaType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "t" {
			m.TAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "aca" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AcaAttr = &parsed
		}
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = &parsed
		}
		if attr.Name.Local == "dt2D" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.Dt2DAttr = &parsed
		}
		if attr.Name.Local == "dtr" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DtrAttr = &parsed
		}
		if attr.Name.Local == "del1" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.Del1Attr = &parsed
		}
		if attr.Name.Local == "del2" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.Del2Attr = &parsed
		}
		if attr.Name.Local == "r1" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.R1Attr = &parsed
		}
		if attr.Name.Local == "r2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.R2Attr = &parsed
		}
		if attr.Name.Local == "ca" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CaAttr = &parsed
		}
		if attr.Name.Local == "si" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SiAttr = &pt
		}
		if attr.Name.Local == "bx" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BxAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CellFormula: %s", err)
		}
		if cd, ok := tok.(xml.CharData); ok {
			m.Content = string(cd)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CellFormula and its children
func (m *CT_CellFormula) Validate() error {
	return m.ValidateWithPath("CT_CellFormula")
}

// ValidateWithPath validates the CT_CellFormula and its children, prefixing error messages with path
func (m *CT_CellFormula) ValidateWithPath(path string) error {
	if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
		return err
	}
	return nil
}
