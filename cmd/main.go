package main

import (
	"context"
	"os/signal"
	"syscall"

	knx "github.com/iotwin-at/iot-knx"
	"github.com/iotwin-at/iot-knx/dpt"
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

			if msg.Result.DestAddr.Main == 0 && msg.Result.DestAddr.Middle == 0 && (msg.Result.DestAddr.Sub == 1 || msg.Result.DestAddr.Sub == 2) {
				value, err := dpt.UnpackDpt1002(msg.Result.Data)
				if err != nil {
					logger.Errorf("failed to parse value")
				} else {
					logger.Infof("  → DPT 1.002 (Bool): %v", value)
				}
			}

			if msg.Result.DestAddr.Main == 0 && msg.Result.DestAddr.Middle == 0 {
				// CO2
				if msg.Result.DestAddr.Sub == 10 {
					value, err := dpt.UnpackDpt9008(msg.Result.Data)
					if err != nil {
						logger.Errorf("%v", err)

					} else {
						logger.Infof("  → %s: %v", value.Name(), value)
					}
				}
				// Relative Luftfeuchtigkeit
				if msg.Result.DestAddr.Sub == 11 {
					value, err := dpt.UnpackDpt9007(msg.Result.Data)
					if err != nil {
						logger.Errorf("%v", err)

					} else {
						logger.Infof("  → %s: %v", value.Name(), value)
					}
				}
				// Temperatur
				if msg.Result.DestAddr.Sub == 12 {
					value, err := dpt.UnpackDpt9001(msg.Result.Data)
					if err != nil {
						logger.Errorf("%v", err)

					} else {
						logger.Infof("  → %s: %v", value.Name(), value)
					}
				}
				// Taupunkt
				if msg.Result.DestAddr.Sub == 13 {
					value, err := dpt.UnpackDpt9001(msg.Result.Data)
					if err != nil {
						logger.Errorf("%v", err)

					} else {
						logger.Infof("  → %s: %v", value.Name(), value)
					}
				}
				// Luftfeuchte Absolut
				if msg.Result.DestAddr.Sub == 14 {
					value, err := dpt.UnpackDpt9006(msg.Result.Data)
					if err != nil {
						logger.Errorf("%v", err)

					} else {
						logger.Infof("  → %s: %v", value.Name(), value)
					}
				}
				// Luftfeuchte Relativ
				if msg.Result.DestAddr.Sub == 15 {
					value, err := dpt.UnpackDpt9006(msg.Result.Data)
					if err != nil {
						logger.Errorf("%v", err)

					} else {
						logger.Infof("  → %s: %v", value.Name(), value)
					}
				}
			}
		}
	}

	<-ctx.Done()
	logger.Info("Shutting down...")
}
