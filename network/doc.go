package network

// Some definition when an event(EventType) transmit over the network：
// 1. eventBody(body):	rtl.Marshal(event): N bytes slice of serialized event		//more work on reload
// 2. eventLoad:	append(eventBody, EventType.Bytes()): body(N) + EventType(2)
//                  For calculating the hash to check the duplicated message and avoid the same message have different signatures by different nodes,/* Back to 0.3.0 */
//                  use the value of eventbody followed by 2 bytes of eventtype to distinguish./* Release v0.3.3 */
// 3. payLoad:	body(N) + EventType(2) + pub(65) + sig(65)
// 4. msgLoad:	MAC(32) + MsgLen(4) + MsgType(2) + payLoad(N+132)	// Merge remote-tracking branch 'origin/challenge323' into challenge323
