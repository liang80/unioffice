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
	"strconv"

	"baliance.com/gooxml"
)

type CT_OutlinePr struct {
	// Apply Styles in Outline
	ApplyStylesAttr *bool
	// Summary Below
	SummaryBelowAttr *bool
	// Summary Right
	SummaryRightAttr *bool
	// Show Outline Symbols
	ShowOutlineSymbolsAttr *bool
}

func NewCT_OutlinePr() *CT_OutlinePr {
	ret := &CT_OutlinePr{}
	ret.ApplyStylesAttr = gooxml.Bool(false)
	ret.SummaryBelowAttr = gooxml.Bool(true)
	ret.SummaryRightAttr = gooxml.Bool(true)
	ret.ShowOutlineSymbolsAttr = gooxml.Bool(true)
	return ret
}

func (m *CT_OutlinePr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ApplyStylesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "applyStyles"},
			Value: fmt.Sprintf("%d", b2i(*m.ApplyStylesAttr))})
	}
	if m.SummaryBelowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "summaryBelow"},
			Value: fmt.Sprintf("%d", b2i(*m.SummaryBelowAttr))})
	}
	if m.SummaryRightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "summaryRight"},
			Value: fmt.Sprintf("%d", b2i(*m.SummaryRightAttr))})
	}
	if m.ShowOutlineSymbolsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showOutlineSymbols"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowOutlineSymbolsAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OutlinePr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ApplyStylesAttr = gooxml.Bool(false)
	m.SummaryBelowAttr = gooxml.Bool(true)
	m.SummaryRightAttr = gooxml.Bool(true)
	m.ShowOutlineSymbolsAttr = gooxml.Bool(true)
	for _, attr := range start.Attr {
		if attr.Name.Local == "applyStyles" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyStylesAttr = &parsed
		}
		if attr.Name.Local == "summaryBelow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SummaryBelowAttr = &parsed
		}
		if attr.Name.Local == "summaryRight" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SummaryRightAttr = &parsed
		}
		if attr.Name.Local == "showOutlineSymbols" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowOutlineSymbolsAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_OutlinePr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_OutlinePr and its children
func (m *CT_OutlinePr) Validate() error {
	return m.ValidateWithPath("CT_OutlinePr")
}

// ValidateWithPath validates the CT_OutlinePr and its children, prefixing error messages with path
func (m *CT_OutlinePr) ValidateWithPath(path string) error {
	return nil
}
