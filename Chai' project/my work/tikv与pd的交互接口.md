## tikv与pd的交互接口
```
接口通过handleStoreHeartbeat实现，handleStoreHeartbeat通过在上锁期间获取StoreStats中记录的store,然后更新storestats，将最后心跳时间store.LastHeartbeatTS更新为当前时间
to:pd\server\cache.go
```


```
其中的proto.Clone函数为心跳的核心部分代码 
to:pd\vendor\github.com\gogo\protobuf\proto\clone.go
```	

## 一些技术支持

#### 有关reflect函数

> Go reflect 应用场景实例 : https://blog.csdn.net/mrbuffoon/article/details/85637417
> 反射 | reflect : https://cloud.tencent.com/developer/section/1143989
> go语言之行--接口(interface)、反射(reflect)详解 : https://www.cnblogs.com/wdliu/p/9222283.html
> Go语言反射reflect深入理解 : https://www.linuxidc.com/Linux/2019-03/157557.htm
简介：
```
反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。

支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。
```

> example_of_reflect.go
> example_of_reflect_2.go
> example_of_reflect_3.go
> example_of_reflect_4.go

```
可以通过尝试运行代码了解代码功能
```
#### go语言自定义包

>Go语言自定义包 : c.biancheng.net/view/5123.html

```
目前大部分的帖子都说的是go语言在linux系统下的操作，windows系统下操作比在linux系统下操作稍微复杂在于，将包放在go/src路径下后，还需要再命令行中进行两步
```
```
go build
go insall
```
pd\vendor\github.com\coreos\etcd\raft\raftpb\raft.pb.go

#### go语言的ide使用尝试
> VSCode+golang 安装配置: https://blog.csdn.net/u013295518/article/details/78766086

## 困惑

#### go语言编译报错: invalid type for composite literal: proto.Message

目前进展：很可能是Message的定义比较混乱，导致编译错误
下面是进行的尝试：进入pd\vendor\github.com\gogo\protobuf\proto\clone.go,尝试寻找整个文档中出现Message的地方对Message的定义

· C:\Users\李晓桐\pd\vendor\github.com\chzyer\readline\remote.go//失败
type Message struct {
	Type MsgType
	Data []byte
}

·C:\Users\李晓桐\pd\vendor\github.com\coreos\etcd\raft\raftpb\raft.pb.go//失败
type Message struct {
	Type             MessageType `protobuf:"varint,1,opt,name=type,enum=raftpb.MessageType" json:"type"`
	To               uint64      `protobuf:"varint,2,opt,name=to" json:"to"`
	From             uint64      `protobuf:"varint,3,opt,name=from" json:"from"`
	Term             uint64      `protobuf:"varint,4,opt,name=term" json:"term"`
	LogTerm          uint64      `protobuf:"varint,5,opt,name=logTerm" json:"logTerm"`
	Index            uint64      `protobuf:"varint,6,opt,name=index" json:"index"`
	Entries          []Entry     `protobuf:"bytes,7,rep,name=entries" json:"entries"`
	Commit           uint64      `protobuf:"varint,8,opt,name=commit" json:"commit"`
	Snapshot         Snapshot    `protobuf:"bytes,9,opt,name=snapshot" json:"snapshot"`
	Reject           bool        `protobuf:"varint,10,opt,name=reject" json:"reject"`
	RejectHint       uint64      `protobuf:"varint,11,opt,name=rejectHint" json:"rejectHint"`
	Context          []byte      `protobuf:"bytes,12,opt,name=context" json:"context,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

·C:\Users\李晓桐\pd\vendor\github.com\coreos\etcd\snap\message.go
type Message struct {
	raftpb.Message
	ReadCloser io.ReadCloser
	TotalSize  int64
	closeC     chan bool
}

·C:\Users\李晓桐\pd\vendor\github.com\gogo\protobuf\proto\lib.go
type Message interface {
	Reset()
	String() string
	ProtoMessage()
}

·C:\Users\李晓桐\pd\vendor\github.com\golang\protobuf\proto\lib.go
type Message interface {
	Reset()
	String() string
	ProtoMessage()
}

·C:\Users\李晓桐\pd\vendor\github.com\juju\errors\error.go
func (e *Err) Message() string {
	return e.message
}

·C:\Users\李晓桐\pd\vendor\github.com\juju\errors\functions.go
type wrapper interface {
	// Message returns the top level error message,
	// not including the message from the Previous
	// error.
	Message() string
	// Underlying returns the Previous error, or nil
	// if there is none.
	Underlying() error
}

·C:\Users\李晓桐\pd\vendor\github.com\pingcap\kvproto\pkg\eraftpb\eraftpb.pb.go
type Message struct {
	MsgType    MessageType `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3,enum=eraftpb.MessageType" json:"msg_type,omitempty"`
	To         uint64      `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	From       uint64      `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	Term       uint64      `protobuf:"varint,4,opt,name=term,proto3" json:"term,omitempty"`
	LogTerm    uint64      `protobuf:"varint,5,opt,name=log_term,json=logTerm,proto3" json:"log_term,omitempty"`
	Index      uint64      `protobuf:"varint,6,opt,name=index,proto3" json:"index,omitempty"`
	Entries    []*Entry    `protobuf:"bytes,7,rep,name=entries" json:"entries,omitempty"`
	Commit     uint64      `protobuf:"varint,8,opt,name=commit,proto3" json:"commit,omitempty"`
	Snapshot   *Snapshot   `protobuf:"bytes,9,opt,name=snapshot" json:"snapshot,omitempty"`
	Reject     bool        `protobuf:"varint,10,opt,name=reject,proto3" json:"reject,omitempty"`
	RejectHint uint64      `protobuf:"varint,11,opt,name=reject_hint,json=rejectHint,proto3" json:"reject_hint,omitempty"`
	Context    []byte      `protobuf:"bytes,12,opt,name=context,proto3" json:"context,omitempty"`
}

·C:\Users\李晓桐\pd\vendor\google.golang.org\grpc\status\status.go
func (s *Status) Message() string {
	if s == nil || s.s == nil {
		return ""
	}
	return s.s.Message
}

