package main

import (
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "log"
    "io"
    "bufio"
    "os"
    "time"
)

//TODO: move vars to appropriate places
var (
    device       string = "ens8"
    snapshot_len int32  = 1024
    promiscuous  bool   = false
    err          error
    timeout      time.Duration = 5 * time.Second
    handle       *pcap.Handle
    buffer       gopacket.SerializeBuffer
	ACK			string = "ack"
	SYN_ACK		string = "sa"
	DATA		string = "data"
)

func constructZMapRoutine() chan string {


	//routine to read in from ZMap
	zmapIncoming := make(chan string)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {

			//Read from ZMap
			input, err := reader.ReadString(byte('\n'))
			if err != nil && err == io.EOF {
				return
			}
			zmapIncoming <- input
		}

	}()

    return zmapIncoming
}

func constructPcapRoutine() chan gopacket.Packet {

	//routine to read in from pcap
	pcapIncoming := make(chan gopacket.Packet)
	go func() {
		// Open device
		handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, 0) //timeout
		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()
		// Use the handle as a packet source to process all packets
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		for {

			//Read from pcap
			packet, err := packetSource.NextPacket()
			if err == io.EOF {
				return
			} else if err != nil {
				log.Println("Error:", err)
				continue
			}
			pcapIncoming <- packet
		}

	}()

    return pcapIncoming

}



func pollTimeoutRoutine( ipMeta * map[string]packet_metadata, timeoutQueue chan packet_metadata ) chan packet_metadata {

    TIMEOUT := 1*time.Second

	timeoutIncoming := make(chan packet_metadata)

    //return from timeout when packet has expired
    //go func() {
        for {
            packet := <-timeoutQueue
            //select {
            //    case packet := <-timeoutQueue:
                    //if timeout has reached, return packet.
                    //else, check that the state has updated in the meanwhile
                    //if not, put the packet back in timeoutQueue
                    if ( ((time.Now()).Sub( packet.Timestamp ) ) > TIMEOUT) {
                        timeoutIncoming <-packet
                    } else {
	                    p, ok := (*ipMeta)[packet.Saddr]
	                    if !ok {
                            continue
                        }
                        //if state hasnt changed
                        if p.ExpectedR != packet.ExpectedR {
                            continue
                        } else {
                            timeoutQueue <-packet
                        }
                    }
            //    default:
         }
    //}()
    return timeoutIncoming

}

// TimeoutQueueStuff TODO:need to move
func constructTimeoutQueue() chan packet_metadata {

    timeoutQueue := make(chan packet_metadata)
    return timeoutQueue
}


//TODO: move the ipStateMap stuff to own file

/* keeps state by storing the packet that was sent 
 * and within the packet stores the expected response */
func constructPacketStateMap() map[string]packet_metadata {

	ipMeta := make( map[string]packet_metadata )
    return ipMeta
}


func metaContains(p * packet_metadata, ipMeta *map[string]packet_metadata) bool {
	_, ok := (*ipMeta)[p.Saddr]
	if !ok {
		return false
	}
    return true
}



func main() {

	//initalize
	ipMeta := constructPacketStateMap()
    timeoutQueue := constructTimeoutQueue()

    //read in config 
    //options := parse()

    zmapIncoming := constructZMapRoutine()
    pcapIncoming := constructPcapRoutine()
    timeoutIncoming := pollTimeoutRoutine( &ipMeta,timeoutQueue )

	//read from both zmap and pcap
	for {
		select {
			case input := <-zmapIncoming:
				ackZMap( input, &ipMeta, &timeoutQueue )
			case input := <-pcapIncoming:
				handlePcap( input, &ipMeta, &timeoutQueue )
            case input := <-timeoutIncoming:
                //TODO

			default:
				//continue to non-blocking poll
		}
	}

} //end of main
