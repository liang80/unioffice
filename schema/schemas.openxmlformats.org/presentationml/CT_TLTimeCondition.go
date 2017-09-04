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
	"log"
)

type CT_TLTimeCondition struct {
	// Trigger Event
	EvtAttr ST_TLTriggerEvent
	// Trigger Delay
	DelayAttr *ST_TLTime
	// Target Element Trigger Choice
	TgtEl *CT_TLTimeTargetElement
	// Time Node
	Tn *CT_TLTriggerTimeNodeID
	// Runtime Node Trigger Choice
	Rtn *CT_TLTriggerRuntimeNode
}

func NewCT_TLTimeCondition() *CT_TLTimeCondition {
	ret := &CT_TLTimeCondition{}
	ret.EvtAttr = ST_TLTriggerEvent(1)
	return ret
}

func (m *CT_TLTimeCondition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.EvtAttr != ST_TLTriggerEventUnset {
		attr, err := m.EvtAttr.MarshalXMLAttr(xml.Name{Local: "evt"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DelayAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "delay"},
			Value: fmt.Sprintf("%v", *m.DelayAttr)})
	}
	e.EncodeToken(start)
	if m.TgtEl != nil {
		setgtEl := xml.StartElement{Name: xml.Name{Local: "p:tgtEl"}}
		e.EncodeElement(m.TgtEl, setgtEl)
	}
	if m.Tn != nil {
		setn := xml.StartElement{Name: xml.Name{Local: "p:tn"}}
		e.EncodeElement(m.Tn, setn)
	}
	if m.Rtn != nil {
		sertn := xml.StartElement{Name: xml.Name{Local: "p:rtn"}}
		e.EncodeElement(m.Rtn, sertn)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLTimeCondition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.EvtAttr = ST_TLTriggerEvent(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "evt" {
			m.EvtAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "delay" {
			parsed, err := ParseUnionST_TLTime(attr.Value)
			if err != nil {
				return err
			}
			m.DelayAttr = &parsed
		}
	}
lCT_TLTimeCondition:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tgtEl":
				m.TgtEl = NewCT_TLTimeTargetElement()
				if err := d.DecodeElement(m.TgtEl, &el); err != nil {
					return err
				}
			case "tn":
				m.Tn = NewCT_TLTriggerTimeNodeID()
				if err := d.DecodeElement(m.Tn, &el); err != nil {
					return err
				}
			case "rtn":
				m.Rtn = NewCT_TLTriggerRuntimeNode()
				if err := d.DecodeElement(m.Rtn, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TLTimeCondition %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTimeCondition
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLTimeCondition and its children
func (m *CT_TLTimeCondition) Validate() error {
	return m.ValidateWithPath("CT_TLTimeCondition")
}

// ValidateWithPath validates the CT_TLTimeCondition and its children, prefixing error messages with path
func (m *CT_TLTimeCondition) ValidateWithPath(path string) error {
	if err := m.EvtAttr.ValidateWithPath(path + "/EvtAttr"); err != nil {
		return err
	}
	if m.DelayAttr != nil {
		if err := m.DelayAttr.ValidateWithPath(path + "/DelayAttr"); err != nil {
			return err
		}
	}
	if m.TgtEl != nil {
		if err := m.TgtEl.ValidateWithPath(path + "/TgtEl"); err != nil {
			return err
		}
	}
	if m.Tn != nil {
		if err := m.Tn.ValidateWithPath(path + "/Tn"); err != nil {
			return err
		}
	}
	if m.Rtn != nil {
		if err := m.Rtn.ValidateWithPath(path + "/Rtn"); err != nil {
			return err
		}
	}
	return nil
}
