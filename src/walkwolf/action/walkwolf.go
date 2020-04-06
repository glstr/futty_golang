package action

import (
	"errors"
	"log"
	"walkwolf/model"

	"github.com/urfave/cli"
)

const (
	CmdDetect = "detect"
)

var (
	ErrNotSupportCmd = errors.New("not support cmd")
)

func Walk(c *cli.Context) error {
	log.Printf("walk start")
	wolf := &WalkWolf{}
	cmd := c.String("cmd")
	url := c.String("rooturl")
	depth := c.String("depth")
	switch cmd {
	case CmdDetect:
		return wolf.Detect(url, depth)
	default:
		return ErrNotSupportCmd
	}
	return nil
}

type WalkWolf struct{}

func (w *WalkWolf) Detect(url string, depth int) error {
	parser := model.NewHTMLParser()
	target := &model.Target{
		RawUrl:  url,
		Element: "a",
		Attr:    "href",
	}
	result, err := parser.GetTargets(target)
	if err != nil {
		log.Printf("error_msg:%s", err.Error())
		return err
	}
	log.Printf("result:%v", result)
	return nil
}

func (w *WalkWolf) DetectAll(url string) error {
	return nil
}
