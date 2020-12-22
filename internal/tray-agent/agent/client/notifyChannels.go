package client

//notifyBuffLength is the buffer length for the NotifyChannel channels of a cache.
const notifyBuffLength = 100

//NotifyChannel identifies a notification channel for a specific event.
type NotifyChannel int

//NotifyChannel identifiers.
const (
	//Notification channel id for the addition of a new peer discovered.
	ChanPeerAdded NotifyChannel = iota
	//Notification channel id for the removal of an available peer.
	ChanPeerDeleted
	//Notification channel id for an update of an available peer.
	ChanPeerUpdated
	//todo the following channels will be merged into a single ChanPeering after changing the chan type of NotifyChannels from string to interface{}
	ChanPeeringOutgoingNew
	ChanPeeringOutgoingDelete
	ChanPeeringIncomingNew
	ChanPeeringIncomingDelete
)

//notifyChannelNames contains all the registered NotifyChannel managed by the AgentController.
//It is used for init and testing purposes.
var notifyChannelNames = []NotifyChannel{
	ChanPeerAdded,
	ChanPeerDeleted,
	ChanPeerUpdated,
	ChanPeeringOutgoingNew,
	ChanPeeringOutgoingDelete,
	ChanPeeringIncomingNew,
	ChanPeeringIncomingDelete,
}
