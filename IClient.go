package knx

import (
	"context"

	"github.com/uoul/go-common/async"
)

type IClient interface {
	// Core Services
	Search(ctx context.Context, req SearchRequest) <-chan async.ActionResult[KnxNetIpPackage[*SearchResponse]]
	Describe(ctx context.Context, req DescriptionRequest) <-chan async.ActionResult[KnxNetIpPackage[*DescriptionResponse]]

	// Routing / Tunnelling Services
	Subscribe() async.Stream[Cemi]
	Unsubscribe(sub async.Stream[Cemi])

	// Send KnxNetIpFrame
	Send(ctx context.Context, command GroupCommand) error
}
