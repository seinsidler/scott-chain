module scott-chain

go 1.16

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/gogo/protobuf v1.3.2
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-datastore v0.4.6
	github.com/ipfs/go-detect-race v0.0.1
	github.com/ipfs/go-ipfs-util v0.0.2
	github.com/ipfs/go-log/v2 v2.3.0
	github.com/jbenet/goprocess v0.1.4
	github.com/libp2p/go-addr-util v0.1.0
	github.com/libp2p/go-buffer-pool v0.0.2
	github.com/libp2p/go-conn-security-multistream v0.3.0
	github.com/libp2p/go-eventbus v0.2.1
	github.com/libp2p/go-libp2p v0.15.0-rc.1.0.20210925144944-eba91ac63ec9 // indirect
	// github.com/libp2p/go-libp2p v0.14.4
	// github.com/libp2p/go-libp2p master
	// scott-chain/
	github.com/libp2p/go-libp2p-asn-util v0.0.0-20210818120414-1f382a4aa43a
	github.com/libp2p/go-libp2p-autonat v0.5.0
	github.com/libp2p/go-libp2p-blankhost v0.2.0
	github.com/libp2p/go-libp2p-circuit v0.4.0
	github.com/libp2p/go-libp2p-core v0.10.0
	github.com/libp2p/go-libp2p-discovery v0.5.1
	github.com/libp2p/go-libp2p-kad-dht v0.13.1
	github.com/libp2p/go-libp2p-mplex v0.4.1
	github.com/libp2p/go-libp2p-nat v0.0.6
	github.com/libp2p/go-libp2p-netutil v0.1.0
	github.com/libp2p/go-libp2p-noise v0.3.0
	github.com/libp2p/go-libp2p-peerstore v0.2.8
	github.com/libp2p/go-libp2p-quic-transport v0.13.0
	github.com/libp2p/go-libp2p-swarm v0.6.0
	github.com/libp2p/go-libp2p-testing v0.5.0
	github.com/libp2p/go-libp2p-tls v0.3.0
	github.com/libp2p/go-libp2p-transport-upgrader v0.5.0
	github.com/libp2p/go-libp2p-yamux v0.5.4
	github.com/libp2p/go-msgio v0.0.6
	github.com/libp2p/go-netroute v0.1.6
	github.com/libp2p/go-stream-muxer-multistream v0.3.0
	github.com/libp2p/go-tcp-transport v0.2.8
	github.com/libp2p/go-ws-transport v0.5.0
	github.com/libp2p/zeroconf/v2 v2.0.0
	github.com/multiformats/go-multiaddr v0.4.0
	github.com/multiformats/go-multiaddr-dns v0.3.1
	github.com/multiformats/go-multistream v0.2.2
	github.com/multiformats/go-varint v0.0.6
	github.com/stretchr/testify v1.7.0
	github.com/whyrusleeping/mdns v0.0.0-20190826153040-b9b60ed33aa9
)

// Ensure that examples always use the go-libp2p version in the same git checkout.
// replace github.com/libp2p/go-libp2p => ../
