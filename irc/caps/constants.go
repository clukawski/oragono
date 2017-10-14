// Copyright (c) 2017 Daniel Oaks <daniel@danieloaks.net>
// released under the MIT license

package caps

// Capability represents an optional feature that a client may request from the server.
type Capability string

const (
	// AccountNotify is this IRCv3 capability: http://ircv3.net/specs/extensions/account-notify-3.1.html
	AccountNotify Capability = "account-notify"
	// AccountTag is this IRCv3 capability: http://ircv3.net/specs/extensions/account-tag-3.2.html
	AccountTag Capability = "account-tag"
	// AwayNotify is this IRCv3 capability: http://ircv3.net/specs/extensions/away-notify-3.1.html
	AwayNotify Capability = "away-notify"
	// Batch is this IRCv3 capability: http://ircv3.net/specs/extensions/batch-3.2.html
	Batch Capability = "batch"
	// CapNotify is this IRCv3 capability: http://ircv3.net/specs/extensions/cap-notify-3.2.html
	CapNotify Capability = "cap-notify"
	// ChgHost is this IRCv3 capability: http://ircv3.net/specs/extensions/chghost-3.2.html
	ChgHost Capability = "chghost"
	// EchoMessage is this IRCv3 capability: http://ircv3.net/specs/extensions/echo-message-3.2.html
	EchoMessage Capability = "echo-message"
	// ExtendedJoin is this IRCv3 capability: http://ircv3.net/specs/extensions/extended-join-3.1.html
	ExtendedJoin Capability = "extended-join"
	// InviteNotify is this IRCv3 capability: http://ircv3.net/specs/extensions/invite-notify-3.2.html
	InviteNotify Capability = "invite-notify"
	// LabeledResponse is this draft IRCv3 capability: http://ircv3.net/specs/extensions/labeled-response.html
	LabeledResponse Capability = "draft/labeled-response"
	// MaxLine is this proposed capability: https://github.com/DanielOaks/ircv3-specifications/blob/master+line-lengths/extensions/line-lengths.md
	MaxLine Capability = "draft/maxline"
	// MessageTags is this draft IRCv3 capability: http://ircv3.net/specs/core/message-tags-3.3.html
	MessageTags Capability = "draft/message-tags-0.2"
	// MultiPrefix is this IRCv3 capability: http://ircv3.net/specs/extensions/multi-prefix-3.1.html
	MultiPrefix Capability = "multi-prefix"
	// Rename is this proposed capability: https://github.com/SaberUK/ircv3-specifications/blob/rename/extensions/rename.md
	Rename Capability = "draft/rename"
	// SASL is this IRCv3 capability: http://ircv3.net/specs/extensions/sasl-3.2.html
	SASL Capability = "sasl"
	// ServerTime is this IRCv3 capability: http://ircv3.net/specs/extensions/server-time-3.2.html
	ServerTime Capability = "server-time"
	// STS is this IRCv3 capability: http://ircv3.net/specs/extensions/sts.html
	STS Capability = "sts"
	// UserhostInNames is this IRCv3 capability: http://ircv3.net/specs/extensions/userhost-in-names-3.2.html
	UserhostInNames Capability = "userhost-in-names"
)

// Name returns the name of the given capability.
func (capability Capability) Name() string {
	return string(capability)
}

// Version is used to select which max version of CAP the client supports.
type Version uint

const (
	// Cap301 refers to the base CAP spec.
	Cap301 Version = 301
	// Cap302 refers to the IRCv3.2 CAP spec.
	Cap302 Version = 302
)
