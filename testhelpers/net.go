package testhelpers

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	ma "gx/ipfs/QmNTCey11oxhb1AxDnQBRHtdhap6Ctud872NjAYPYYXPuc/go-multiaddr"
	"gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	pstore "gx/ipfs/QmRhFARzTHcFh8wUxwN5KvyTGq73FLC65EfFAhz8Ng7aGb/go-libp2p-peerstore"
	inet "gx/ipfs/QmTGxDz2CjBucFzPNTiWwzQmTWdrBnzqbqrMucDYMsjuPb/go-libp2p-net"
	"gx/ipfs/QmTu65MVbemtUxJEWgsTtzv9Zv9P8rvmqNA4eG9TrTRGYc/go-libp2p-peer"
	"gx/ipfs/QmVmDhyTTUcQXFD1rRQ64fGLMSAoaQvNH3hwuaCFAPq2hy/errors"
	smux "gx/ipfs/QmVtV1y2e8W4eQgzsP6qfSpCCZ6zWYE4m6NzJjB7iswwrT/go-stream-muxer"
	"gx/ipfs/QmZNkThpqfVXs9GNbexPrfBbXSLNYeKrE7jwFM2oqHbyqN/go-libp2p-protocol"
	msmux "gx/ipfs/QmabLh8TrJ3emfAoQk5AbqbLTbMyj7XqumMFmAFxa9epo8/go-multistream"
	"gx/ipfs/QmcCk4LZRJPAKuwY9dusFea7LckELZgo5HagErTbm39o38/go-libp2p-interface-connmgr"
	"gx/ipfs/Qmd52WKRSwrBK5gUaJKawryZQ5by6UbNB8KVW2Zy6JtbyW/go-libp2p-host"
	mh "gx/ipfs/QmerPMzPk1mJVowm8KgmoknWa4yCYvvugMPsgWmDNUvDLW/go-multihash"

	"github.com/filecoin-project/go-filecoin/types"
)

var _ host.Host = &FakeHost{}

// FakeHost is a test host.Host
type FakeHost struct {
	ConnectImpl func(context.Context, pstore.PeerInfo) error
}

// NewFakeHost constructs a FakeHost with no other parameters needed
func NewFakeHost() host.Host {
	nopfunc := func(_ context.Context, _ pstore.PeerInfo) error { return nil }
	return &FakeHost{ConnectImpl: nopfunc}
}

// minimal implementation of host.Host interface

