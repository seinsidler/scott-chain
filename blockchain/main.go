package main

import (
	"context"
	"flag"
	"fmt"
	"scott-chain/blockchain/chain"

	// "chain"

	// "log"
	"time"

	// "github.com/libp2p/go-libp2p"

	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	"github.com/multiformats/go-multiaddr"
)
var logger = log.Logger("rendezvous")
const difficulty int = chain.Difficulty

func main() {
	// var node chain.Node
	t := time.Now()
	genesisBlock := chain.Block{}
	genesisBlock = chain.Block{Index: 0, Timestamp: t.String(), BPM: 0, Hash: chain.CalculateHash(genesisBlock), PrevHash: "", Difficulty: difficulty, Nonce: ""}

	chain.Blockchain = append(chain.Blockchain, genesisBlock)

	// LibP2P code uses golog to log messages. They log with different
	// string IDs (i.e. "swarm"). We can control the verbosity level for
	// all loggers with:
	// golog.SetAllLoggers("DEBUG") // Change to DEBUG for extra info

	// Parse options from the command line
	// OLD STUFF ============================
	// discoverN := flag.Bool("discovery", true, "enable peer discovery")
	// listenF := flag.Int("l", 0, "wait for incoming connections")
	// target := flag.String("d", "", "target peer to dial")
	// secio := flag.Bool("secio", false, "enable secio")
	// seed := flag.Int64("seed", 0, "set random seed for id generation")
	// flag.Parse()
	// config, err := flag.Parse()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // Make a host that listens on the given multiaddress
	// if *discoverN == false {
	// 	ha, err := node.MakeBasicHost()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	} 
	// 	node = chain.Node{Port: *listenF, Target: *target, Secio: *secio, Seed: *seed, Ha: ha}

	// } else {
	// 	ha, err := node.MakeBasicHost()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	node = chain.Node{Port: *listenF, Target: *target, Secio: *secio, Seed: *seed, Ha: ha}

	// }

	// // ha, err := node.MakeBasicHost()
	// // new host for discvoery with dht
	// // ha, err := node.MakeDiscoveryHost()
	
	// // dht here it goes

	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// // node = chain.Node{Port: *listenF, Target: *target, Secio: *secio, Seed: *seed, Ha: ha}

	// if *listenF == 0 {
	// 	log.Fatal("Please provide a port to bind on with -l")
	// }

	// // Make a host that listens on the given multiaddress

	// if *target == "" && *discoverN == true {
	// 	// ha, err := node.MakeDiscoveryHost()
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	// 	node = chain.Node{Port: *listenF, Target: *target, Secio: *secio, Seed: *seed, Ha: ha}
		

	// 	log.Println("listening for connections")
	// 	// Set a stream handler on host A. /p2p/1.0.0 is
	// 	// a user-defined protocol name.
	// 	node.Ha.SetStreamHandler("/p2p/1.0.0", node.HandleStream)
	// 	// initailize the DHT with host as a peer
	// 	ctx := context.Background()
	// 	kademliaDHT, err := dht.New(ctx, node.Ha)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// // Bootstrap the DHT. In the default configuration, this spawns a Background
	// // thread that will refresh the peer table every five minutes.
	// logger.Debug("Bootstrapping the DHT")
	// if err = kademliaDHT.Bootstrap(ctx); err != nil {
	// 	panic(err)
	// }
	// 	for _, addr := range bootstrapPeers {

	// 		iaddr, _ := ipfsaddr.ParseString(addr)
		
	// 		peerinfo, _ := peerstore.InfoFromP2pAddr(iaddr.Multiaddr())
		
	// 		if err := node.Ha.Connect(ctx, *peerinfo); err != nil {
	// 			fmt.Println(err)
	// 		} else {
	// 			fmt.Println("Connection established with bootstrap node: ", *peerinfo)
	// 		}
	// 	}
	// 	select {} // hang forever
	// 	/**** This is where the listener code ends ****/
	// } else {
	// 	// ha, err := node.MakeBasicHost()
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	// 	// node = chain.Node{Port: *listenF, Target: *target, Secio: *secio, Seed: *seed, Ha: ha}

	// 	node.Ha.SetStreamHandler("/p2p/1.0.0", node.HandleStream)

	// 	// The following code extracts target's peer ID from the
	// 	// given multiaddress
	// 	ipfsaddr, err := ma.NewMultiaddr(*target)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	peerid, err := peer.Decode(pid)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	// Decapsulate the /ipfs/<peerID> part from the target
	// 	// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
	// 	targetPeerAddr, _ := ma.NewMultiaddr(
	// 		fmt.Sprintf("/ipfs/%s", peer.Encode(peerid)))
	// 	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

	// 	// We have a peer ID and a targetAddr so we add it to the peerstore
	// 	// so LibP2P knows how to contact it
	// 	node.Ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

	// 	log.Println("opening stream")
	// 	// make a new stream from host B to host A
	// 	// it should be handled on host A by the handler we set above because
	// 	// we use the same /p2p/1.0.0 protocol
	// 	s, err := node.Ha.NewStream(context.Background(), peerid, "/p2p/1.0.0")
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	// Create a buffered stream so that read and writes are non blocking.
	// 	node.HandleStream(s)

	// 	select {} // hang forever

	// }
	// OLD STUFF ============================
	log.SetAllLoggers(log.LevelWarn)
	log.SetLogLevel("rendezvous", "info")
	help := flag.Bool("h", false, "Display Help")
	config, err := ParseFlags()
	if err != nil {
		panic(err)
	}

	if *help {
		fmt.Println("This program demonstrates a simple p2p chat application using libp2p")
		fmt.Println()
		fmt.Println("Usage: Run './chat in two different terminals. Let them connect to the bootstrap nodes, announce themselves and connect to the peers")
		flag.PrintDefaults()
		return
	}

	// libp2p.New constructs a new libp2p Host. Other options can be added
	// here.
	ctx := context.Background()
	host, err := libp2p.New(ctx, libp2p.ListenAddrs([]multiaddr.Multiaddr(config.ListenAddresses)...))
	if err != nil {
		panic(err)
	}
	fmt.Printf(host.ID().String())
	// logger.Info("Host created. We are:", host.ID())
	// logger.Info(host.Addrs())

	// // Set a function as stream handler. This function is called when a peer
	// // initiates a connection and starts a stream with this peer.
	// // ctx := context.Background()
	// host.SetStreamHandler(protocol.ID(config.ProtocolID), node.HandleStream)

	// // Start a DHT, for use in peer discovery. We can't just make a new DHT
	// // client because we want each peer to maintain its own local copy of the
	// // DHT, so that the bootstrapping node of the DHT can go down without
	// // inhibiting future peer discovery.
	// // ctx := context.Background()
	// kademliaDHT, err := dht.New(ctx, host)
	// if err != nil {
	// 	panic(err)
	// }

	// // Bootstrap the DHT. In the default configuration, this spawns a Background
	// // thread that will refresh the peer table every five minutes.
	// logger.Debug("Bootstrapping the DHT")
	// if err = kademliaDHT.Bootstrap(ctx); err != nil {
	// 	panic(err)
	// }

	// // Let's connect to the bootstrap nodes first. They will tell us about the
	// // other nodes in the network.
	// var wg sync.WaitGroup
	// for _, peerAddr := range config.BootstrapPeers {
	// 	peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		if err := host.Connect(ctx, *peerinfo); err != nil {
	// 			logger.Warning(err)
	// 		} else {
	// 			logger.Info("Connection established with bootstrap node:", *peerinfo)
	// 		}
	// 	}()
	// }
	// wg.Wait()

	// // We use a rendezvous point "meet me here" to announce our location.
	// // This is like telling your friends to meet you at the Eiffel Tower.
	// logger.Info("Announcing ourselves...")
	// routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
	// discovery.Advertise(ctx, routingDiscovery, config.RendezvousString)
	// logger.Debug("Successfully announced!")

	// // Now, look for others who have announced
	// // This is like your friend telling you the location to meet you.
	// logger.Debug("Searching for other peers...")
	// peerChan, err := routingDiscovery.FindPeers(ctx, config.RendezvousString)
	// if err != nil {
	// 	panic(err)
	// }

	// for peer := range peerChan {
	// 	if peer.ID == host.ID() {
	// 		continue
	// 	}
	// 	logger.Debug("Found peer:", peer)

	// 	logger.Debug("Connecting to:", peer)
	// 	stream, err := host.NewStream(ctx, peer.ID, protocol.ID(config.ProtocolID))

	// 	if err != nil {
	// 		logger.Warning("Connection failed:", err)
	// 		continue
	// 	} else {
	// 		// rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// 		// go writeData(rw)
	// 		// go readData(rw)
	// 		node.HandleStream(stream)
	// 	}

	// 	logger.Info("Connected to:", peer)
	// }

	// select {}
}
