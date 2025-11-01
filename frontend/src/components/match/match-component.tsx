import { useMatchStore } from '@/stores/match-store'
import { ChessBoard } from './chess-board'
import { Button } from '@/components/ui/button'
import { useNavigate } from '@tanstack/react-router'

export const Match = () => {
  const { status, playerColor, completionMessage, errorMessage, reset } = useMatchStore()
  const navigate = useNavigate()

  const handleBackHome = () => {
    reset()
    navigate({ to: '/' })
  }

  if (status === 'idle') {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center">
          <p className="text-xl mb-4">No active match</p>
          <Button onClick={handleBackHome}>Go Home</Button>
        </div>
      </div>
    )
  }

  if (status === 'waiting') {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center space-y-4">
          <div className="animate-spin rounded-full h-16 w-16 border-b-2 border-primary mx-auto"></div>
          <p className="text-2xl font-semibold">Waiting for opponent...</p>
          <p className="text-muted-foreground">Finding you a match</p>
        </div>
      </div>
    )
  }

  if (status === 'completed') {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center space-y-6">
          <h2 className="text-4xl font-bold">Game Over</h2>
          <p className="text-2xl">{completionMessage}</p>
          <Button onClick={handleBackHome}>Play Again</Button>
        </div>
      </div>
    )
  }

  // status === 'playing'
  return (
    <div className="min-h-screen flex flex-col items-center justify-center p-4">
      <div className="w-full max-w-4xl space-y-6">
        <div className="text-center space-y-2">
          <h2 className="text-3xl font-bold">Chaos Chess</h2>
          <div className="flex items-center justify-center gap-4">
            <div className={`px-4 py-2 rounded-lg ${
              playerColor === 'white' ? 'bg-white text-black' : 'bg-black text-white'
            }`}>
              You are playing as {playerColor}
            </div>
          </div>
        </div>

        {errorMessage && (
          <div className="bg-destructive/10 border border-destructive text-destructive px-4 py-3 rounded-lg text-center">
            {errorMessage}
          </div>
        )}

        <div className="flex justify-center">
          <ChessBoard />
        </div>

        <div className="text-center">
          <Button variant="outline" onClick={handleBackHome}>
            Leave Match
          </Button>
        </div>
      </div>
    </div>
  )
}