package chain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/libp2p/go-libp2p-core/host"
	net "github.com/libp2p/go-libp2p-core/network"
)
type Node struct {
	// Port    int
	// Target	string
	// Secio 	bool
	// Seed	int64
	Ha 		host.Host
}

var mutex = &sync.Mutex{}

// makeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress. It will use secio if secio is true.
//func MakeBasicHost(listenPort int, secio bool, randseed int64) (host.Host, error) {
// func (node Node)MakeBasicHost() (host.Host, error) {
// 	// If the seed is zero, use real cryptographic randomness. Otherwise, use a
// 	// deterministic randomness source to make generated keys stay the same
// 	// across multiple runs
// 	var r io.Reader
// 	if node.Seed == 0 {
// 		r = rand.Reader
// 	} else {
// 		r = mrand.New(mrand.NewSource(node.Seed))
// 	}

// 	// Generate a key pair for this host. We will use it
// 	// to obtain a valid host ID.
// 	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	opts := []libp2p.Option{
// 		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", node.Port)),
// 		libp2p.Identity(priv),
// 	}

// 	if !node.Secio {
// 		opts = append(opts, libp2p.NoSecurity)
// 	}

// 	// basicHost, err := libp2p.New(context.Background(), opts...)
// 	basicHost, err := libp2p.New(opts...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Build host multiaddress
// 	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))

// 	// Now we can build a full multiaddress to reach this host
// 	// by encapsulating both addresses:
// 	addr := basicHost.Addrs()[0]
// 	fullAddr := addr.Encapsulate(hostAddr)
// 	log.Printf("I am %s\n", fullAddr)
// 	if node.Secio {
// 		log.Printf("Now run \"go run main.go -l %d -d %s -secio\" on a different terminal\n", node.Port+1, fullAddr)
// 	} else {
// 		log.Printf("Now run \"go run main.go -l %d -d %s\" on a different terminal\n", node.Port+1, fullAddr)
// 	}

// 	return basicHost, nil
// }
// func (node Node)MakeDiscoveryHost() (host.Host, error) {
// 	// host, err := libp2p.New()
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// return host, nil
// 	ctx := context.Background()

// 	// libp2p.New constructs a new libp2p Host.
// 	// Other options can be added here.
// 	sourceMultiAddr, _ := ma.NewMultiaddr("/ip4/0.0.0.0/tcp/4000")

// 	r := mrand.New(mrand.NewSource(int64(10)))
// 	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
// 	if err != nil {
// 		panic(err)
// 	}
// 	host, err := libp2p.New(
// 		libp2p.ListenAddrs(sourceMultiAddr),
// 		libp2p.Identity(prvKey),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("This node: ", host.ID().Pretty(), " ", host.Addrs())

// 	_, err = dht.New(ctx, host)
// 	if err != nil {
// 		panic(err)
// 	}

// 	select {}
// }
func (node Node) HandleStream(s net.Stream) {

	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	go node.ReadData(rw)
	go node.WriteData(rw)

	// stream 's' will stay open until you close it (or the other side closes it).
}
func (node Node) ReadData(rw *bufio.ReadWriter) {

	go func() {
		
		for {
			fmt.Println("READ DATA")
			str, err := rw.ReadString('\n')
			if err != nil {
				// where the stream reset occurs 
				log.Println(err)
				
			}

			if str == "" {
				return
			}
			if str != "\n" {

				chain := make([]Block, 0)
				if err := json.Unmarshal([]byte(str), &chain); err != nil {
					log.Fatal(err)
				}

				mutex.Lock()
				if len(chain) > len(Blockchain) {
					Blockchain = chain
					bytes, err := json.MarshalIndent(Blockchain, "", "  ")
					if err != nil {

						log.Fatal(err)
					}
					// Green console color: 	\x1b[32m
					// Reset console color: 	\x1b[0m %s\x1b[0m>
					fmt.Printf("\x1b[32m %s\x1b[0m>", string(bytes))
				}
				mutex.Unlock()
			}
		}
	}()
}
func (node Node) WriteData(rw *bufio.ReadWriter) {

	go func() {
		for {
			fmt.Println("WRITE DATA")
			time.Sleep(5 * time.Second)
			mutex.Lock()
			bytes, err := json.Marshal(Blockchain)
			if err != nil {
				log.Println(err)
			}
			mutex.Unlock()

			mutex.Lock()
			rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
			rw.Flush()
			mutex.Unlock()

		}
	}()

	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		sendData = strings.Replace(sendData, "\r", "", -1)
		sendData = strings.Replace(sendData, "\n", "", -1)
		bpm, err := strconv.Atoi(sendData)
		if err != nil {
			log.Fatal(err)
		}
		newBlock := GenerateBlock(Blockchain[len(Blockchain)-1], bpm)

		if IsBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
			mutex.Lock()
			Blockchain = append(Blockchain, newBlock)
			mutex.Unlock()
		}

		bytes, err := json.Marshal(Blockchain)
		if err != nil {
			log.Println(err)
		}

		spew.Dump(Blockchain)

		mutex.Lock()
		rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		rw.Flush()
		mutex.Unlock()
	}

}
