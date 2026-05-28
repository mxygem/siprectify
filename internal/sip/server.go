package sip

import (
	"log"
	"log/slog"

	"github.com/emiago/sipgo"
	"github.com/emiago/sipgo/sip"
)

type Server struct {
	*sipgo.Server
}

func New(sipPort, rtpBasePort int) *Server {
	ua, err := sipgo.NewUA(sipgo.WithUserAgent("siprectiv"))
	if err != nil {
		log.Fatalf("creating new UA: %v", err)
	}

	srv, err := sipgo.NewServer(ua)
	if err != nil {
		log.Fatalf("creating new server: %v", err)
	}

	s := &Server{srv}

	srv.OnInvite(s.onInvite)
	srv.OnAck(s.onAck)
	srv.OnBye(s.onBye)

	return s
}

func (s *Server) onInvite(req *sip.Request, tx sip.ServerTransaction) {
	slog.Info("reached onInvite", "method", req.Method, "address", &req.Laddr, "body", string(req.MessageData.Body()))

	resp := sip.NewResponseFromRequest(req, 200, "OK", req.MessageData.Body())
	resp.AppendHeader(sip.NewHeader("Content-Type", "application/sdp"))

	if err := tx.Respond(resp); err != nil {
		slog.Error("responding to invite", "error", err)
	}
}

func (s *Server) onAck(req *sip.Request, tx sip.ServerTransaction) {
	slog.Info("reached onAck", "method", req.Method, "address", &req.Laddr, "body", string(req.MessageData.Body()))
}

func (s *Server) onBye(req *sip.Request, tx sip.ServerTransaction) {
	slog.Info("reached onBye", "method", req.Method, "address", &req.Laddr, "body", string(req.MessageData.Body()))

	if err := tx.Respond(sip.NewResponseFromRequest(req, 200, "OK", nil)); err != nil {
		slog.Error("responding to bye", "error", err)
	}
}
