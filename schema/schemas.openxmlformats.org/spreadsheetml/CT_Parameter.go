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

type CT_Parameter struct {
	// Parameter Name
	NameAttr *string
	// SQL Data Type
	SqlTypeAttr *int32
	// Parameter Type
	ParameterTypeAttr ST_ParameterType
	// Refresh on Change
	RefreshOnChangeAttr *bool
	// Parameter Prompt String
	PromptAttr *string
	// Boolean
	BooleanAttr *bool
	// Double
	DoubleAttr *float64
	// Integer
	IntegerAttr *int32
	// String
	StringAttr *string
	// Cell Reference
	CellAttr *string
}

func NewCT_Parameter() *CT_Parameter {
	ret := &CT_Parameter{}
	ret.ParameterTypeAttr = ST_ParameterTypePrompt
	ret.RefreshOnChangeAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_Parameter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.SqlTypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sqlType"},
			Value: fmt.Sprintf("%v", *m.SqlTypeAttr)})
	}
	if m.ParameterTypeAttr != ST_ParameterTypeUnset {
		attr, err := m.ParameterTypeAttr.MarshalXMLAttr(xml.Name{Local: "parameterType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RefreshOnChangeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshOnChange"},
			Value: fmt.Sprintf("%d", b2i(*m.RefreshOnChangeAttr))})
	}
	if m.PromptAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "prompt"},
			Value: fmt.Sprintf("%v", *m.PromptAttr)})
	}
	if m.BooleanAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "boolean"},
			Value: fmt.Sprintf("%d", b2i(*m.BooleanAttr))})
	}
	if m.DoubleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "double"},
			Value: fmt.Sprintf("%v", *m.DoubleAttr)})
	}
	if m.IntegerAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "integer"},
			Value: fmt.Sprintf("%v", *m.IntegerAttr)})
	}
	if m.StringAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "string"},
			Value: fmt.Sprintf("%v", *m.StringAttr)})
	}
	if m.CellAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cell"},
			Value: fmt.Sprintf("%v", *m.CellAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Parameter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ParameterTypeAttr = ST_ParameterTypePrompt
	m.RefreshOnChangeAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "sqlType" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.SqlTypeAttr = &pt
		}
		if attr.Name.Local == "parameterType" {
			m.ParameterTypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "refreshOnChange" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RefreshOnChangeAttr = &parsed
		}
		if attr.Name.Local == "prompt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PromptAttr = &parsed
		}
		if attr.Name.Local == "boolean" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BooleanAttr = &parsed
		}
		if attr.Name.Local == "double" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.DoubleAttr = &parsed
		}
		if attr.Name.Local == "integer" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.IntegerAttr = &pt
		}
		if attr.Name.Local == "string" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StringAttr = &parsed
		}
		if attr.Name.Local == "cell" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CellAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Parameter: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Parameter and its children
func (m *CT_Parameter) Validate() error {
	return m.ValidateWithPath("CT_Parameter")
}

// ValidateWithPath validates the CT_Parameter and its children, prefixing error messages with path
func (m *CT_Parameter) ValidateWithPath(path string) error {
	if err := m.ParameterTypeAttr.ValidateWithPath(path + "/ParameterTypeAttr"); err != nil {
		return err
	}
	return nil
}
