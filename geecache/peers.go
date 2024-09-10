package geecache

import pb "./geecachepb"

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	// 针对protobuf修改的数据类型
	Get(in *pb.Request, out *pb.Response) error
}
