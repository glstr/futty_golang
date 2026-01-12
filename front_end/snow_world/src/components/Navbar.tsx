import { Link, useLocation } from 'react-router-dom'
import { useState, useRef, useEffect } from 'react'
import './Navbar.css'

function Navbar() {
  const location = useLocation()
  const [isToolsOpen, setIsToolsOpen] = useState(false)
  const dropdownRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setIsToolsOpen(false)
      }
    }

    document.addEventListener('mousedown', handleClickOutside)
    return () => {
      document.removeEventListener('mousedown', handleClickOutside)
    }
  }, [])

  const isToolsActive = location.pathname.startsWith('/tools')

  return (
    <nav className="navbar">
      <div className="navbar-container">
        <div className="navbar-logo">
          <Link to="/">我的主页</Link>
        </div>
        <ul className="navbar-menu">
          <li className="navbar-item">
            <Link 
              to="/" 
              className={location.pathname === '/' ? 'navbar-link active' : 'navbar-link'}
            >
              首页
            </Link>
          </li>
          <li className="navbar-item">
            <Link 
              to="/profile" 
              className={location.pathname === '/profile' ? 'navbar-link active' : 'navbar-link'}
            >
              个人主页
            </Link>
          </li>
          <li className="navbar-item">
            <Link 
              to="/chat" 
              className={location.pathname === '/chat' ? 'navbar-link active' : 'navbar-link'}
            >
              聊天
            </Link>
          </li>
          <li className="navbar-item">
            <Link 
              to="/gallery" 
              className={location.pathname === '/gallery' ? 'navbar-link active' : 'navbar-link'}
            >
              照片墙
            </Link>
          </li>
          <li className="navbar-item">
            <Link 
              to="/map" 
              className={location.pathname === '/map' ? 'navbar-link active' : 'navbar-link'}
            >
              世界地图
            </Link>
          </li>
          <li className="navbar-item navbar-dropdown" ref={dropdownRef}>
            <button
              className={`navbar-link navbar-dropdown-toggle ${isToolsActive ? 'active' : ''}`}
              onClick={() => setIsToolsOpen(!isToolsOpen)}
            >
              工具
              <span className="dropdown-arrow">▼</span>
            </button>
            {isToolsOpen && (
              <ul className="navbar-dropdown-menu">
                <li>
                  <Link 
                    to="/tools/trace-router" 
                    className={location.pathname === '/tools/trace-router' ? 'dropdown-link active' : 'dropdown-link'}
                    onClick={() => setIsToolsOpen(false)}
                  >
                    路由追踪
                  </Link>
                </li>
              </ul>
            )}
          </li>
          <li className="navbar-item">
            <Link 
              to="/about" 
              className={location.pathname === '/about' ? 'navbar-link active' : 'navbar-link'}
            >
              关于
            </Link>
          </li>
        </ul>
      </div>
    </nav>
  )
}

export default Navbar

