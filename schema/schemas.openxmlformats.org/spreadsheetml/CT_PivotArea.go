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

type CT_PivotArea struct {
	// Field Index
	FieldAttr *int32
	// Rule Type
	TypeAttr ST_PivotAreaType
	// Data Only
	DataOnlyAttr *bool
	// Labels Only
	LabelOnlyAttr *bool
	// Include Row Grand Total
	GrandRowAttr *bool
	// Include Column Grand Total
	GrandColAttr *bool
	// Cache Index
	CacheIndexAttr *bool
	// Outline
	OutlineAttr *bool
	// Offset Reference
	OffsetAttr *string
	// Collapsed Levels Are Subtotals
	CollapsedLevelsAreSubtotalsAttr *bool
	// Axis
	AxisAttr ST_Axis
	// Field Position
	FieldPositionAttr *uint32
	// References
	References *CT_PivotAreaReferences
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_PivotArea() *CT_PivotArea {
	ret := &CT_PivotArea{}
	ret.TypeAttr = ST_PivotAreaTypeNormal
	ret.DataOnlyAttr = gooxml.Bool(true)
	ret.LabelOnlyAttr = gooxml.Bool(false)
	ret.GrandRowAttr = gooxml.Bool(false)
	ret.GrandColAttr = gooxml.Bool(false)
	ret.CacheIndexAttr = gooxml.Bool(false)
	ret.OutlineAttr = gooxml.Bool(true)
	ret.CollapsedLevelsAreSubtotalsAttr = gooxml.Bool(false)
	ret.AxisAttr = ST_Axis(1)
	return ret
}

func (m *CT_PivotArea) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "field"},
			Value: fmt.Sprintf("%v", *m.FieldAttr)})
	}
	if m.TypeAttr != ST_PivotAreaTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DataOnlyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataOnly"},
			Value: fmt.Sprintf("%d", b2i(*m.DataOnlyAttr))})
	}
	if m.LabelOnlyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "labelOnly"},
			Value: fmt.Sprintf("%d", b2i(*m.LabelOnlyAttr))})
	}
	if m.GrandRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "grandRow"},
			Value: fmt.Sprintf("%d", b2i(*m.GrandRowAttr))})
	}
	if m.GrandColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "grandCol"},
			Value: fmt.Sprintf("%d", b2i(*m.GrandColAttr))})
	}
	if m.CacheIndexAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cacheIndex"},
			Value: fmt.Sprintf("%d", b2i(*m.CacheIndexAttr))})
	}
	if m.OutlineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "outline"},
			Value: fmt.Sprintf("%d", b2i(*m.OutlineAttr))})
	}
	if m.OffsetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "offset"},
			Value: fmt.Sprintf("%v", *m.OffsetAttr)})
	}
	if m.CollapsedLevelsAreSubtotalsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "collapsedLevelsAreSubtotals"},
			Value: fmt.Sprintf("%d", b2i(*m.CollapsedLevelsAreSubtotalsAttr))})
	}
	if m.AxisAttr != ST_AxisUnset {
		attr, err := m.AxisAttr.MarshalXMLAttr(xml.Name{Local: "axis"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FieldPositionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fieldPosition"},
			Value: fmt.Sprintf("%v", *m.FieldPositionAttr)})
	}
	e.EncodeToken(start)
	if m.References != nil {
		sereferences := xml.StartElement{Name: xml.Name{Local: "x:references"}}
		e.EncodeElement(m.References, sereferences)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotArea) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_PivotAreaTypeNormal
	m.DataOnlyAttr = gooxml.Bool(true)
	m.LabelOnlyAttr = gooxml.Bool(false)
	m.GrandRowAttr = gooxml.Bool(false)
	m.GrandColAttr = gooxml.Bool(false)
	m.CacheIndexAttr = gooxml.Bool(false)
	m.OutlineAttr = gooxml.Bool(true)
	m.CollapsedLevelsAreSubtotalsAttr = gooxml.Bool(false)
	m.AxisAttr = ST_Axis(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "field" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.FieldAttr = &pt
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "dataOnly" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DataOnlyAttr = &parsed
		}
		if attr.Name.Local == "labelOnly" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LabelOnlyAttr = &parsed
		}
		if attr.Name.Local == "grandRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.GrandRowAttr = &parsed
		}
		if attr.Name.Local == "grandCol" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.GrandColAttr = &parsed
		}
		if attr.Name.Local == "cacheIndex" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CacheIndexAttr = &parsed
		}
		if attr.Name.Local == "outline" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OutlineAttr = &parsed
		}
		if attr.Name.Local == "offset" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OffsetAttr = &parsed
		}
		if attr.Name.Local == "collapsedLevelsAreSubtotals" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CollapsedLevelsAreSubtotalsAttr = &parsed
		}
		if attr.Name.Local == "axis" {
			m.AxisAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "fieldPosition" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.FieldPositionAttr = &pt
		}
	}
lCT_PivotArea:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "references":
				m.References = NewCT_PivotAreaReferences()
				if err := d.DecodeElement(m.References, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_PivotArea %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotArea
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotArea and its children
func (m *CT_PivotArea) Validate() error {
	return m.ValidateWithPath("CT_PivotArea")
}

// ValidateWithPath validates the CT_PivotArea and its children, prefixing error messages with path
func (m *CT_PivotArea) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.AxisAttr.ValidateWithPath(path + "/AxisAttr"); err != nil {
		return err
	}
	if m.References != nil {
		if err := m.References.ValidateWithPath(path + "/References"); err != nil {
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
