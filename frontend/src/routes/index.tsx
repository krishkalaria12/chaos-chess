import { createFileRoute } from '@tanstack/react-router'
import { Button } from '@/components/ui/button'

export const Route = createFileRoute('/')({
  component: App,
})

function App() {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center">
      <div className="text-center space-y-8">
        <h1 className="text-6xl font-bold">
          Chaos Chess
        </h1>
        <Button
          size="lg"
        >
          Start Match
        </Button>
      </div>
    </div>
  )
}
