import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/about')({
  component: About,
})

function About() {
  return (
    <main className="p-4">
      <h1 className="text-2xl font-bold mb-4">About</h1>
      <p>Sobre o Finassisty.</p>
    </main>
  )
}
