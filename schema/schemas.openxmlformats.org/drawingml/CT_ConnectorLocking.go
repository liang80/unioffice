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

type CT_ConnectorLocking struct {
	ExtLst                 *CT_OfficeArtExtensionList
	NoGrpAttr              *bool
	NoSelectAttr           *bool
	NoRotAttr              *bool
	NoChangeAspectAttr     *bool
	NoMoveAttr             *bool
	NoResizeAttr           *bool
	NoEditPointsAttr       *bool
	NoAdjustHandlesAttr    *bool
	NoChangeArrowheadsAttr *bool
	NoChangeShapeTypeAttr  *bool
}

func NewCT_ConnectorLocking() *CT_ConnectorLocking {
	ret := &CT_ConnectorLocking{}
	ret.NoGrpAttr = gooxml.Bool(false)
	ret.NoSelectAttr = gooxml.Bool(false)
	ret.NoRotAttr = gooxml.Bool(false)
	ret.NoChangeAspectAttr = gooxml.Bool(false)
	ret.NoMoveAttr = gooxml.Bool(false)
	ret.NoResizeAttr = gooxml.Bool(false)
	ret.NoEditPointsAttr = gooxml.Bool(false)
	ret.NoAdjustHandlesAttr = gooxml.Bool(false)
	ret.NoChangeArrowheadsAttr = gooxml.Bool(false)
	ret.NoChangeShapeTypeAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_ConnectorLocking) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NoGrpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noGrp"},
			Value: fmt.Sprintf("%d", b2i(*m.NoGrpAttr))})
	}
	if m.NoSelectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noSelect"},
			Value: fmt.Sprintf("%d", b2i(*m.NoSelectAttr))})
	}
	if m.NoRotAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noRot"},
			Value: fmt.Sprintf("%d", b2i(*m.NoRotAttr))})
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
	if m.NoEditPointsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noEditPoints"},
			Value: fmt.Sprintf("%d", b2i(*m.NoEditPointsAttr))})
	}
	if m.NoAdjustHandlesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noAdjustHandles"},
			Value: fmt.Sprintf("%d", b2i(*m.NoAdjustHandlesAttr))})
	}
	if m.NoChangeArrowheadsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noChangeArrowheads"},
			Value: fmt.Sprintf("%d", b2i(*m.NoChangeArrowheadsAttr))})
	}
	if m.NoChangeShapeTypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noChangeShapeType"},
			Value: fmt.Sprintf("%d", b2i(*m.NoChangeShapeTypeAttr))})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ConnectorLocking) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NoGrpAttr = gooxml.Bool(false)
	m.NoSelectAttr = gooxml.Bool(false)
	m.NoRotAttr = gooxml.Bool(false)
	m.NoChangeAspectAttr = gooxml.Bool(false)
	m.NoMoveAttr = gooxml.Bool(false)
	m.NoResizeAttr = gooxml.Bool(false)
	m.NoEditPointsAttr = gooxml.Bool(false)
	m.NoAdjustHandlesAttr = gooxml.Bool(false)
	m.NoChangeArrowheadsAttr = gooxml.Bool(false)
	m.NoChangeShapeTypeAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "noGrp" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoGrpAttr = &parsed
		}
		if attr.Name.Local == "noSelect" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoSelectAttr = &parsed
		}
		if attr.Name.Local == "noRot" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoRotAttr = &parsed
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
		if attr.Name.Local == "noEditPoints" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoEditPointsAttr = &parsed
		}
		if attr.Name.Local == "noAdjustHandles" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoAdjustHandlesAttr = &parsed
		}
		if attr.Name.Local == "noChangeArrowheads" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoChangeArrowheadsAttr = &parsed
		}
		if attr.Name.Local == "noChangeShapeType" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoChangeShapeTypeAttr = &parsed
		}
	}
lCT_ConnectorLocking:
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
				log.Printf("skipping unsupported element on CT_ConnectorLocking %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ConnectorLocking
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ConnectorLocking and its children
func (m *CT_ConnectorLocking) Validate() error {
	return m.ValidateWithPath("CT_ConnectorLocking")
}

// ValidateWithPath validates the CT_ConnectorLocking and its children, prefixing error messages with path
func (m *CT_ConnectorLocking) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
