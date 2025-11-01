import { sendEvent } from "./event"

export interface SendPlayMoveType {
    from: string
    to: string
}

export class SendPlayMoveEvent {
    from: string
    to: string

    constructor(from: string, to: string) {
        this.from = from
        this.to = to
    }
}

export class ReceivePlayMoveEvent {
    from: string
    to: string
    sent: string

    constructor(from: string, to: string, sent: string) {
        this.from = from
        this.to = to
        this.sent = sent
    }
}

export function sendPlayMove({ from, to }: SendPlayMoveType) {
    const outgoingEvent = new SendPlayMoveEvent(from, to)
    sendEvent("send_play_move", outgoingEvent)
}