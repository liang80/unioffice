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
	"log"
	"strconv"

	"baliance.com/gooxml"
)

type CT_GraphicalObjectFrameLocking struct {
	NoGrpAttr          *bool
	NoDrilldownAttr    *bool
	NoSelectAttr       *bool
	NoChangeAspectAttr *bool
	NoMoveAttr         *bool
	NoResizeAttr       *bool
	ExtLst             *CT_OfficeArtExtensionList
}

func NewCT_GraphicalObjectFrameLocking() *CT_GraphicalObjectFrameLocking {
	ret := &CT_GraphicalObjectFrameLocking{}
	ret.NoGrpAttr = gooxml.Bool(false)
	ret.NoDrilldownAttr = gooxml.Bool(false)
	ret.NoSelectAttr = gooxml.Bool(false)
	ret.NoChangeAspectAttr = gooxml.Bool(false)
	ret.NoMoveAttr = gooxml.Bool(false)
	ret.NoResizeAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_GraphicalObjectFrameLocking) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NoGrpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noGrp"},
			Value: fmt.Sprintf("%d", b2i(*m.NoGrpAttr))})
	}
	if m.NoDrilldownAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noDrilldown"},
			Value: fmt.Sprintf("%d", b2i(*m.NoDrilldownAttr))})
	}
	if m.NoSelectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noSelect"},
			Value: fmt.Sprintf("%d", b2i(*m.NoSelectAttr))})
	}
	if m.NoChangeAspectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noChangeAspect"},
			Value: fmt.Sprintf("%d", b2i(*m.NoChangeAspectAttr))})
	}
	if m.NoMoveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noMove"},
			Value: fmt.Sprintf("%d", b2i(*m.NoMoveAttr))})
	}
	if m.NoResizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noResize"},
			Value: fmt.Sprintf("%d", b2i(*m.NoResizeAttr))})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GraphicalObjectFrameLocking) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NoGrpAttr = gooxml.Bool(false)
	m.NoDrilldownAttr = gooxml.Bool(false)
	m.NoSelectAttr = gooxml.Bool(false)
	m.NoChangeAspectAttr = gooxml.Bool(false)
	m.NoMoveAttr = gooxml.Bool(false)
	m.NoResizeAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "noGrp" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoGrpAttr = &parsed
		}
		if attr.Name.Local == "noDrilldown" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoDrilldownAttr = &parsed
		}
		if attr.Name.Local == "noSelect" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoSelectAttr = &parsed
		}
		if attr.Name.Local == "noChangeAspect" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoChangeAspectAttr = &parsed
		}
		if attr.Name.Local == "noMove" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoMoveAttr = &parsed
		}
		if attr.Name.Local == "noResize" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoResizeAttr = &parsed
		}
	}
lCT_GraphicalObjectFrameLocking:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_GraphicalObjectFrameLocking %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GraphicalObjectFrameLocking
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GraphicalObjectFrameLocking and its children
func (m *CT_GraphicalObjectFrameLocking) Validate() error {
	return m.ValidateWithPath("CT_GraphicalObjectFrameLocking")
}

// ValidateWithPath validates the CT_GraphicalObjectFrameLocking and its children, prefixing error messages with path
func (m *CT_GraphicalObjectFrameLocking) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
