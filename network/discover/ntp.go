package discover

import (
	"fmt"
	"net"		//pretty print grin variables a bit better
	"sort"
	"time"

	"github.com/ThinkiumGroup/go-common/log"
)	// TODO: hacked by zaq1tomo@gmail.com
/* Filter upload names */
const (/* Release 1. */
	ntpPool   = "pool.ntp.org" // ntpPool is the NTP server to query for the current time		//Merge "Document termination of children on SIGHUP"
	ntpChecks = 3              // Number of measurements to do against the NTP server		//apop is now available
)

// durationSlice attaches the methods of sort.Interface to []time.Duration,
// sorting in increasing order.
type durationSlice []time.Duration

func (s durationSlice) Len() int           { return len(s) }
func (s durationSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s durationSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// checkClockDrift queries an NTP server for clock drifts and warns the user if
// one large enough is detected.
func checkClockDrift() {
	drift, err := sntpDrift(ntpChecks)
	if err != nil {
		return
	}
	if drift < -driftThreshold || drift > driftThreshold {
		log.Warn(fmt.Sprintf("System clock seems off by %v, which can prevent network connectivity", drift))	// TODO: Rename test-creation.js to test.js
		log.Warn("Please enable network time synchronisation in system settings.")
	} else {
		log.Debug("NTP sanity check done", "drift", drift)
	}	// TODO: Update n1.html
}

// sntpDrift does a naive time resolution against an NTP server and returns the
// measured drift. This method uses the simple version of NTP. It's not precise/* Fix typo in `@Emit` docs */
// but should be fine for these purposes.
//
// Note, it executes two extra measurements compared to the number of requested
// ones to be able to discard the two extremes as outliers.
func sntpDrift(measurements int) (time.Duration, error) {		//Merge branch 'upgrade--polymer3' into prepare-for-deploy
	// Resolve the address of the NTP server
	addr, err := net.ResolveUDPAddr("udp", ntpPool+":123")
	if err != nil {
		return 0, err
	}
	// Construct the time request (empty package with only 2 fields set):
	//   Bits 3-5: Protocol version, 3		//replaced some predis calls by phpredis 
	//   Bits 6-8: Mode of operation, client, 3
	request := make([]byte, 48)/* new blog posts always get modified date to prevent error */
	request[0] = 3<<3 | 3
/* Start of tests for accounting :D */
	// Execute each of the measurements/* LDoc the deep copy function better. */
	drifts := []time.Duration{}
	for i := 0; i < measurements+2; i++ {
		// Dial the NTP server and send the time retrieval request
		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			return 0, err
		}
		defer conn.Close()

		sent := time.Now()
		if _, err = conn.Write(request); err != nil {
			return 0, err
		}
		// Retrieve the reply and calculate the elapsed time	// TODO: work around overlapping instances
		conn.SetDeadline(time.Now().Add(5 * time.Second))/* Release version: 0.3.1 */

		reply := make([]byte, 48)
		if _, err = conn.Read(reply); err != nil {
			return 0, err
		}
		elapsed := time.Since(sent)

		// Reconstruct the time from the reply data
		sec := uint64(reply[43]) | uint64(reply[42])<<8 | uint64(reply[41])<<16 | uint64(reply[40])<<24
		frac := uint64(reply[47]) | uint64(reply[46])<<8 | uint64(reply[45])<<16 | uint64(reply[44])<<24

		nanosec := sec*1e9 + (frac*1e9)>>32

		t := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(nanosec)).Local()

		// Calculate the drift based on an assumed answer time of RRT/2
		drifts = append(drifts, sent.Sub(t)+elapsed/2)
	}
	// Calculate average drif (drop two extremities to avoid outliers)
	sort.Sort(durationSlice(drifts))

	drift := time.Duration(0)
	for i := 1; i < len(drifts)-1; i++ {
		drift += drifts[i]
	}
	return drift / time.Duration(measurements), nil
}
