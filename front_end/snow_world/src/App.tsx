import { Routes, Route } from 'react-router-dom'
import Navbar from './components/Navbar'
import CatFollower from './components/CatFollower'
import Home from './pages/Home'
import Profile from './pages/Profile'
import Chat from './pages/Chat'
import Gallery from './pages/Gallery'
import WorldMap from './pages/WorldMap'
import About from './pages/About'
import './App.css'

function App() {
  return (
    <>
      <Navbar />
      <CatFollower />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/profile" element={<Profile />} />
        <Route path="/chat" element={<Chat />} />
        <Route path="/gallery" element={<Gallery />} />
        <Route path="/map" element={<WorldMap />} />
        <Route path="/about" element={<About />} />
      </Routes>
    </>
  )
}

export default App
