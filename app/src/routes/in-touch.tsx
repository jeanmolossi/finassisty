import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/in-touch')({
  component: InTouch,
})

function InTouch() {
  return (
    <main className="p-4">
      <h1 className="text-2xl font-bold mb-4">In Touch</h1>
      <p>Fale conosco.</p>
    </main>
  )
}
