export interface Instrument {
    from: string;
    to: string;
}

export interface Tick {
    instrument: Instrument;
    price: number;
}
