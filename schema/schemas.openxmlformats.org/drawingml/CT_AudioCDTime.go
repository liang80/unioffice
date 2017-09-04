// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"baliance.com/gooxml"
)

type CT_AudioCDTime struct {
	TrackAttr uint8
	TimeAttr  *uint32
}

func NewCT_AudioCDTime() *CT_AudioCDTime {
	ret := &CT_AudioCDTime{}
	ret.TimeAttr = gooxml.Uint32(0)
	return ret
}

func (m *CT_AudioCDTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "track"},
		Value: fmt.Sprintf("%v", m.TrackAttr)})
	if m.TimeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "time"},
			Value: fmt.Sprintf("%v", *m.TimeAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AudioCDTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TimeAttr = gooxml.Uint32(0)
	for _, attr := range start.Attr {
		if attr.Name.Local == "track" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			m.TrackAttr = uint8(parsed)
		}
		if attr.Name.Local == "time" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TimeAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AudioCDTime: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_AudioCDTime and its children
func (m *CT_AudioCDTime) Validate() error {
	return m.ValidateWithPath("CT_AudioCDTime")
}

// ValidateWithPath validates the CT_AudioCDTime and its children, prefixing error messages with path
func (m *CT_AudioCDTime) ValidateWithPath(path string) error {
	return nil
}
