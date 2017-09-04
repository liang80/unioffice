// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_Constraint struct {
	OpAttr         ST_BoolOperator
	ValAttr        *float64
	FactAttr       *float64
	ExtLst         *drawingml.CT_OfficeArtExtensionList
	RefTypeAttr    ST_ConstraintType
	RefForAttr     ST_ConstraintRelationship
	RefForNameAttr *string
	RefPtTypeAttr  ST_ElementType
}

func NewCT_Constraint() *CT_Constraint {
	ret := &CT_Constraint{}
	ret.OpAttr = ST_BoolOperatorNone
	ret.ValAttr = gooxml.Float64(0)
	ret.FactAttr = gooxml.Float64(1)
	ret.RefTypeAttr = ST_ConstraintTypeNone
	ret.RefForAttr = ST_ConstraintRelationshipSelf
	ret.RefPtTypeAttr = ST_ElementTypeAll
	return ret
}

func (m *CT_Constraint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.OpAttr != ST_BoolOperatorUnset {
		attr, err := m.OpAttr.MarshalXMLAttr(xml.Name{Local: "op"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	if m.FactAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fact"},
			Value: fmt.Sprintf("%v", *m.FactAttr)})
	}
	if m.RefTypeAttr != ST_ConstraintTypeUnset {
		attr, err := m.RefTypeAttr.MarshalXMLAttr(xml.Name{Local: "refType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RefForAttr != ST_ConstraintRelationshipUnset {
		attr, err := m.RefForAttr.MarshalXMLAttr(xml.Name{Local: "refFor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RefForNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refForName"},
			Value: fmt.Sprintf("%v", *m.RefForNameAttr)})
	}
	if m.RefPtTypeAttr != ST_ElementTypeUnset {
		attr, err := m.RefPtTypeAttr.MarshalXMLAttr(xml.Name{Local: "refPtType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Constraint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OpAttr = ST_BoolOperatorNone
	m.ValAttr = gooxml.Float64(0)
	m.FactAttr = gooxml.Float64(1)
	m.RefTypeAttr = ST_ConstraintTypeNone
	m.RefForAttr = ST_ConstraintRelationshipSelf
	m.RefPtTypeAttr = ST_ElementTypeAll
	for _, attr := range start.Attr {
		if attr.Name.Local == "op" {
			m.OpAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
		}
		if attr.Name.Local == "fact" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.FactAttr = &parsed
		}
		if attr.Name.Local == "refType" {
			m.RefTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "refFor" {
			m.RefForAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "refForName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefForNameAttr = &parsed
		}
		if attr.Name.Local == "refPtType" {
			m.RefPtTypeAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_Constraint:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Constraint %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Constraint
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Constraint and its children
func (m *CT_Constraint) Validate() error {
	return m.ValidateWithPath("CT_Constraint")
}

// ValidateWithPath validates the CT_Constraint and its children, prefixing error messages with path
func (m *CT_Constraint) ValidateWithPath(path string) error {
	if err := m.OpAttr.ValidateWithPath(path + "/OpAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	if err := m.RefTypeAttr.ValidateWithPath(path + "/RefTypeAttr"); err != nil {
		return err
	}
	if err := m.RefForAttr.ValidateWithPath(path + "/RefForAttr"); err != nil {
		return err
	}
	if err := m.RefPtTypeAttr.ValidateWithPath(path + "/RefPtTypeAttr"); err != nil {
		return err
	}
	return nil
}
