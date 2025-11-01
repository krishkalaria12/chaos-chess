import { create } from 'zustand'

export type MatchStatus = 'idle' | 'waiting' | 'playing' | 'completed'

interface MatchState {
  status: MatchStatus
  playerColor: string | null
  boardOrientation: 'white' | 'black'
  position: string
  completionMessage: string | null
  errorMessage: string | null

  // Actions
  setWaiting: () => void
  startMatch: (color: string, orientation: 'white' | 'black', position: string) => void
  updatePosition: (position: string) => void
  completeMatch: (message: string) => void
  setError: (message: string) => void
  clearError: () => void
  reset: () => void
}

export const useMatchStore = create<MatchState>((set) => ({
  status: 'idle',
  playerColor: null,
  boardOrientation: 'white',
  position: 'start',
  completionMessage: null,
  errorMessage: null,

  setWaiting: () => set({ status: 'waiting' }),

  startMatch: (color, orientation, position) => set({
    status: 'playing',
    playerColor: color,
    boardOrientation: orientation,
    position,
  }),

  updatePosition: (position) => set({ position }),

  completeMatch: (message) => set({
    status: 'completed',
    completionMessage: message,
  }),

  setError: (message) => set({ errorMessage: message }),

  clearError: () => set({ errorMessage: null }),

  reset: () => set({
    status: 'idle',
    playerColor: null,
    boardOrientation: 'white',
    position: 'start',
    completionMessage: null,
    errorMessage: null,
  }),
}))
