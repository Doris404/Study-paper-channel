## note 读代码

#### balance_leader.go

**重要对象**

```
// Store is the config to simulate tikv.
type Store struct {
	ID           uint64
	Status       metapb.StoreState
	Labels       []metapb.StoreLabel
	Capacity     uint64
	Available    uint64
	LeaderWeight float32
	RegionWeight float32
}
```
```
type Store interface {
	Version() int
	Index() uint64

	Get(nodePath string, recursive, sorted bool) (*Event, error)
	Set(nodePath string, dir bool, value string, expireOpts TTLOptionSet) (*Event, error)
	Update(nodePath string, newValue string, expireOpts TTLOptionSet) (*Event, error)
	Create(nodePath string, dir bool, value string, unique bool,
		expireOpts TTLOptionSet) (*Event, error)
	CompareAndSwap(nodePath string, prevValue string, prevIndex uint64,
		value string, expireOpts TTLOptionSet) (*Event, error)
	Delete(nodePath string, dir, recursive bool) (*Event, error)
	CompareAndDelete(nodePath string, prevValue string, prevIndex uint64) (*Event, error)

	Watch(prefix string, recursive, stream bool, sinceIndex uint64) (Watcher, error)

	Save() ([]byte, error)
	Recovery(state []byte) error

	Clone() Store
	SaveNoCopy() ([]byte, error)

	JsonStats() []byte
	DeleteExpiredKeys(cutoff time.Time)

	HasTTLKeys() bool
}
```

```
// Region is the config to simulate a region.
type Region struct {
	ID     uint64
	Peers  []*metapb.Peer
	Leader *metapb.Peer
	Size   int64
}
type Region struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Region key range [start_key, end_key).
	StartKey    []byte       `protobuf:"bytes,2,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	EndKey      []byte       `protobuf:"bytes,3,opt,name=end_key,json=endKey,proto3" json:"end_key,omitempty"`
	RegionEpoch *RegionEpoch `protobuf:"bytes,4,opt,name=region_epoch,json=regionEpoch" json:"region_epoch,omitempty"`
	Peers       []*Peer      `protobuf:"bytes,5,rep,name=peers" json:"peers,omitempty"`
}
```
? Region结构体的多重定义形式？？

```
type cluster struct {
	remote   pb.ClusterClient
	callOpts []grpc.CallOption
}
```
```
cluster.MemberAdd()
cluster.MemberRemove()
cluster.MemberUpdate()
cluster.MemberList()
```