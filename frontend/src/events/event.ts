import { websocket } from "@/lib/connect-websocket"
import { ReceivePlayMoveEvent } from "./play-move-event"
import { useMatchStore } from "@/stores/match-store"
import { Chess } from "chess.js"

export interface EventType {
    type: string
    payload: any
}

export class WebSocketEvent {
    type: string
    payload: any

    constructor(type: string, payload: any){
        this.type = type
        this.payload = payload
    }
}

export function routeEvent(event: EventType) {
    if (event.type == undefined) {
        throw new Error("The event type is not defined")
    }

    const matchStore = useMatchStore.getState()

    switch(event.type){
        case "match_start": {
            const payload = event.payload
            console.log('Match started:', payload)
            matchStore.startMatch(payload.color, payload.orientation, payload.position)
            break
        }

        case "receive_play_move": {
            const payload = event.payload
            const moveEvent = Object.assign(
                new ReceivePlayMoveEvent(payload.from, payload.to, payload.sent),
                event.payload
            )

            console.log('Move received:', moveEvent)

            // Update the board position by making the move
            const currentPosition = matchStore.position
            const game = currentPosition === 'start' ? new Chess() : new Chess(currentPosition)

            try {
                game.move({
                    from: moveEvent.from,
                    to: moveEvent.to,
                    promotion: 'q'
                })
                matchStore.updatePosition(game.fen())
            } catch (error) {
                console.error('Invalid move received:', error)
            }
            break
        }

        case "error": {
            const payload = event.payload
            console.error('Error from server:', payload.message)
            matchStore.setError(payload.message)
            break
        }

        case "match_complete": {
            const payload = event.payload
            console.log('Match completed:', payload.message)
            matchStore.completeMatch(payload.message)
            break
        }

        default:
            console.warn("Unsupported Event type:", event.type)
    }
}

export function sendEvent(eventName: string, payload: any) {
    const event = new WebSocketEvent(eventName, payload)
    websocket.send(JSON.stringify(event))
}