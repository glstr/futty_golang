package tcpserver

import (
	"log"
	"net"
)

type LongConnManager struct {
	sessions map[string]*Session
}

func NewLongConnManager() *LongConnManager {
	return &LongConnManager{
		sessions: make(map[string]*Session),
	}
}

func (m *LongConnManager) HandleConn(c net.Conn) error {
	go m.handleConn(c)
	return nil
}

func (m *LongConnManager) handleConn(c net.Conn) {
	session := NewSession(c)
	m.addSession(session)
	for {
		//read request
		req, err := session.ReadRequest()
		if err != nil {
			log.Printf("session:%s, error_msg:%v", session.SessionId(), err)
			if session.Status() == StatusNotConnect {
				log.Printf("session:%s, connect closed", session.SessionId())
				m.delSession(session)
				return
			}
		}
		log.Printf("session:%s, req:%s", session.SessionId(), req.String())
		//router
		if session.Status() != StatusConnectLogin {
			res, err := m.handleRequest(req)
			if err != nil {
				log.Printf("session:%s, error_msg:%v", session.SessionId(), err)
				m.delSession(session)
				return
			}
			log.Printf("session:%s, res:%s", session.SessionId(), res.String())
			err = session.WriteResponse(res)
			if err != nil {
				log.Printf("session:%s, error_msg:%s", session.SessionId(), err)
				m.delSession(session)
				return
			}
			continue
		}
	}
}

func (m *LongConnManager) addSession(s *Session) error {
	sessionId := s.SessionId()
	m.sessions[sessionId] = s
	return nil
}

func (m *LongConnManager) delSession(s *Session) error {
	sessionId := s.SessionId()
	delete(m.sessions, sessionId)
	return nil
}

func (m *LongConnManager) handleRequest(req *MessageRequest) (*MessageResponse,
	error) {
	handler, err := GetHanderByRequest(req)
	if err != nil {
		return nil, err
	}
	return handler.HandleReq(req)
}

func GetHanderByRequest(req *MessageRequest) (Handler, error) {
	return NewController(), nil
}
