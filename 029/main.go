package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func main() {
	ipAddress := "8.8.8.8"

	var rtts []time.Duration

	for i := 0; i < 4; i++ {
		rtt, err := ping(ipAddress)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		rtts = append(rtts, rtt)
		time.Sleep(1 * time.Second)
	}

	var total time.Duration
	for _, rtt := range rtts {
		total += rtt
	}
	avg := total / time.Duration(len(rtts))

	fmt.Printf("Ping statistics for %s:\n", ipAddress)
	fmt.Printf("  Packets: Sent = 4, Received = %d, Lost = %d (%.0f%% loss),\n",
		len(rtts), 4-len(rtts), float64(4-len(rtts))/4*100)
	fmt.Printf("Approximate round trip times in milli-seconds:\n")
	fmt.Printf("  Minimum = %s, Maximum = %s, Average = %s\n",
		minDuration(rtts), maxDuration(rtts), avg)
}

func ping(ipAddress string) (time.Duration, error) {
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return 0, err
	}
	defer c.Close()

	dst, err := net.ResolveIPAddr("ip4", ipAddress)
	if err != nil {
		return 0, err
	}

	message := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte("PING"),
		},
	}
	mb, err := message.Marshal(nil)
	if err != nil {
		return 0, err
	}

	start := time.Now()
	_, err = c.WriteTo(mb, dst)
	if err != nil {
		return 0, err
	}

	reply := make([]byte, 1500)
	err = c.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return 0, err
	}
	_, _, err = c.ReadFrom(reply)
	if err != nil {
		return 0, err
	}
	duration := time.Since(start)

	return duration, nil
}

func minDuration(rtts []time.Duration) time.Duration {
	min := rtts[0]
	for _, rtt := range rtts[1:] {
		if rtt < min {
			min = rtt
		}
	}
	return min
}

func maxDuration(rtts []time.Duration) time.Duration {
	max := rtts[0]
	for _, rtt := range rtts[1:] {
		if rtt > max {
			max = rtt
		}
	}
	return max
}