func (fh *FakeHost) Addrs() []ma.Multiaddr              { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Close() error                       { panic("not implemented") } // nolint: golint
func (fh *FakeHost) ConnManager() ifconnmgr.ConnManager { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Connect(ctx context.Context, pi pstore.PeerInfo) error { // nolint: golint
	return fh.ConnectImpl(ctx, pi)
}
func (fh *FakeHost) ID() peer.ID                                      { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Network() inet.Network                            { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Mux() *msmux.MultistreamMuxer                     { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Peerstore() pstore.Peerstore                      { panic("not implemented") } // nolint: golint
func (fh *FakeHost) RemoveStreamHandler(protocol.ID)                  { panic("not implemented") } // nolint: golint
func (fh *FakeHost) SetStreamHandler(protocol.ID, inet.StreamHandler) { panic("not implemented") } // nolint: golint
func (fh *FakeHost) SetStreamHandlerMatch(protocol.ID, func(string) bool, inet.StreamHandler) { // nolint: golint
	panic("not implemented")
}

// NewStream is required for the host.Host interface; returns a new FakeStream.
func (fh *FakeHost) NewStream(context.Context, peer.ID, ...protocol.ID) (inet.Stream, error) { // nolint: golint
	return newFakeStream(), nil
}

var _ inet.Dialer = &FakeDialer{}

// FakeDialer is a test inet.Dialer
type FakeDialer struct {
	PeersImpl func() []peer.ID
}

// Minimal implementation of the inet.Dialer interface

// Peers returns a fake inet.Dialer PeersImpl
func (fd *FakeDialer) Peers() []peer.ID {
	return fd.PeersImpl()
}
func (fd *FakeDialer) Peerstore() pstore.Peerstore                          { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) LocalPeer() peer.ID                                   { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) DialPeer(context.Context, peer.ID) (inet.Conn, error) { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) ClosePeer(peer.ID) error                              { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) Connectedness(peer.ID) inet.Connectedness             { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) Conns() []inet.Conn                                   { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) ConnsToPeer(peer.ID) []inet.Conn                      { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) Notify(inet.Notifiee)                                 { panic("not implemented") } // nolint: golint
func (fd *FakeDialer) StopNotify(inet.Notifiee)                             { panic("not implemented") } // nolint: golint

// fakeStream is a test inet.Stream
type fakeStream struct {
	_   smux.Stream
	pid protocol.ID
}

var _ inet.Stream = &fakeStream{}

func newFakeStream() fakeStream { return fakeStream{} }

// Minimal implementation of the inet.Stream interface

func (fs fakeStream) Protocol() protocol.ID              { return fs.pid }            // nolint: golint
func (fs fakeStream) SetProtocol(id protocol.ID)         { fs.pid = id }              // nolint: golint
func (fs fakeStream) Stat() inet.Stat                    { panic("not implemented") } // nolint: golint
func (fs fakeStream) Conn() inet.Conn                    { panic("not implemented") } // nolint: golint
func (fs fakeStream) Write(_ []byte) (int, error)        { return 1, nil }            // nolint: golint
func (fs fakeStream) Read(_ []byte) (int, error)         { return 1, nil }            // nolint: golint
func (fs fakeStream) Close() error                       { return nil }               // nolint: golint
func (fs fakeStream) Reset() error                       { return nil }               // nolint: golint
func (fs fakeStream) SetDeadline(_ time.Time) error      { return nil }               // nolint: golint
func (fs fakeStream) SetReadDeadline(_ time.Time) error  { return nil }               // nolint: golint
func (fs fakeStream) SetWriteDeadline(_ time.Time) error { return nil }               // nolint: golint

// RandPeerID is a libp2p random peer ID generator.
// These peer.ID generators were copied from libp2p/go-testutil. We didn't bring in the
// whole repo as a dependency because we only need this small bit. However if we find
// ourselves using more and more pieces we should just take a dependency on it.
func RandPeerID() (peer.ID, error) {
	buf := make([]byte, 16)
	if n, err := rand.Read(buf); n != 16 || err != nil {
		if n != 16 && err == nil {
			err = errors.New("couldnt read 16 random bytes")
		}
		panic(err)
	}
	h, _ := mh.Sum(buf, mh.SHA2_256, -1)
	return peer.ID(h), nil
}

// TestFetcher is an object with the same method set as Fetcher plus a method
// for adding blocks to the source.  It is used to implement an object that
// behaves like Fetcher but does not go to the network for use in tests.
type TestFetcher struct {
	sourceBlocks map[string]*types.Block // sourceBlocks maps block cid strings to blocks.
}

// NewTestFetcher returns a TestFetcher with no source blocks.
func NewTestFetcher() *TestFetcher {
	return &TestFetcher{
		sourceBlocks: make(map[string]*types.Block),
	}
}

// AddSourceBlocks adds the input blocks to the fetcher source.
func (f *TestFetcher) AddSourceBlocks(blocks ...*types.Block) {
	for _, block := range blocks {
		f.sourceBlocks[block.Cid().String()] = block
	}
}

// GetBlocks returns any blocks in the source with matching cids.
func (f *TestFetcher) GetBlocks(ctx context.Context, cids []cid.Cid) ([]*types.Block, error) {
	var ret []*types.Block
	for _, c := range cids {
		if block, ok := f.sourceBlocks[c.String()]; ok {
			ret = append(ret, block)
		} else {
			return nil, fmt.Errorf("failed to fetch block: %s", c.String())
		}
	}
	return ret, nil
}
