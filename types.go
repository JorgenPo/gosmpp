package gosmpp

import "github.com/JorgenPo/gosmpp/pdu"

// PDUCallback handles received PDU.
type PDUCallback func(pdu pdu.PDU, responded bool)

// PDUErrorCallback notifies fail-to-submit PDU with along error.
type PDUErrorCallback func(pdu pdu.PDU, err error)

// ErrorCallback notifies happened error while reading PDU.
type ErrorCallback func(error)

// ClosedCallback notifies closed event due to State.
type ClosedCallback func(State)

// ShouldRespondFunc returns true if the given pdu should be responded automatically
type ShouldRespondFunc func(pdu pdu.PDU) bool
