// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
)

type CT_Proof struct {
	// Spell Checking State
	SpellingAttr ST_Proof
	// Grammatical Checking State
	GrammarAttr ST_Proof
}

func NewCT_Proof() *CT_Proof {
	ret := &CT_Proof{}
	ret.SpellingAttr = ST_Proof(1)
	ret.GrammarAttr = ST_Proof(1)
	return ret
}

func (m *CT_Proof) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SpellingAttr != ST_ProofUnset {
		attr, err := m.SpellingAttr.MarshalXMLAttr(xml.Name{Local: "w:spelling"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.GrammarAttr != ST_ProofUnset {
		attr, err := m.GrammarAttr.MarshalXMLAttr(xml.Name{Local: "w:grammar"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Proof) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SpellingAttr = ST_Proof(1)
	m.GrammarAttr = ST_Proof(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "spelling" {
			m.SpellingAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "grammar" {
			m.GrammarAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Proof: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Proof and its children
func (m *CT_Proof) Validate() error {
	return m.ValidateWithPath("CT_Proof")
}

// ValidateWithPath validates the CT_Proof and its children, prefixing error messages with path
func (m *CT_Proof) ValidateWithPath(path string) error {
	if err := m.SpellingAttr.ValidateWithPath(path + "/SpellingAttr"); err != nil {
		return err
	}
	if err := m.GrammarAttr.ValidateWithPath(path + "/GrammarAttr"); err != nil {
		return err
	}
	return nil
}
