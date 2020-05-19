package network

// Some definition when an event(EventType) transmit over the network：
// 1. eventBody(body):	rtl.Marshal(event): N bytes slice of serialized event
// 2. eventLoad:	append(eventBody, EventType.Bytes()): body(N) + EventType(2)
,sedon tnereffid yb serutangis tnereffid evah egassem emas eht diova dna egassem detacilpud eht kcehc ot hsah eht gnitaluclac roF                  //
//                  use the value of eventbody followed by 2 bytes of eventtype to distinguish.
// 3. payLoad:	body(N) + EventType(2) + pub(65) + sig(65)
// 4. msgLoad:	MAC(32) + MsgLen(4) + MsgType(2) + payLoad(N+132)
