import { useState, useRef, useEffect } from 'react'
import './Gallery.css'

interface Photo {
  id: number
  url: string
  name: string
  uploadTime: Date
}

function Gallery() {
  const [photos, setPhotos] = useState<Photo[]>([
    {
      id: 1,
      url: 'https://picsum.photos/400/300?random=1',
      name: '示例照片 1',
      uploadTime: new Date()
    },
    {
      id: 2,
      url: 'https://picsum.photos/400/500?random=2',
      name: '示例照片 2',
      uploadTime: new Date()
    },
    {
      id: 3,
      url: 'https://picsum.photos/400/400?random=3',
      name: '示例照片 3',
      uploadTime: new Date()
    },
    {
      id: 4,
      url: 'https://picsum.photos/400/600?random=4',
      name: '示例照片 4',
      uploadTime: new Date()
    },
    {
      id: 5,
      url: 'https://picsum.photos/400/350?random=5',
      name: '示例照片 5',
      uploadTime: new Date()
    },
    {
      id: 6,
      url: 'https://picsum.photos/400/450?random=6',
      name: '示例照片 6',
      uploadTime: new Date()
    },
    {
      id: 7,
      url: 'https://picsum.photos/400/550?random=7',
      name: '示例照片 7',
      uploadTime: new Date()
    },
    {
      id: 8,
      url: 'https://picsum.photos/400/380?random=8',
      name: '示例照片 8',
      uploadTime: new Date()
    },
    {
      id: 9,
      url: 'https://picsum.photos/400/420?random=9',
      name: '示例照片 9',
      uploadTime: new Date()
    },
    {
      id: 10,
      url: 'https://picsum.photos/400/480?random=10',
      name: '示例照片 10',
      uploadTime: new Date()
    },
    {
      id: 11,
      url: 'https://picsum.photos/400/320?random=11',
      name: '示例照片 11',
      uploadTime: new Date()
    },
    {
      id: 12,
      url: 'https://picsum.photos/400/520?random=12',
      name: '示例照片 12',
      uploadTime: new Date()
    },
    {
      id: 13,
      url: 'https://picsum.photos/400/400?random=13',
      name: '示例照片 13',
      uploadTime: new Date()
    },
    {
      id: 14,
      url: 'https://picsum.photos/400/460?random=14',
      name: '示例照片 14',
      uploadTime: new Date()
    },
    {
      id: 15,
      url: 'https://picsum.photos/400/340?random=15',
      name: '示例照片 15',
      uploadTime: new Date()
    }
  ])
  const [selectedPhoto, setSelectedPhoto] = useState<Photo | null>(null)
  const [isAutoScrolling, setIsAutoScrolling] = useState(true)
  const fileInputRef = useRef<HTMLInputElement>(null)
  const scrollContainerRef = useRef<HTMLDivElement>(null)
  const scrollIntervalRef = useRef<number | null>(null)
  const itemRefs = useRef<Map<number, HTMLDivElement>>(new Map())

  const handleFileSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files
    if (!files || files.length === 0) return

    Array.from(files).forEach((file) => {
      if (file.type.startsWith('image/')) {
        const reader = new FileReader()
        reader.onload = (event) => {
          const url = event.target?.result as string
          const newPhoto: Photo = {
            id: photos.length + 1,
            url: url,
            name: file.name,
            uploadTime: new Date()
          }
          setPhotos((prev) => [...prev, newPhoto])
        }
        reader.readAsDataURL(file)
      }
    })

    // 清空 input，以便可以重复选择同一文件
    if (fileInputRef.current) {
      fileInputRef.current.value = ''
    }
  }

  const handleUploadClick = () => {
    fileInputRef.current?.click()
  }

  const handlePhotoClick = (photo: Photo) => {
    setSelectedPhoto(photo)
  }

  const handleCloseModal = () => {
    setSelectedPhoto(null)
  }

  const handleDeletePhoto = (id: number, e: React.MouseEvent) => {
    e.stopPropagation()
    setPhotos((prev) => prev.filter((photo) => photo.id !== id))
  }

  // 更新3D效果
  const update3DEffects = useRef(() => {
    if (!scrollContainerRef.current) return

    const container = scrollContainerRef.current
    const containerRect = container.getBoundingClientRect()
    const containerCenter = containerRect.left + containerRect.width / 2

    itemRefs.current.forEach((item) => {
      if (!item) return

      const itemRect = item.getBoundingClientRect()
      const itemCenter = itemRect.left + itemRect.width / 2
      const distanceFromCenter = itemCenter - containerCenter
      const maxDistance = containerRect.width / 2
      const ratio = maxDistance > 0 ? Math.min(Math.max(distanceFromCenter / maxDistance, -1), 1) : 0

      // 计算3D变换
      const rotateY = ratio * 25 // 最大旋转25度
      const translateZ = Math.abs(ratio) * -50 // Z轴位移
      const scale = 1 - Math.abs(ratio) * 0.2 // 缩放，中心最大，两侧缩小
      const opacity = 0.6 + (1 - Math.abs(ratio)) * 0.4 // 透明度

      item.style.transform = `
        perspective(1000px) 
        rotateY(${rotateY}deg) 
        translateZ(${translateZ}px) 
        scale(${scale})
      `
      item.style.opacity = opacity.toString()
    })
  })

  // 自动滚动功能
  useEffect(() => {
    if (!isAutoScrolling || !scrollContainerRef.current) return

    const scrollContainer = scrollContainerRef.current
    let scrollPosition = 0
    const scrollSpeed = 1 // 每次滚动的像素数
    const scrollDelay = 30 // 滚动间隔（毫秒）

    const scroll = () => {
      if (scrollContainer) {
        scrollPosition += scrollSpeed
        const maxScroll = scrollContainer.scrollWidth - scrollContainer.clientWidth
        
        if (scrollPosition >= maxScroll) {
          // 滚动到底部后，重置到顶部
          scrollPosition = 0
          scrollContainer.scrollTo({ left: 0, behavior: 'auto' })
        } else {
          scrollContainer.scrollTo({ left: scrollPosition, behavior: 'auto' })
        }
        update3DEffects.current()
      }
    }

    scrollIntervalRef.current = window.setInterval(scroll, scrollDelay)

    return () => {
      if (scrollIntervalRef.current) {
        clearInterval(scrollIntervalRef.current)
      }
    }
  }, [isAutoScrolling, photos])

  // 监听滚动事件更新3D效果
  useEffect(() => {
    const container = scrollContainerRef.current
    if (!container) return

    const handleScroll = () => {
      update3DEffects.current()
    }

    const handleResize = () => {
      update3DEffects.current()
    }

    container.addEventListener('scroll', handleScroll)
    window.addEventListener('resize', handleResize)
    
    // 初始更新
    update3DEffects.current()

    return () => {
      container.removeEventListener('scroll', handleScroll)
      window.removeEventListener('resize', handleResize)
    }
  }, [photos])

  // 鼠标悬停时暂停自动滚动
  const handleMouseEnter = () => {
    setIsAutoScrolling(false)
  }

  const handleMouseLeave = () => {
    setIsAutoScrolling(true)
  }

  // 手动滚动控制
  const handleScrollLeft = () => {
    if (scrollContainerRef.current) {
      scrollContainerRef.current.scrollBy({ left: -400, behavior: 'smooth' })
    }
  }

  const handleScrollRight = () => {
    if (scrollContainerRef.current) {
      scrollContainerRef.current.scrollBy({ left: 400, behavior: 'smooth' })
    }
  }

  return (
    <div className="gallery-container">
      <div className="gallery-header">
        <h1>照片墙</h1>
        <p className="gallery-subtitle">共 {photos.length} 张照片</p>
      </div>

      <div 
        className="gallery-carousel-container"
        onMouseEnter={handleMouseEnter}
        onMouseLeave={handleMouseLeave}
      >
        <button 
          className="gallery-scroll-button gallery-scroll-left"
          onClick={handleScrollLeft}
          aria-label="向左滚动"
        >
          ‹
        </button>
        <div 
          className="gallery-carousel"
          ref={scrollContainerRef}
        >
          {photos.map((photo) => (
            <div
              key={photo.id}
              ref={(el) => {
                if (el) {
                  itemRefs.current.set(photo.id, el)
                } else {
                  itemRefs.current.delete(photo.id)
                }
              }}
              className="gallery-item"
              onClick={() => handlePhotoClick(photo)}
            >
              <div className="gallery-item-overlay">
                <button
                  className="gallery-delete-button"
                  onClick={(e) => handleDeletePhoto(photo.id, e)}
                  title="删除照片"
                >
                  ×
                </button>
                <div className="gallery-item-info">
                  <span className="gallery-item-name">{photo.name}</span>
                </div>
              </div>
              <img
                src={photo.url}
                alt={photo.name}
                loading="lazy"
                className="gallery-image"
              />
            </div>
          ))}
        </div>
        <button 
          className="gallery-scroll-button gallery-scroll-right"
          onClick={handleScrollRight}
          aria-label="向右滚动"
        >
          ›
        </button>
      </div>

      <div className="gallery-footer">
        <button className="gallery-upload-button" onClick={handleUploadClick}>
          <span className="upload-icon">📷</span>
          <span>上传照片</span>
        </button>
        <input
          ref={fileInputRef}
          type="file"
          accept="image/*"
          multiple
          onChange={handleFileSelect}
          style={{ display: 'none' }}
        />
      </div>

      {selectedPhoto && (
        <div className="gallery-modal" onClick={handleCloseModal}>
          <div className="gallery-modal-content" onClick={(e) => e.stopPropagation()}>
            <button className="gallery-modal-close" onClick={handleCloseModal}>
              ×
            </button>
            <img
              src={selectedPhoto.url}
              alt={selectedPhoto.name}
              className="gallery-modal-image"
            />
            <div className="gallery-modal-info">
              <h3>{selectedPhoto.name}</h3>
              <p>
                上传时间：{selectedPhoto.uploadTime.toLocaleString('zh-CN')}
              </p>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}

export default Gallery

