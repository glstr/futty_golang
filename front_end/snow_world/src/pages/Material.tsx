import { useState, useMemo, useEffect } from 'react'
import { apiConfig } from '../config/api'
import './Material.css'

interface MaterialItem {
  id: number
  imageUrl: string
  title: string
  description: string
  author: string
  date: string
}

function Material() {
  const [keyword, setKeyword] = useState('')
  const [materialData, setMaterialData] = useState<MaterialItem[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const fetchMaterials = async () => {
      try {
        setIsLoading(true)
        const response = await fetch(`${apiConfig.baseURL}/snow/get_material_list`)
        const data = await response.json()

        if (data.error_code === 0 && data.material_list) {
          setMaterialData(data.material_list)
        } else {
          setError(data.error_msg || '获取数据失败')
        }
      } catch (err) {
        setError('网络请求失败，请稍后重试')
        console.error('Fetch materials error:', err)
      } finally {
        setIsLoading(false)
      }
    }

    fetchMaterials()
  }, [])

  const filteredList = useMemo(() => {
    const k = keyword.trim().toLowerCase()
    if (!k) {
      return materialData
    }
    return materialData.filter((item) => {
      const text =
        item.title +
        item.description +
        item.author +
        item.date
      return text.toLowerCase().includes(k)
    })
  }, [keyword, materialData])

  if (isLoading) {
    return (
      <div className="material-container">
        <div className="material-loading">加载中...</div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="material-container">
        <div className="material-error">
          <p>{error}</p>
          <button onClick={() => window.location.reload()} className="retry-button">
            重试
          </button>
        </div>
      </div>
    )
  }

  return (
    <div className="material-container">
      <div className="material-header">
        <h1>素材展示</h1>
        <p className="material-subtitle">每个卡片包含图片、标题、描述、作者和时间</p>
      </div>

      <div className="material-search">
        <input
          type="text"
          className="material-search-input"
          placeholder="搜索素材名、描述或作者"
          value={keyword}
          onChange={(e) => setKeyword(e.target.value)}
        />
      </div>

      <div className="material-grid">
        {filteredList.map((item) => (
          <div key={item.id} className="material-card">
            <div className="material-image-wrapper">
              <img src={item.imageUrl} alt={item.title} className="material-image" loading="lazy" />
            </div>
            <div className="material-card-body">
              <h2 className="material-title">{item.title}</h2>
              <p className="material-description">{item.description}</p>
              <div className="material-meta">
                <span className="material-author">{item.author}</span>
                <span className="material-date">{item.date}</span>
              </div>
            </div>
          </div>
        ))}
        {filteredList.length === 0 && (
          <div className="material-empty">未找到匹配的素材，请尝试其他关键词</div>
        )}
      </div>
    </div>
  )
}

export default Material
