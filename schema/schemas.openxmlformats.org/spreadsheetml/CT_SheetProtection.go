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

type CT_SheetProtection struct {
	// Legacy Password
	PasswordAttr *string
	// Cryptographic Algorithm Name
	AlgorithmNameAttr *string
	// Password Hash Value
	HashValueAttr *string
	// Salt Value for Password Verifier
	SaltValueAttr *string
	// Iterations to Run Hashing Algorithm
	SpinCountAttr *uint32
	// Sheet Locked
	SheetAttr *bool
	// Objects Locked
	ObjectsAttr *bool
	// Scenarios Locked
	ScenariosAttr *bool
	// Format Cells Locked
	FormatCellsAttr *bool
	// Format Columns Locked
	FormatColumnsAttr *bool
	// Format Rows Locked
	FormatRowsAttr *bool
	// Insert Columns Locked
	InsertColumnsAttr *bool
	// Insert Rows Locked
	InsertRowsAttr *bool
	// Insert Hyperlinks Locked
	InsertHyperlinksAttr *bool
	// Delete Columns Locked
	DeleteColumnsAttr *bool
	// Delete Rows Locked
	DeleteRowsAttr *bool
	// Select Locked Cells Locked
	SelectLockedCellsAttr *bool
	// Sort Locked
	SortAttr *bool
	// AutoFilter Locked
	AutoFilterAttr *bool
	// Pivot Tables Locked
	PivotTablesAttr *bool
	// Select Unlocked Cells Locked
	SelectUnlockedCellsAttr *bool
}

