package main

import (
	"context"
	"os/signal"
	"syscall"

	knx "github.com/iotwin-at/iot-knx"
	"github.com/uoul/go-common/log"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	logger := log.NewConsoleLogger(log.TRACE)
	client := knx.NewRoutingClient(ctx, logger)

	sub := client.Subscribe()
	defer client.Unsubscribe(sub)

	logger.Info("Application is running...")
LP1:
	for {
		select {
		case <-ctx.Done():
			break LP1
		case msg := <-sub:
			logger.Info("------------------------ Incomming Dataframe ------------------------")
			logger.Infof("Apci: 0x%04X", msg.Result.APCI)
			logger.Infof("Source Address: %d.%d.%d", msg.Result.SourceAddr.Area, msg.Result.SourceAddr.Line, msg.Result.SourceAddr.Device)
			logger.Infof("Destination Address: %d/%d/%d", msg.Result.DestAddr.Main, msg.Result.DestAddr.Middle, msg.Result.DestAddr.Sub)
			logger.Infof("Data: %b", msg.Result.Data)
		}
	}

	<-ctx.Done()
	logger.Info("Shutting down...")
}
