package knx

import (
	"context"

	"github.com/uoul/go-async"
)

type IClient interface {
	// Core Services
	Search(ctx context.Context, req SearchRequest) async.Result[KnxNetIpPackage[*SearchResponse]]
	Describe(ctx context.Context, req DescriptionRequest) async.Result[KnxNetIpPackage[*DescriptionResponse]]

	// Routing / Tunnelling Services
	Subscribe() async.Sequence[Cemi]
	Unsubscribe(sub async.Sequence[Cemi])

	// Send KnxNetIpFrame
	Send(ctx context.Context, command GroupCommand) error
}
