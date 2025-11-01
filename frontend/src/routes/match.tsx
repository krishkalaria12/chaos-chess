import { Match } from '@/components/match/match-component'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/match')({
  component: Match,
})
