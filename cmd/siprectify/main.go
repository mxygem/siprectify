package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		sipPort = flag.Int("sip-port", 5060, "SIP listening port")
		rtpBase = flag.Int("rtp-base", 5004, "Base port for RTP media (will use base and base+1)")
		debug   = flag.Bool("debug", false, "Enable debug logging")
	)
	flag.Parse()

	if *debug {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Graceful shutdown on SIGINT/SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		log.Printf("received signal %v, shutting down", sig)
		cancel()
	}()

	log.Printf("siprectify starting: SIP listener on :%d, RTP base port %d", *sipPort, *rtpBase)

	// TODO: Step 2 - Initialize SIP server
	// sip.New(*sipPort)

	// TODO: Step 3 - Initialize RTP receiver
	// rtp.New(*rtpBase)

	// TODO: Step 4 - Initialize recorder
	// recorder.New()

	<-ctx.Done()
	log.Println("siprectify stopped")
}
