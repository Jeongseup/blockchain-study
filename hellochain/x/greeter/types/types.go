package types

import (
	"fmt"
	"strings"

	codec "github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)



const (
	ModuleName = "greeter"
	StoreKey = ModuleName
)

var (
	ModuleCodec = codec.New()
	
)

type Greeting struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`     // address of the account "sending" the greeting
	Recipient sdk.AccAddress `json:"receiver" yaml:"receiver"` // address of the account "receiving" the greeting
	Body      string         `json:"body" yaml:"body"`         // string body of the greeting
}

type GreetingsList []Greeting

func NewGreeting(sender sdk.AccAddress, body string, receiver sdk.AccAddress) Greeting {
	return Greeting{
		Recipient: receiver,
		Sender: sender,
		Body: body,
	}
}

func (g Greeting) String() string {
	return strings.TrimSpace(
		fmt.Sprintf(`Sender: %s Recipient: %s Body: %s`, g.Sender.String(), g.Recipient.String(), g.Body),
	)
}

type QueryResGreetings map[string][]Greeting


func (q QueryResGreetings) String() string {
	b := ModuleCodec.MustMarshalJSON(q)
	return string(b)
}

// NewQueryResGreetings constructs a new instance
func NewQueryResGreetings() QueryResGreetings {
	return make(map[string][]Greeting)
}