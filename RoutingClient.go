package knx

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/uoul/go-common/async"
	"github.com/uoul/go-common/log"
)

type RoutingClient struct {
	logger        log.ILogger
	multicastAddr string
	retryInterval time.Duration

	conn    *net.UDPConn
	subs    map[async.Stream[Cemi]]bool
	subMux  sync.RWMutex
	connMux sync.RWMutex
}

// ------------------------------------------------------------------------------------
// Public
// ------------------------------------------------------------------------------------

// Describe implements IClient.
func (r *RoutingClient) Describe(ctx context.Context, req DescriptionRequest) <-chan async.ActionResult[KnxNetIpPackage[*DescriptionResponse]] {
	panic("unimplemented")
}

// Search implements IClient.
func (r *RoutingClient) Search(ctx context.Context, req SearchRequest) <-chan async.ActionResult[KnxNetIpPackage[*SearchResponse]] {
	panic("unimplemented")
}

// Send implements IClient.
func (r *RoutingClient) Send(ctx context.Context, command GroupCommand) error {
	r.connMux.RLock()
	defer r.connMux.RUnlock()
	if r.conn == nil {
		return NewErrNetConnection("no udp connection available")
	}
	// TODO: Implement Frame conversion
	return NewErrNotImplemented("not implemented")
}

// Subscribe implements IClient.
func (r *RoutingClient) Subscribe() async.Stream[Cemi] {
	r.subMux.Lock()
	defer r.subMux.Unlock()
	sub := async.NewStream[Cemi]()
	r.subs[sub] = true
	return sub
}

// Unsubscribe implements IClient.
func (r *RoutingClient) Unsubscribe(sub async.Stream[Cemi]) {
	r.subMux.Lock()
	defer r.subMux.Unlock()
	delete(r.subs, sub)
}

// ------------------------------------------------------------------------------------
// Private
// ------------------------------------------------------------------------------------

func readUdp(mux *sync.RWMutex, conn *net.UDPConn) async.Stream[[]byte] {
	ch := async.NewStream[[]byte]()
	go func() {
		defer close(ch)
		for {
			buffer := make([]byte, 1024)
			mux.RLock()
			n, err := conn.Read(buffer)
			mux.RUnlock()
			if err != nil {
				ch <- async.NewErrorActionResult[[]byte](
					NewErrNetConnection("failed to read from udp - %v", err),
				)
				return
			}
			ch <- async.ActionResult[[]byte]{
				Result: buffer[:n],
				Error:  nil,
			}
		}
	}()
	return ch
}

func (k *RoutingClient) notify(p Cemi) {
	k.subMux.RLock()
	defer k.subMux.RUnlock()
	for sub := range k.subs {
		sub <- async.ActionResult[Cemi]{
			Result: p,
			Error:  nil,
		}
	}
}

func (k *RoutingClient) run(ctx context.Context) error {
	// Create UDP address
	addr, err := net.ResolveUDPAddr("udp4", k.multicastAddr)
	if err != nil {
		return NewErrAddressResolution("failed to create udp address for %s - %v", k.multicastAddr, err)
	}
	// Connect to multicast address
	k.conn, err = net.ListenMulticastUDP("udp4", nil, addr)
	if err != nil {
		return NewErrNetConnection("failed to listen on %s - %v", addr.String(), err)
	}
	defer func() {
		k.conn.Close()
		k.connMux.Lock()
		k.conn = nil
		k.connMux.Unlock()
	}()
	// Create Reader
	reader := readUdp(&k.connMux, k.conn)
	k.logger.Debugf("KnxRoutingClient(%s) listening...", k.multicastAddr)
	// Run
	for {
		select {
		case <-ctx.Done():
			return nil
		case data := <-reader:
			// Check for error
			if data.Error != nil {
				return data.Error
			}
			// Parse Header
			header, _, err := parseKnxNetIpHeader(data.Result[:6])
			if err != nil {
				k.logger.Warningf("%v", err)
				continue
			}
			// Process Package depending on ServiceType
			switch header.ServiceType {
			case SEARCH_REQUEST:
				err = k.processSearchReq(data.Result)
			case SEARCH_RESPONSE:
				err = k.processSearchRes(data.Result)
			case DESCRIPTION_REQUEST:
				err = k.processDescriptionReq(data.Result)
			case DESCRIPTION_RESPONSE:
				err = k.processDescriptionResp(data.Result)
			case ROUTING_INDICATION:
				err = k.processRoutingIndication(data.Result)
			case ROUTING_LOST_MESSAGE:
				err = k.processRoutingLostMsg(data.Result)
			case ROUTING_BUSY:
				err = k.processRoutingBusy(data.Result)
			default:
				continue // ignore
			}
			if err != nil {
				k.logger.Warningf("%v", err)
				continue
			}
		}
	}
}

func (r *RoutingClient) processSearchReq(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

func (r *RoutingClient) processSearchRes(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

func (r *RoutingClient) processDescriptionReq(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

func (r *RoutingClient) processDescriptionResp(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

func (r *RoutingClient) processRoutingIndication(data []byte) error {
	msg, err := parseKnxNetIpPackage[*RoutingIndication](data)
	if err != nil {
		return err
	}
	r.notify(msg.Body.Cemi)
	return nil
}

func (r *RoutingClient) processRoutingLostMsg(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

func (r *RoutingClient) processRoutingBusy(data []byte) error {
	return NewErrNotImplemented("not implemented")
}

// ------------------------------------------------------------------------------------
// Options
// ------------------------------------------------------------------------------------

func WithRoutingClientMulticastAddr(addr string) func(*RoutingClient) {
	return func(rc *RoutingClient) {
		rc.multicastAddr = addr
	}
}

func WithRoutingClientRetryInterval(interval time.Duration) func(*RoutingClient) {
	return func(rc *RoutingClient) {
		rc.retryInterval = interval
	}
}

// ------------------------------------------------------------------------------------
// Constructor
// ------------------------------------------------------------------------------------

func NewRoutingClient(ctx context.Context, logger log.ILogger, opts ...func(*RoutingClient)) IClient {
	c := &RoutingClient{
		logger:        logger,
		multicastAddr: "224.0.23.12:3671",
		retryInterval: 30 * time.Second,
		conn:          nil,
		subs:          map[async.Stream[Cemi]]bool{},
	}
	for _, o := range opts {
		o(c)
	}
	go func() {
		for {
			err := c.run(ctx)
			if err == nil {
				logger.Debugf("RoutingClient shutdown successfully - context ended")
				return // Context dead
			}
			logger.Errorf("%v", err)
			time.Sleep(c.retryInterval)
		}
	}()
	return c
}
