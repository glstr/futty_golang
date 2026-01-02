import { Link, useLocation } from 'react-router-dom'
import './Navbar.css'

function Navbar() {
  const location = useLocation()

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

