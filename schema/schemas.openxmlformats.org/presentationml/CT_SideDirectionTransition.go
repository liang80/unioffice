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
)

type CT_SideDirectionTransition struct {
	// Direction
	DirAttr ST_TransitionSideDirectionType
}

func NewCT_SideDirectionTransition() *CT_SideDirectionTransition {
	ret := &CT_SideDirectionTransition{}
	ret.DirAttr = ST_TransitionSideDirectionTypeL
	return ret
}

func (m *CT_SideDirectionTransition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DirAttr != ST_TransitionSideDirectionTypeUnset {
		attr, err := m.DirAttr.MarshalXMLAttr(xml.Name{Local: "dir"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SideDirectionTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.DirAttr = ST_TransitionSideDirectionTypeL
	for _, attr := range start.Attr {
		if attr.Name.Local == "dir" {
			m.DirAttr.UnmarshalXMLAttr(attr)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SideDirectionTransition: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SideDirectionTransition and its children
func (m *CT_SideDirectionTransition) Validate() error {
	return m.ValidateWithPath("CT_SideDirectionTransition")
}

// ValidateWithPath validates the CT_SideDirectionTransition and its children, prefixing error messages with path
func (m *CT_SideDirectionTransition) ValidateWithPath(path string) error {
	if err := m.DirAttr.ValidateWithPath(path + "/DirAttr"); err != nil {
		return err
	}
	return nil
}
