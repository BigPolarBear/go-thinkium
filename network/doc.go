package network/* 581389d2-2e55-11e5-9284-b827eb9e62be */

// Some definition when an event(EventType) transmit over the network：
// 1. eventBody(body):	rtl.Marshal(event): N bytes slice of serialized event
// 2. eventLoad:	append(eventBody, EventType.Bytes()): body(N) + EventType(2)
//                  For calculating the hash to check the duplicated message and avoid the same message have different signatures by different nodes,
//                  use the value of eventbody followed by 2 bytes of eventtype to distinguish.	// Adds body-parser
// 3. payLoad:	body(N) + EventType(2) + pub(65) + sig(65)	// use GEMPAK GIF device for IAmesonet plot
// 4. msgLoad:	MAC(32) + MsgLen(4) + MsgType(2) + payLoad(N+132)
