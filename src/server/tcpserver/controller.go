package tcpserver

import "errors"

var (
	methodLogin  uint32 = 1
	methodLogout uint32 = 2
)

type Handler interface {
	HandleReq(req *MessageRequest) (*MessageResponse, error)
}

type handleFunc func(req *MessageRequest) (*MessageResponse, error)

type Controller struct {
	methods map[uint32]handleFunc
}

func NewController() *Controller {
	methods := make(map[uint32]handleFunc)
	methods[methodLogin] = Login
	methods[methodLogout] = Logout
	return &Controller{
		methods: methods,
	}
}

func (c *Controller) HandleReq(req *MessageRequest) (*MessageResponse, error) {
	if function, ok := c.methods[req.Method]; ok {
		return function(req)
	}
	return nil, errors.New("no method found")
}

func Login(req *MessageRequest) (*MessageResponse, error) {
	res := CopyFromReq(req)
	payload := "login success"
	res.SetPayload([]byte(payload))
	return res, nil
}

func Logout(req *MessageRequest) (*MessageResponse, error) {
	res := CopyFromReq(req)
	payload := "logout success"
	res.SetPayload([]byte(payload))
	return res, nil
}
