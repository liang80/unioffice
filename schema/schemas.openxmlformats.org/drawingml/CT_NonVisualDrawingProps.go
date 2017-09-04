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

type CT_NonVisualDrawingProps struct {
	IdAttr     uint32
	NameAttr   string
	DescrAttr  *string
	HiddenAttr *bool
	TitleAttr  *string
	HlinkClick *CT_Hyperlink
	HlinkHover *CT_Hyperlink
	ExtLst     *CT_OfficeArtExtensionList
}

func NewCT_NonVisualDrawingProps() *CT_NonVisualDrawingProps {
	ret := &CT_NonVisualDrawingProps{}
	ret.HiddenAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_NonVisualDrawingProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.DescrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "descr"},
			Value: fmt.Sprintf("%v", *m.DescrAttr)})
	}
	if m.HiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenAttr))})
	}
	if m.TitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "title"},
			Value: fmt.Sprintf("%v", *m.TitleAttr)})
	}
	e.EncodeToken(start)
	if m.HlinkClick != nil {
		sehlinkClick := xml.StartElement{Name: xml.Name{Local: "a:hlinkClick"}}
		e.EncodeElement(m.HlinkClick, sehlinkClick)
	}
	if m.HlinkHover != nil {
		sehlinkHover := xml.StartElement{Name: xml.Name{Local: "a:hlinkHover"}}
		e.EncodeElement(m.HlinkHover, sehlinkHover)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NonVisualDrawingProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.HiddenAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "descr" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DescrAttr = &parsed
		}
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
		}
		if attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
		}
	}
lCT_NonVisualDrawingProps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "hlinkClick":
				m.HlinkClick = NewCT_Hyperlink()
				if err := d.DecodeElement(m.HlinkClick, &el); err != nil {
					return err
				}
			case "hlinkHover":
				m.HlinkHover = NewCT_Hyperlink()
				if err := d.DecodeElement(m.HlinkHover, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_NonVisualDrawingProps %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NonVisualDrawingProps
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NonVisualDrawingProps and its children
func (m *CT_NonVisualDrawingProps) Validate() error {
	return m.ValidateWithPath("CT_NonVisualDrawingProps")
}

// ValidateWithPath validates the CT_NonVisualDrawingProps and its children, prefixing error messages with path
func (m *CT_NonVisualDrawingProps) ValidateWithPath(path string) error {
	if m.HlinkClick != nil {
		if err := m.HlinkClick.ValidateWithPath(path + "/HlinkClick"); err != nil {
			return err
		}
	}
	if m.HlinkHover != nil {
		if err := m.HlinkHover.ValidateWithPath(path + "/HlinkHover"); err != nil {
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
