import { useEffect, useState } from 'react'
import './WorldMap.css'

// 使用 OpenStreetMap 作为地图数据源
function WorldMap() {
  const [mapLoaded, setMapLoaded] = useState(false)
  const [center, setCenter] = useState<[number, number]>([39.9042, 116.4074]) // 北京坐标
  const [zoom, setZoom] = useState(3)

  useEffect(() => {
    // 检查是否已经加载
    if ((window as any).L) {
      initMap()
      return
    }

    // 动态加载 Leaflet CSS
    const existingLink = document.querySelector('link[href*="leaflet"]')
    if (!existingLink) {
      const link = document.createElement('link')
      link.rel = 'stylesheet'
      link.href = 'https://unpkg.com/leaflet@1.9.4/dist/leaflet.css'
      link.integrity = 'sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY='
      link.crossOrigin = ''
      document.head.appendChild(link)
    }

    // 动态加载 Leaflet JS
    const existingScript = document.querySelector('script[src*="leaflet"]')
    if (!existingScript) {
      const script = document.createElement('script')
      script.src = 'https://unpkg.com/leaflet@1.9.4/dist/leaflet.js'
      script.integrity = 'sha256-20nQCchB9co0qIjJZRGuk2/Z9VM+kNiyxNV1lvTlZBo='
      script.crossOrigin = ''
      script.onload = () => {
        initMap()
      }
      document.body.appendChild(script)
    } else {
      // 如果脚本已存在，等待一下再初始化
      setTimeout(() => {
        if ((window as any).L) {
          initMap()
        }
      }, 100)
    }
  }, [])

  const initMap = () => {
    const L = (window as any).L
    if (!L) {
      console.error('Leaflet not loaded')
      return
    }

    // 检查地图容器是否存在
    const mapContainer = document.getElementById('world-map')
    if (!mapContainer) {
      console.error('Map container not found')
      return
    }

    // 如果地图已经初始化，先移除
    if ((mapContainer as any)._leaflet_id) {
      return
    }

    // 创建地图
    const map = L.map('world-map', {
      zoomControl: true,
      attributionControl: true
    }).setView(center, zoom)

    // 添加 OpenStreetMap 图层
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
      maxZoom: 19,
      minZoom: 2
    }).addTo(map)

    // 添加一些主要城市的标记
    const cities = [
      { name: '北京', coords: [39.9042, 116.4074] },
      { name: '上海', coords: [31.2304, 121.4737] },
      { name: '广州', coords: [23.1291, 113.2644] },
      { name: '深圳', coords: [22.5431, 114.0579] },
      { name: '纽约', coords: [40.7128, -74.0060] },
      { name: '伦敦', coords: [51.5074, -0.1278] },
      { name: '东京', coords: [35.6762, 139.6503] },
      { name: '巴黎', coords: [48.8566, 2.3522] },
      { name: '悉尼', coords: [-33.8688, 151.2093] },
      { name: '莫斯科', coords: [55.7558, 37.6173] }
    ]

    cities.forEach(city => {
      L.marker(city.coords as [number, number])
        .addTo(map)
        .bindPopup(city.name)
    })

    // 地图点击事件
    map.on('click', (e: any) => {
      const { lat, lng } = e.latlng
      L.popup()
        .setLatLng([lat, lng])
        .setContent(`坐标: ${lat.toFixed(4)}, ${lng.toFixed(4)}`)
        .openOn(map)
    })

    // 地图缩放事件
    map.on('zoomend', () => {
      setZoom(map.getZoom())
      setCenter([map.getCenter().lat, map.getCenter().lng])
    })

    setMapLoaded(true)
  }

  return (
    <div className="world-map-container">
      <div className="world-map-header">
        <h1>世界地图</h1>
        <div className="map-info">
          <span>缩放级别: {zoom}</span>
          <span>中心: {center[0].toFixed(2)}, {center[1].toFixed(2)}</span>
        </div>
      </div>
      <div id="world-map" className="world-map"></div>
      <div className="map-controls">
        <div className="map-legend">
          <h3>地图说明</h3>
          <ul>
            <li>点击地图查看坐标</li>
            <li>使用鼠标滚轮或工具栏缩放</li>
            <li>拖拽地图进行平移</li>
            <li>支持缩放至城市级别</li>
          </ul>
        </div>
      </div>
    </div>
  )
}

export default WorldMap

