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
)

type AG_ConstraintAttributes struct {
	TypeAttr    ST_ConstraintType
	ForAttr     ST_ConstraintRelationship
	ForNameAttr *string
	PtTypeAttr  ST_ElementType
}

func NewAG_ConstraintAttributes() *AG_ConstraintAttributes {
	ret := &AG_ConstraintAttributes{}
	ret.TypeAttr = ST_ConstraintType(1)
	ret.ForAttr = ST_ConstraintRelationshipSelf
	ret.PtTypeAttr = ST_ElementTypeAll
	return ret
}

func (m *AG_ConstraintAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_ConstraintTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ForAttr != ST_ConstraintRelationshipUnset {
		attr, err := m.ForAttr.MarshalXMLAttr(xml.Name{Local: "for"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ForNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "forName"},
			Value: fmt.Sprintf("%v", *m.ForNameAttr)})
	}
	if m.PtTypeAttr != ST_ElementTypeUnset {
		attr, err := m.PtTypeAttr.MarshalXMLAttr(xml.Name{Local: "ptType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	return nil
}

func (m *AG_ConstraintAttributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_ConstraintType(1)
	m.ForAttr = ST_ConstraintRelationshipSelf
	m.PtTypeAttr = ST_ElementTypeAll
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "for" {
			m.ForAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "forName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ForNameAttr = &parsed
		}
		if attr.Name.Local == "ptType" {
			m.PtTypeAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_ConstraintAttributes: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_ConstraintAttributes and its children
func (m *AG_ConstraintAttributes) Validate() error {
	return m.ValidateWithPath("AG_ConstraintAttributes")
}

// ValidateWithPath validates the AG_ConstraintAttributes and its children, prefixing error messages with path
func (m *AG_ConstraintAttributes) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.ForAttr.ValidateWithPath(path + "/ForAttr"); err != nil {
		return err
	}
	if err := m.PtTypeAttr.ValidateWithPath(path + "/PtTypeAttr"); err != nil {
		return err
	}
	return nil
}
