# siprectify

Minimal [SIPREC](https://datatracker.ietf.org/doc/html/rfc6341) Recording Server in [Go](https://go.dev/), built with [sipgo](https://github.com/emiago/sipgo) and [Pion RTP](https://github.com/pion/rtp), to ingest and persist call media for downstream analysis.

## What is SIPREC?

[SIPREC (Session Recording Protocol)](https://datatracker.ietf.org/doc/html/rfc6341) is a standard for real-time media recording in VoIP systems. A Session Recording Client (SRC) — typically a Session Border Controller — sends your SRS a SIP INVITE containing two SDP offers, one per leg of the original call. Your SRS accepts the call, starts receiving RTP packets on the advertised IP/ports, decodes the audio, and persists it.

Key RFCs:
- **RFC 3261**: SIP (Session Initiation Protocol) — signaling layer
- **RFC 3550**: RTP (Real-time Transport Protocol) — media layer  
- **RFC 6341**: SIPREC core — recording session conventions
- **RFC 7866**: SIPREC metadata — participant and stream identification

## Project Goals

Demonstrate a working SIPREC recorder as a resume artifact for voice/telephony infrastructure roles.

**Weekend 1**: Get something receiving and persisting basic audio (SIP + RTP + G.711 WAV output).
**Weekend 2**: Add SIPREC-specific features (metadata parsing, dual-stream handling, test client).
**Stretch**: Observability, more codecs, containerization, gRPC control plane.

## Architecture

```
siprectify/
├── README.md
├── NOTES.md                    # Learning journal & blog draft
├── cmd/
│   ├── siprectify/
│   │   └── main.go            # Server entry point
│   └── siprectify-client/      # (Step 7) Test client
├── internal/
│   ├── sip/                    # Signaling layer
│   ├── rtp/                    # Media receiver & parser
│   └── recorder/               # Audio persistence (codec decode, WAV write)
├── testdata/                   # Sample SDP, RTP captures, media files
└── docs/                       # RFC notes, architecture diagrams
```

## What's Working

- [ ] **Step 1**: Repo hygiene, README, and dependency setup *(in progress)*
- [ ] **Step 2**: SIP signaling layer (listen on UDP/5060, respond to INVITE with 200 OK)
- [ ] **Step 3**: RTP receiver (open UDP sockets, parse RTP packets, log sequence/timestamp)
- [ ] **Step 4**: Codec decode & WAV output (G.711 µ-law/A-law to 16-bit PCM)
- [ ] **Step 5**: SIPREC detection (Require: siprec header, multipart body, metadata parsing)
- [ ] **Step 6**: Dual-stream output (separate caller/callee WAV files with labels from metadata)
- [ ] **Step 7**: Test client (Go SIPREC INVITE generator for end-to-end testing)
- [ ] **Step 8**: Blog post (600–1000 words on what SIPREC is and lessons learned)

## Building

```bash
go build -o siprectify ./cmd/siprectify
```

## Testing

Requires [sipp](https://sipp.sourceforge.net/) for SIP load testing:

```bash
brew install sipp
```

```bash
# Terminal 1: Run the server
./siprectify

# Terminal 2: Send SIP INVITE via sipp
sipp -sf uac_siprec.xml 127.0.0.1
```

## Dependencies

- **sipgo** — SIP protocol handling (UAS, message parsing)
- **pion/rtp** — RTP packet marshaling/unmarshaling
- **go-audio/wav** — WAV file writing (coming in Step 4)

## References

- **Pion examples**: [github.com/pion/example-webrtc-applications](https://github.com/pion/example-webrtc-applications) (save-to-disk, rtp-to-webrtc patterns)
- **emiago/sipgo examples**: [github.com/emiago/sipgo/examples](https://github.com/emiago/sipgo/tree/main/examples)
- **Wireshark**: Capture SIP + RTP on loopback (`lo0`) with dissectors enabled — invaluable for understanding packet flow
- **RFC 3550** (RTP): ~80 pages, worth reading front-to-back for sequence numbers, timestamps, jitter
- **RFC 6341 + 7866** (SIPREC): Shorter; focus on metadata format and stream-labeling conventions

## Author Notes

Built as a voice/telephony infrastructure learning project. See [NOTES.md](NOTES.md) for a public learning journal, design decisions, and RFC sections that tripped me up.
