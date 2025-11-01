import { routeEvent, WebSocketEvent } from "@/events/event"

export let websocket: WebSocket;
export const connectWebsocket = () => {
    websocket = new WebSocket('ws://localhost:8080/ws')

    websocket.onopen = () => {
      console.log('Connected to server')
    }

    websocket.onmessage = (evt) => {
      console.log('Received message:', evt.data);

      const eventData = JSON.parse(evt.data)
      const event: WebSocketEvent = {
        type: eventData.type,
        payload: eventData.payload
      }
      routeEvent(event)
    }

    websocket.onerror = (error) => {
      console.error('WebSocket error:', error)
    }

    websocket.onclose = () => {
      console.log('Disconnected from server')
    }
}