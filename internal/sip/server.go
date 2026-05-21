package sip

import (
	"log"

	"github.com/emiago/sipgo"
	"github.com/emiago/sipgo/sip"
)

type Server struct {
	*sipgo.Server
}

func New(port int) *Server {
	ua, err := sipgo.NewUA(sipgo.WithUserAgent("siprectiv"))
	if err != nil {
		log.Fatalf("creating new UA: %v", err)
	}

	srv, err := sipgo.NewServer(ua)
	if err != nil {
		log.Fatalf("creating new server: %v", err)
	}

	srv.OnInvite(onInvite)
	srv.OnAck(onAck)
	srv.OnBye(onBye)

	return &Server{srv}
}

func onInvite(req *sip.Request, tx sip.ServerTransaction) {
	log.Printf("reached oInvite")
	log.Printf("%s request from %s", req.Method, &req.Laddr)
	log.Printf("msg: %s", string(req.MessageData.Body()))

	resp := sip.NewResponseFromRequest(req, 200, "OK", req.MessageData.Body())
	resp.AppendHeader(sip.NewHeader("Content-Type", "application/sip"))

	tx.Respond(resp)
}

func onAck(req *sip.Request, tx sip.ServerTransaction) {
	log.Println("reached onAck")
}

func onBye(req *sip.Request, tx sip.ServerTransaction) {
	log.Println("reached onBye")
	tx.Respond(sip.NewResponseFromRequest(req, 200, "OK", nil))
}
