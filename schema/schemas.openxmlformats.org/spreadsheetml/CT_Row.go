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
	"log"
	"strconv"

	"baliance.com/gooxml"
)

type CT_Row struct {
	// Row Index
	RAttr *uint32
	// Spans
	SpansAttr *ST_CellSpans
	// Style Index
	SAttr *uint32
	// Custom Format
	CustomFormatAttr *bool
	// Row Height
	HtAttr *float64
	// Hidden
	HiddenAttr *bool
	// Custom Height
	CustomHeightAttr *bool
	// Outline Level
	OutlineLevelAttr *uint8
	// Collapsed
	CollapsedAttr *bool
	// Thick Top Border
	ThickTopAttr *bool
	// Thick Bottom
	ThickBotAttr *bool
	// Show Phonetic
	PhAttr *bool
	// Cell
	C []*CT_Cell
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Row() *CT_Row {
	ret := &CT_Row{}
	ret.SAttr = gooxml.Uint32(0)
	ret.CustomFormatAttr = gooxml.Bool(false)
	ret.HiddenAttr = gooxml.Bool(false)
	ret.CustomHeightAttr = gooxml.Bool(false)
	ret.OutlineLevelAttr = gooxml.Uint8(0)
	ret.CollapsedAttr = gooxml.Bool(false)
	ret.ThickTopAttr = gooxml.Bool(false)
	ret.ThickBotAttr = gooxml.Bool(false)
	ret.PhAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_Row) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
			Value: fmt.Sprintf("%v", *m.RAttr)})
	}
	if m.SpansAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spans"},
			Value: fmt.Sprintf("%v", *m.SpansAttr)})
	}
	if m.SAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "s"},
			Value: fmt.Sprintf("%v", *m.SAttr)})
	}
	if m.CustomFormatAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "customFormat"},
			Value: fmt.Sprintf("%d", b2i(*m.CustomFormatAttr))})
	}
	if m.HtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ht"},
			Value: fmt.Sprintf("%v", *m.HtAttr)})
	}
	if m.HiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenAttr))})
	}
	if m.CustomHeightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "customHeight"},
			Value: fmt.Sprintf("%d", b2i(*m.CustomHeightAttr))})
	}
	if m.OutlineLevelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "outlineLevel"},
			Value: fmt.Sprintf("%v", *m.OutlineLevelAttr)})
	}
	if m.CollapsedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "collapsed"},
			Value: fmt.Sprintf("%d", b2i(*m.CollapsedAttr))})
	}
	if m.ThickTopAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "thickTop"},
			Value: fmt.Sprintf("%d", b2i(*m.ThickTopAttr))})
	}
	if m.ThickBotAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "thickBot"},
			Value: fmt.Sprintf("%d", b2i(*m.ThickBotAttr))})
	}
	if m.PhAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ph"},
			Value: fmt.Sprintf("%d", b2i(*m.PhAttr))})
	}
	e.EncodeToken(start)
	if m.C != nil {
		sec := xml.StartElement{Name: xml.Name{Local: "x:c"}}
		e.EncodeElement(m.C, sec)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Row) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SAttr = gooxml.Uint32(0)
	m.CustomFormatAttr = gooxml.Bool(false)
	m.HiddenAttr = gooxml.Bool(false)
	m.CustomHeightAttr = gooxml.Bool(false)
	m.OutlineLevelAttr = gooxml.Uint8(0)
	m.CollapsedAttr = gooxml.Bool(false)
	m.ThickTopAttr = gooxml.Bool(false)
	m.ThickBotAttr = gooxml.Bool(false)
	m.PhAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RAttr = &pt
		}
		if attr.Name.Local == "spans" {
			parsed, err := ParseSliceST_CellSpans(attr.Value)
			if err != nil {
				return err
			}
			m.SpansAttr = &parsed
		}
		if attr.Name.Local == "s" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SAttr = &pt
		}
		if attr.Name.Local == "customFormat" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustomFormatAttr = &parsed
		}
		if attr.Name.Local == "ht" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.HtAttr = &parsed
		}
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
		}
		if attr.Name.Local == "customHeight" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustomHeightAttr = &parsed
		}
		if attr.Name.Local == "outlineLevel" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.OutlineLevelAttr = &pt
		}
		if attr.Name.Local == "collapsed" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CollapsedAttr = &parsed
		}
		if attr.Name.Local == "thickTop" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ThickTopAttr = &parsed
		}
		if attr.Name.Local == "thickBot" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ThickBotAttr = &parsed
		}
		if attr.Name.Local == "ph" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PhAttr = &parsed
		}
	}
lCT_Row:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "c":
				tmp := NewCT_Cell()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.C = append(m.C, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Row %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Row
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Row and its children
func (m *CT_Row) Validate() error {
	return m.ValidateWithPath("CT_Row")
}

// ValidateWithPath validates the CT_Row and its children, prefixing error messages with path
func (m *CT_Row) ValidateWithPath(path string) error {
	for i, v := range m.C {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/C[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
