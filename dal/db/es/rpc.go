package es

import "errors"

var (
	ErrESNotInitialized = errors.New("es client not initialized")
)

// Init 初始化 ES Client
func Init(url string) error {
	return InitESClient(&ESConfig{URL: url})
}

// Write 写入文档到指定索引（自动生成ID）
func Write(index string, doc interface{}) error {
	cli := GetClient()
	if cli == nil {
		return ErrESNotInitialized
	}
	return cli.WriteDoc(index, doc)
}

// Read 根据ID读取文档
func Read(index string, id string, out interface{}) error {
	cli := GetClient()
	if cli == nil {
		return ErrESNotInitialized
	}
	return cli.GetDoc(index, id, out)
}

// Search 搜索指定索引，query 为 ES 查询体（map 或结构体）
// out 需为指向切片的指针，用于接收 _source 列表
func Search(index string, query interface{}, out interface{}) error {
	cli := GetClient()
	if cli == nil {
		return ErrESNotInitialized
	}
	return cli.Search(index, query, out)
}

