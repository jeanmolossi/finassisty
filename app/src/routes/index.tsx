import { Switch } from '@base-ui-components/react/switch'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
  component: Home,
})

function Home() {
  return (
    <main className="p-4">
      <h1 className="text-2xl font-bold mb-4">Finassisty Home</h1>
      <Switch.Root defaultChecked>
        <Switch.Thumb />
      </Switch.Root>
    </main>
  )
}
