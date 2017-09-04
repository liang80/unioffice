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
	"strconv"

	"baliance.com/gooxml"
)

type CT_TLBuildDiagram struct {
	// Diagram Build Types
	BldAttr      ST_TLDiagramBuildType
	SpidAttr     *uint32
	GrpIdAttr    *uint32
	UiExpandAttr *bool
}

func NewCT_TLBuildDiagram() *CT_TLBuildDiagram {
	ret := &CT_TLBuildDiagram{}
	ret.BldAttr = ST_TLDiagramBuildTypeWhole
	ret.UiExpandAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_TLBuildDiagram) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BldAttr != ST_TLDiagramBuildTypeUnset {
		attr, err := m.BldAttr.MarshalXMLAttr(xml.Name{Local: "bld"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.GrpIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "grpId"},
			Value: fmt.Sprintf("%v", *m.GrpIdAttr)})
	}
	if m.UiExpandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uiExpand"},
			Value: fmt.Sprintf("%d", b2i(*m.UiExpandAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLBuildDiagram) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.BldAttr = ST_TLDiagramBuildTypeWhole
	m.UiExpandAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "bld" {
			m.BldAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "spid" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SpidAttr = &pt
		}
		if attr.Name.Local == "grpId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.GrpIdAttr = &pt
		}
		if attr.Name.Local == "uiExpand" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UiExpandAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TLBuildDiagram: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TLBuildDiagram and its children
func (m *CT_TLBuildDiagram) Validate() error {
	return m.ValidateWithPath("CT_TLBuildDiagram")
}

// ValidateWithPath validates the CT_TLBuildDiagram and its children, prefixing error messages with path
func (m *CT_TLBuildDiagram) ValidateWithPath(path string) error {
	if err := m.BldAttr.ValidateWithPath(path + "/BldAttr"); err != nil {
		return err
	}
	return nil
}
