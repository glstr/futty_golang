package service

import (
	"sync"
)

type MaterialInfo struct {
	Id          int64  `json:"id"`
	ImageUrl    string `json:"imageUrl"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Date        string `json:"date"`
}

type MaterialService interface {
	GetMaterialList() ([]*MaterialInfo, error)
}

var (
	materialServiceOnce sync.Once
	materialService     MaterialService
)

func GetMaterialService() MaterialService {
	materialServiceOnce.Do(func() {
		materialService = &DefaultMaterialService{}
	})
	return materialService
}

type DefaultMaterialService struct{}

func (m *DefaultMaterialService) GetMaterialList() ([]*MaterialInfo, error) {
	// Mock data
	materials := []*MaterialInfo{
		{
			Id:          1,
			ImageUrl:    "https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?auto=format&fit=crop&w=800&q=80",
			Title:       "抽象流体背景",
			Description: "充满动感的蓝色流体渐变背景，适用于科技感设计。",
			Author:      "DesignMaster",
			Date:        "2025-01-01",
		},
		{
			Id:          2,
			ImageUrl:    "https://images.unsplash.com/photo-1558655146-d09347e92766?auto=format&fit=crop&w=800&q=80",
			Title:       "极简办公桌面",
			Description: "干净整洁的白色桌面摄影，包含笔记本和咖啡。",
			Author:      "StudioLens",
			Date:        "2025-01-02",
		},
		{
			Id:          3,
			ImageUrl:    "https://images.unsplash.com/photo-1493246507139-91e8fad9978e?auto=format&fit=crop&w=800&q=80",
			Title:       "高山云海",
			Description: "壮丽的雪山与云海交织，适合旅游类项目使用。",
			Author:      "NatureWalker",
			Date:        "2025-01-03",
		},
		{
			Id:          4,
			ImageUrl:    "https://images.unsplash.com/photo-1550745165-9bc0b252726f?auto=format&fit=crop&w=800&q=80",
			Title:       "复古胶片街景",
			Description: "充满怀旧氛围的城市街头摄影，胶片质感。",
			Author:      "FilmSoul",
			Date:        "2025-01-04",
		},
		{
			Id:          5,
			ImageUrl:    "https://images.unsplash.com/photo-1507238691740-187a5b1d37b8?auto=format&fit=crop&w=800&q=80",
			Title:       "现代建筑几何",
			Description: "玻璃幕墙与钢结构的几何美学，建筑设计素材。",
			Author:      "CityScaper",
			Date:        "2025-01-05",
		},
		{
			Id:          6,
			ImageUrl:    "https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=800&q=80",
			Title:       "人像摄影精选",
			Description: "高品质人像特写，光影处理细腻自然。",
			Author:      "PortraitPro",
			Date:        "2025-01-06",
		},
		{
			Id:          7,
			ImageUrl:    "https://images.unsplash.com/photo-1621839673705-6617adf9e890?auto=format&fit=crop&w=800&q=80",
			Title:       "3D 渲染图标",
			Description: "一套色彩鲜艳的 3D 立体图标素材，适用于 App 设计。",
			Author:      "IconLab",
			Date:        "2025-01-07",
		},
		{
			Id:          8,
			ImageUrl:    "https://images.unsplash.com/photo-1470252649378-9c29740c9fa8?auto=format&fit=crop&w=800&q=80",
			Title:       "晨雾森林",
			Description: "清晨阳光穿透迷雾的森林景象，唯美自然素材。",
			Author:      "ForestSeeker",
			Date:        "2025-01-08",
		},
		{
			Id:          9,
			ImageUrl:    "https://images.unsplash.com/photo-1550684848-fac1c5b4e853?auto=format&fit=crop&w=800&q=80",
			Title:       "赛博朋克街道",
			Description: "霓虹灯闪烁的夜晚街道，充满未来科技感。",
			Author:      "CyberVision",
			Date:        "2025-01-09",
		},
		{
			Id:          10,
			ImageUrl:    "https://images.unsplash.com/photo-1519681393784-d120267933ba?auto=format&fit=crop&w=800&q=80",
			Title:       "星空延时摄影",
			Description: "璀璨银河与繁星点点的夜空素材。",
			Author:      "StarGazer",
			Date:        "2025-01-10",
		},
		{
			Id:          11,
			ImageUrl:    "https://images.unsplash.com/photo-1497215728101-856f4ea42174?auto=format&fit=crop&w=800&q=80",
			Title:       "极简主义办公空间",
			Description: "宽敞明亮的现代办公室室内设计摄影。",
			Author:      "OfficeSpace",
			Date:        "2025-01-11",
		},
		{
			Id:          12,
			ImageUrl:    "https://images.unsplash.com/photo-1516035069371-29a1b244cc32?auto=format&fit=crop&w=800&q=80",
			Title:       "复古相机特写",
			Description: "经典机械相机的微距摄影，展现精密工艺。",
			Author:      "VintageTech",
			Date:        "2025-01-12",
		},
	}
	return materials, nil
}
