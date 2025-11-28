package server

const (
	Message = "NYGmessagePort: "

	// covert power mechanism token
	underTest  = "UnderTest: true"
	_underTest = "UnderTest: false"
	CovertUse  = "Covert: true"
	_CovertUse = "Covert: false"

	// end

	// draw power mechanism token
	drawOffer       = "DrawOffer: true"
	_drawOffer      = "DrawOffer: false"
	_decline        = "DrawOfferDeclined: true "
	DrawUse         = "Draw: true"
	_DrawUse        = "Draw: false"
	DrawMeetingDone = "DrawMeetingDone: true"
	DMsg            = "DrawMsg: "
	//end

	// rewind power mechanism token
	backClock  = "ClockRestart: true"
	_backClock = "ClockRestart: false"
	RewindUse  = "Rewind: true"
	_RewindUse = "Rewind: false"

	// end

	// tag power mechanism token
	tagIn   = Lock
	_tagIn  = Unlock
	TagUse  = "Tag: true"
	_TagUse = "Tag: false"

	Alert     = "AlertTeam: " // team name nickname has switched with team name nickname
	TMsg      = "TagMsg: true"
	TaggedIn  = "TAGGEDIN"
	TaggedOut = "TAGGEDOUT"
	//end

	// nexus power mechanism token
	nexusWord = "NexusWord: "
	NexusUse  = "Nexus: true"
	_NexusUse = "Nexus: false"
)
