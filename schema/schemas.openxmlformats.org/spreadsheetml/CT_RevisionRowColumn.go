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

type CT_RevisionRowColumn struct {
	// Sheet Id
	SIdAttr uint32
	// End Of List
	EolAttr *bool
	// Reference
	RefAttr string
	// User Action
	ActionAttr ST_rwColActionType
	// Edge Deleted
	EdgeAttr *bool
	// Undo
	Undo []*CT_UndoInfo
	// Revised Row Column
	Rcc []*CT_RevisionCellChange
	// Revision Format
	Rfmt    []*CT_RevisionFormatting
	RIdAttr *uint32
	UaAttr  *bool
	RaAttr  *bool
}

func NewCT_RevisionRowColumn() *CT_RevisionRowColumn {
	ret := &CT_RevisionRowColumn{}
	ret.EolAttr = gooxml.Bool(false)
	ret.ActionAttr = ST_rwColActionType(1)
	ret.EdgeAttr = gooxml.Bool(false)
	ret.UaAttr = gooxml.Bool(false)
	ret.RaAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_RevisionRowColumn) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sId"},
		Value: fmt.Sprintf("%v", m.SIdAttr)})
	if m.EolAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "eol"},
			Value: fmt.Sprintf("%d", b2i(*m.EolAttr))})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
		Value: fmt.Sprintf("%v", m.RefAttr)})
	attr, err := m.ActionAttr.MarshalXMLAttr(xml.Name{Local: "action"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.EdgeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "edge"},
			Value: fmt.Sprintf("%d", b2i(*m.EdgeAttr))})
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
	if m.Undo != nil {
		seundo := xml.StartElement{Name: xml.Name{Local: "x:undo"}}
		e.EncodeElement(m.Undo, seundo)
	}
	if m.Rcc != nil {
		sercc := xml.StartElement{Name: xml.Name{Local: "x:rcc"}}
		e.EncodeElement(m.Rcc, sercc)
	}
	if m.Rfmt != nil {
		serfmt := xml.StartElement{Name: xml.Name{Local: "x:rfmt"}}
		e.EncodeElement(m.Rfmt, serfmt)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RevisionRowColumn) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.EolAttr = gooxml.Bool(false)
	m.ActionAttr = ST_rwColActionType(1)
	m.EdgeAttr = gooxml.Bool(false)
	m.UaAttr = gooxml.Bool(false)
	m.RaAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "sId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "eol" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EolAttr = &parsed
		}
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = parsed
		}
		if attr.Name.Local == "action" {
			m.ActionAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "edge" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EdgeAttr = &parsed
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
lCT_RevisionRowColumn:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "undo":
				tmp := NewCT_UndoInfo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Undo = append(m.Undo, tmp)
			case "rcc":
				tmp := NewCT_RevisionCellChange()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcc = append(m.Rcc, tmp)
			case "rfmt":
				tmp := NewCT_RevisionFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rfmt = append(m.Rfmt, tmp)
			default:
				log.Printf("skipping unsupported element on CT_RevisionRowColumn %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RevisionRowColumn
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RevisionRowColumn and its children
func (m *CT_RevisionRowColumn) Validate() error {
	return m.ValidateWithPath("CT_RevisionRowColumn")
}

// ValidateWithPath validates the CT_RevisionRowColumn and its children, prefixing error messages with path
func (m *CT_RevisionRowColumn) ValidateWithPath(path string) error {
	if m.ActionAttr == ST_rwColActionTypeUnset {
		return fmt.Errorf("%s/ActionAttr is a mandatory field", path)
	}
	if err := m.ActionAttr.ValidateWithPath(path + "/ActionAttr"); err != nil {
		return err
	}
	for i, v := range m.Undo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Undo[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rcc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rcc[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rfmt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rfmt[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
