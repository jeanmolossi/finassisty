import { Outlet, Link, createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
  component: Root,
})

function Root() {
  return (
    <>
      <nav className="p-4 flex gap-4 bg-gray-100">
        <Link to="/">Home</Link>
        <Link to="/about">About</Link>
        <Link to="/dashboard">Dashboard</Link>
        <Link to="/in-touch">In Touch</Link>
      </nav>
      <Outlet />
    </>
  )
}
