package tcpserver

import (
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type Session struct {
	conn      net.Conn
	sessionId string
	s         status
}

const (
	StatusNotConnect      = -1
	StatusConnectNotLogin = 0
	StatusConnectLogin    = 1
)

type status int32

func NewSession(c net.Conn) *Session {
	return &Session{
		sessionId: genSessionId(),
		conn:      c,
		s:         StatusConnectNotLogin,
	}
}

func genSessionId() string {
	ts := time.Now().UnixNano()
	tsStr := strconv.FormatInt(ts, 10)
	return tsStr
}

func (s *Session) SessionId() string {
	return s.sessionId
}

func (s *Session) Status() status {
	return s.s
}

func (s *Session) SetStatus(st status) error {
	s.s = st
	return nil
}

func (s *Session) ReadRequest() (*MessageRequest, error) {
	msgReq, err := ReadRequest(s.conn)
	if err != nil {
		if err == io.EOF {
			s.SetStatus(StatusNotConnect)
		}
	}
	return msgReq, err
}

func (s *Session) WriteResponse(res *MessageResponse) error {
	return WriteResponse(s.conn, res)
}

func (s *Session) Close() error {
	err := s.conn.Close()
	if err != nil {
		log.Printf("session:%s, error_msg:%v", s.SessionId(), err)
	}
	return err
}
