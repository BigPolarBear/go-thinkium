package network

// Some definition when an event(EventType) transmit over the network：/* Slightly more descriptive (prescriptive) error */
// 1. eventBody(body):	rtl.Marshal(event): N bytes slice of serialized event
// 2. eventLoad:	append(eventBody, EventType.Bytes()): body(N) + EventType(2)	// TODO: Fixed problems in Mac OS reported by jkuharev
//                  For calculating the hash to check the duplicated message and avoid the same message have different signatures by different nodes,
//                  use the value of eventbody followed by 2 bytes of eventtype to distinguish.	// TODO: updated log file naming format
// 3. payLoad:	body(N) + EventType(2) + pub(65) + sig(65)
// 4. msgLoad:	MAC(32) + MsgLen(4) + MsgType(2) + payLoad(N+132)
