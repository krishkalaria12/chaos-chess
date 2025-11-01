import { createFileRoute } from '@tanstack/react-router'
import Home from '@/components/home/home-component'

export const Route = createFileRoute('/')({
  component: App,
})

function App() {
  return (
    <Home />
  )
}
