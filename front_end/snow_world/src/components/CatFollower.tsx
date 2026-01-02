import { useState, useEffect, useRef } from 'react'
import './CatFollower.css'

// 二次元小乌龟图片路径
// 图片已保存在 public/images/cute_turtle.svg
const catImages = [
  '/images/cute_turtle.svg'
]

function CatFollower() {
  const [lookDirection, setLookDirection] = useState<'left' | 'right' | 'center'>('center')
  const [isClicking, setIsClicking] = useState(false)
  const [isHovering, setIsHovering] = useState(false)
  const [isYawning, setIsYawning] = useState(false)
  const [headTilt, setHeadTilt] = useState<'left' | 'right' | 'center'>('center')
  const [currentImageIndex, setCurrentImageIndex] = useState(0)
  const [imageError, setImageError] = useState(false)
  const catRef = useRef<HTMLDivElement>(null)
  const lastMouseX = useRef<number>(0)

  useEffect(() => {
    const handleMouseMove = (e: MouseEvent) => {
      if (!catRef.current) return

      const catRect = catRef.current.getBoundingClientRect()
      const catCenterX = catRect.left + catRect.width / 2
      
      // 判断鼠标在猫的左边还是右边
      if (e.clientX < catCenterX - 50) {
        setLookDirection('left')
      } else if (e.clientX > catCenterX + 50) {
        setLookDirection('right')
      } else {
        setLookDirection('center')
      }

      lastMouseX.current = e.clientX
    }

    const handleMouseClick = () => {
      setIsClicking(true)
      setTimeout(() => setIsClicking(false), 300)
    }

    const handleMouseEnter = () => {
      setIsHovering(true)
    }

    const handleMouseLeave = () => {
      setIsHovering(false)
      setLookDirection('center')
    }

    window.addEventListener('mousemove', handleMouseMove)
    window.addEventListener('click', handleMouseClick)

    return () => {
      window.removeEventListener('mousemove', handleMouseMove)
      window.removeEventListener('click', handleMouseClick)
    }
  }, [])

  // 自动打哈欠
  useEffect(() => {
    const yawnInterval = setInterval(() => {
      setIsYawning(true)
      setTimeout(() => setIsYawning(false), 2000)
    }, 8000 + Math.random() * 5000) // 8-13秒随机打哈欠

    return () => clearInterval(yawnInterval)
  }, [])

  // 自动摆头
  useEffect(() => {
    const headTiltInterval = setInterval(() => {
      const tilts: ('left' | 'right' | 'center')[] = ['left', 'right', 'center']
      const randomTilt = tilts[Math.floor(Math.random() * tilts.length)]
      setHeadTilt(randomTilt)
      
      setTimeout(() => {
        setHeadTilt('center')
      }, 2000)
    }, 5000 + Math.random() * 3000) // 5-8秒随机摆头

    return () => clearInterval(headTiltInterval)
  }, [])

  const handleImageError = () => {
    if (currentImageIndex < catImages.length - 1) {
      setCurrentImageIndex(currentImageIndex + 1)
    } else {
      setImageError(true)
    }
  }

  return (
    <div
      ref={catRef}
      className={`cat-follower ${lookDirection} ${isClicking ? 'clicking' : ''} ${isHovering ? 'hovering' : ''} ${isYawning ? 'yawning' : ''} head-tilt-${headTilt}`}
    >
      {!imageError ? (
        <img
          src={catImages[currentImageIndex]}
          alt="可爱的二次元小乌龟"
          className="cat-image"
          onError={handleImageError}
        />
      ) : (
        // 如果所有图片都加载失败，显示CSS绘制的小乌龟作为fallback
        <div className="cat-body">
          <div className="cat-head">
            <div className="cat-ear cat-ear-left"></div>
            <div className="cat-ear cat-ear-right"></div>
            <div className="cat-face">
              <div className="cat-eye cat-eye-left">
                <div className="cat-pupil"></div>
              </div>
              <div className="cat-eye cat-eye-right">
                <div className="cat-pupil"></div>
              </div>
              <div className="cat-nose"></div>
              <div className="cat-mouth">
                {isYawning && <div className="cat-yawn"></div>}
              </div>
            </div>
          </div>
          <div className="cat-body-main">
            <div className="cat-fur-pattern"></div>
            <div className="cat-paw cat-paw-left"></div>
            <div className="cat-paw cat-paw-right"></div>
          </div>
          <div className="cat-tail">
            <div className="cat-tail-fur"></div>
          </div>
        </div>
      )}
    </div>
  )
}

export default CatFollower

