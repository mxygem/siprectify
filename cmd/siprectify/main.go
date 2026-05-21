package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mxygem/siprectify/internal/sip"
)

func main() {
	var (
		sipPortFlag = flag.Int("sip-port", 5060, "SIP listening port")
		rtpBaseFlag = flag.Int("rtp-base", 5004, "Base port for RTP media (will use base and base+1)")
		debug       = flag.Bool("debug", false, "Enable debug logging")
	)
	flag.Parse()

	if *debug {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
	sipPort := *sipPortFlag
	rtpBase := *rtpBaseFlag

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

	log.Printf("siprectify starting: SIP listener on :%d, RTP base port %d", sipPort, rtpBase)

	server := sip.New(sipPort)
	if err := server.ListenAndServe(ctx, "udp", fmt.Sprintf("127.0.0.1:%d", sipPort)); err != nil {
		log.Fatalf("listen and serve: %s", err)
	}

	// TODO: Step 3 - Initialize RTP receiver
	// rtp.New(*rtpBase)

	// TODO: Step 4 - Initialize recorder
	// recorder.New()

	<-ctx.Done()
	log.Println("siprectify stopped")
}
