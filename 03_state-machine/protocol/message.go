package protocol

// EchoRequesrArgs is an RPC input for `EchoRequest` call
type EchoRequestArgs struct {
	Payload string
}

// EchoRequesrReply is an RPC outout for `EchoRequest` call
type EchoRequestReply struct {
	Payload string
}
