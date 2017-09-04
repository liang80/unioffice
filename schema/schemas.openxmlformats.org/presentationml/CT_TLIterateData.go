// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml"
)

type CT_TLIterateData struct {
	// Iterate Type
	TypeAttr ST_IterateType
	// Backwards
	BackwardsAttr *bool
	// Time Absolute
	TmAbs *CT_TLIterateIntervalTime
	// Time Percentage
	TmPct *CT_TLIterateIntervalPercentage
}

func NewCT_TLIterateData() *CT_TLIterateData {
	ret := &CT_TLIterateData{}
	ret.TypeAttr = ST_IterateTypeEl
	ret.BackwardsAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_TLIterateData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_IterateTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BackwardsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "backwards"},
			Value: fmt.Sprintf("%d", b2i(*m.BackwardsAttr))})
	}
	e.EncodeToken(start)
	if m.TmAbs != nil {
		setmAbs := xml.StartElement{Name: xml.Name{Local: "p:tmAbs"}}
		e.EncodeElement(m.TmAbs, setmAbs)
	}
	if m.TmPct != nil {
		setmPct := xml.StartElement{Name: xml.Name{Local: "p:tmPct"}}
		e.EncodeElement(m.TmPct, setmPct)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLIterateData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_IterateTypeEl
	m.BackwardsAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "backwards" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BackwardsAttr = &parsed
		}
	}
lCT_TLIterateData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tmAbs":
				m.TmAbs = NewCT_TLIterateIntervalTime()
				if err := d.DecodeElement(m.TmAbs, &el); err != nil {
					return err
				}
			case "tmPct":
				m.TmPct = NewCT_TLIterateIntervalPercentage()
				if err := d.DecodeElement(m.TmPct, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TLIterateData %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLIterateData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLIterateData and its children
func (m *CT_TLIterateData) Validate() error {
	return m.ValidateWithPath("CT_TLIterateData")
}

// ValidateWithPath validates the CT_TLIterateData and its children, prefixing error messages with path
func (m *CT_TLIterateData) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if m.TmAbs != nil {
		if err := m.TmAbs.ValidateWithPath(path + "/TmAbs"); err != nil {
			return err
		}
	}
	if m.TmPct != nil {
		if err := m.TmPct.ValidateWithPath(path + "/TmPct"); err != nil {
			return err
		}
	}
	return nil
}
