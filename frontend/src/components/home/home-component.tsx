import { Button } from '@/components/ui/button'
import { connectWebsocket } from '@/lib/connect-websocket'
import { useMatchStore } from '@/stores/match-store'
import { useNavigate } from '@tanstack/react-router'

const Home = () => {
  const { setWaiting } = useMatchStore()
  const navigate = useNavigate()

  const handleStartMatch = () => {
    // Set status to waiting
    setWaiting()

    // Navigate to match page
    navigate({ to: '/match' })

    // Connect to websocket
    connectWebsocket()
  }

  return (
    <div className="min-h-screen flex flex-col items-center justify-center">
      <div className="text-center space-y-8">
        <h1 className="text-6xl font-bold">
          Chaos Chess
        </h1>
        <Button
          size="lg"
          onClick={handleStartMatch}
        >
          Start Match
        </Button>
      </div>
    </div>
  )
}

export default Home