func NewCT_SheetProtection() *CT_SheetProtection {
	ret := &CT_SheetProtection{}
	ret.SheetAttr = gooxml.Bool(false)
	ret.ObjectsAttr = gooxml.Bool(false)
	ret.ScenariosAttr = gooxml.Bool(false)
	ret.FormatCellsAttr = gooxml.Bool(true)
	ret.FormatColumnsAttr = gooxml.Bool(true)
	ret.FormatRowsAttr = gooxml.Bool(true)
	ret.InsertColumnsAttr = gooxml.Bool(true)
	ret.InsertRowsAttr = gooxml.Bool(true)
	ret.InsertHyperlinksAttr = gooxml.Bool(true)
	ret.DeleteColumnsAttr = gooxml.Bool(true)
	ret.DeleteRowsAttr = gooxml.Bool(true)
	ret.SelectLockedCellsAttr = gooxml.Bool(false)
	ret.SortAttr = gooxml.Bool(true)
	ret.AutoFilterAttr = gooxml.Bool(true)
	ret.PivotTablesAttr = gooxml.Bool(true)
	ret.SelectUnlockedCellsAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_SheetProtection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PasswordAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "password"},
			Value: fmt.Sprintf("%v", *m.PasswordAttr)})
	}
	if m.AlgorithmNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "algorithmName"},
			Value: fmt.Sprintf("%v", *m.AlgorithmNameAttr)})
	}
	if m.HashValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hashValue"},
			Value: fmt.Sprintf("%v", *m.HashValueAttr)})
	}
	if m.SaltValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "saltValue"},
			Value: fmt.Sprintf("%v", *m.SaltValueAttr)})
	}
	if m.SpinCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spinCount"},
			Value: fmt.Sprintf("%v", *m.SpinCountAttr)})
	}
	if m.SheetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheet"},
			Value: fmt.Sprintf("%d", b2i(*m.SheetAttr))})
	}
	if m.ObjectsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "objects"},
			Value: fmt.Sprintf("%d", b2i(*m.ObjectsAttr))})
	}
	if m.ScenariosAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "scenarios"},
			Value: fmt.Sprintf("%d", b2i(*m.ScenariosAttr))})
	}
	if m.FormatCellsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formatCells"},
			Value: fmt.Sprintf("%d", b2i(*m.FormatCellsAttr))})
	}
	if m.FormatColumnsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formatColumns"},
			Value: fmt.Sprintf("%d", b2i(*m.FormatColumnsAttr))})
	}
	if m.FormatRowsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formatRows"},
			Value: fmt.Sprintf("%d", b2i(*m.FormatRowsAttr))})
	}
	if m.InsertColumnsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "insertColumns"},
			Value: fmt.Sprintf("%d", b2i(*m.InsertColumnsAttr))})
	}
	if m.InsertRowsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "insertRows"},
			Value: fmt.Sprintf("%d", b2i(*m.InsertRowsAttr))})
	}
	if m.InsertHyperlinksAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "insertHyperlinks"},
			Value: fmt.Sprintf("%d", b2i(*m.InsertHyperlinksAttr))})
	}
	if m.DeleteColumnsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "deleteColumns"},
			Value: fmt.Sprintf("%d", b2i(*m.DeleteColumnsAttr))})
	}
	if m.DeleteRowsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "deleteRows"},
			Value: fmt.Sprintf("%d", b2i(*m.DeleteRowsAttr))})
	}
	if m.SelectLockedCellsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "selectLockedCells"},
			Value: fmt.Sprintf("%d", b2i(*m.SelectLockedCellsAttr))})
	}
	if m.SortAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sort"},
			Value: fmt.Sprintf("%d", b2i(*m.SortAttr))})
	}
	if m.AutoFilterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoFilter"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoFilterAttr))})
	}
	if m.PivotTablesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pivotTables"},
			Value: fmt.Sprintf("%d", b2i(*m.PivotTablesAttr))})
	}
	if m.SelectUnlockedCellsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "selectUnlockedCells"},
			Value: fmt.Sprintf("%d", b2i(*m.SelectUnlockedCellsAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SheetProtection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SheetAttr = gooxml.Bool(false)
	m.ObjectsAttr = gooxml.Bool(false)
	m.ScenariosAttr = gooxml.Bool(false)
	m.FormatCellsAttr = gooxml.Bool(true)
	m.FormatColumnsAttr = gooxml.Bool(true)
	m.FormatRowsAttr = gooxml.Bool(true)
	m.InsertColumnsAttr = gooxml.Bool(true)
	m.InsertRowsAttr = gooxml.Bool(true)
	m.InsertHyperlinksAttr = gooxml.Bool(true)
	m.DeleteColumnsAttr = gooxml.Bool(true)
	m.DeleteRowsAttr = gooxml.Bool(true)
	m.SelectLockedCellsAttr = gooxml.Bool(false)
	m.SortAttr = gooxml.Bool(true)
	m.AutoFilterAttr = gooxml.Bool(true)
	m.PivotTablesAttr = gooxml.Bool(true)
	m.SelectUnlockedCellsAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "password" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PasswordAttr = &parsed
		}
		if attr.Name.Local == "algorithmName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgorithmNameAttr = &parsed
		}
		if attr.Name.Local == "hashValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HashValueAttr = &parsed
		}
		if attr.Name.Local == "saltValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SaltValueAttr = &parsed
		}
		if attr.Name.Local == "spinCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SpinCountAttr = &pt
		}
		if attr.Name.Local == "sheet" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SheetAttr = &parsed
		}
		if attr.Name.Local == "objects" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ObjectsAttr = &parsed
		}
		if attr.Name.Local == "scenarios" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ScenariosAttr = &parsed
		}
		if attr.Name.Local == "formatCells" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FormatCellsAttr = &parsed
		}
		if attr.Name.Local == "formatColumns" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FormatColumnsAttr = &parsed
		}
		if attr.Name.Local == "formatRows" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FormatRowsAttr = &parsed
		}
		if attr.Name.Local == "insertColumns" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.InsertColumnsAttr = &parsed
		}
		if attr.Name.Local == "insertRows" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.InsertRowsAttr = &parsed
		}
		if attr.Name.Local == "insertHyperlinks" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.InsertHyperlinksAttr = &parsed
		}
		if attr.Name.Local == "deleteColumns" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DeleteColumnsAttr = &parsed
		}
		if attr.Name.Local == "deleteRows" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DeleteRowsAttr = &parsed
		}
		if attr.Name.Local == "selectLockedCells" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SelectLockedCellsAttr = &parsed
		}
		if attr.Name.Local == "sort" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SortAttr = &parsed
		}
		if attr.Name.Local == "autoFilter" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoFilterAttr = &parsed
		}
		if attr.Name.Local == "pivotTables" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PivotTablesAttr = &parsed
		}
		if attr.Name.Local == "selectUnlockedCells" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SelectUnlockedCellsAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SheetProtection: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SheetProtection and its children
func (m *CT_SheetProtection) Validate() error {
	return m.ValidateWithPath("CT_SheetProtection")
}

// ValidateWithPath validates the CT_SheetProtection and its children, prefixing error messages with path
func (m *CT_SheetProtection) ValidateWithPath(path string) error {
	return nil
}
