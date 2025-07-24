import { BrowserRouter, Routes, Route, Link } from 'react-router-dom'
import Home from './Home'
import About from './About'
import Dashboard from './Dashboard'
import InTouch from './InTouch'

export default function App() {
  return (
    <BrowserRouter>
      <nav className="p-4 flex gap-4 bg-gray-100">
        <Link to="/">Home</Link>
        <Link to="/about">About</Link>
        <Link to="/dashboard">Dashboard</Link>
        <Link to="/in-touch">In Touch</Link>
      </nav>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/in-touch" element={<InTouch />} />
      </Routes>
    </BrowserRouter>
  )
}
