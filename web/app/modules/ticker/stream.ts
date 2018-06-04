import { Subject } from "rxjs";
import { Instrument, Tick } from "./types";

const STREAM_URL = `ws://${location.host}/stream`;

interface Message {
    action: "subscribe" | "unsubscribe";
    instrument: Instrument;
}

export class Stream {
    private subject = new Subject<Tick>();
    private socket: WebSocket | null = null;
    private messageBuffer: Message[] = [];

    handleOpen = () => {
        this.flushMessageBuffer();

        if (this.socket !== null) {
            this.socket.onopen = null;
        }
    };

    handleMessage = (evt: MessageEvent) => {
        const tick: Tick = JSON.parse(evt.data);
        this.subject.next(tick);
    };

    connect() {
        this.socket = new WebSocket(STREAM_URL);
        this.socket.onopen = this.handleOpen;
        this.socket.onmessage = this.handleMessage;
    }

    disconnect() {
        if (this.socket !== null) {
            this.socket.onmessage = null;
            this.socket.close();
        }
    }

    observe() {
        return this.subject.asObservable();
    }

    subscribe(instrument: Instrument) {
        this.send({
            action: "subscribe",
            instrument,
        });
    }

    unsubscribe(instrument: Instrument) {
        this.send({
            action: "unsubscribe",
            instrument,
        });
    }

    private flushMessageBuffer() {
        for (const message of this.messageBuffer) {
            this.send(message);
        }
    }

    private send(message: Message) {
        if (this.socket === null || this.socket.readyState !== WebSocket.OPEN) {
            this.messageBuffer.push(message);
            return;
        }

        this.socket.send(
            JSON.stringify(message)
        );
    }
}
