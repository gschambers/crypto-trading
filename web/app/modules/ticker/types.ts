interface PriceSummary {
    price: number;
    volume: number;
}

export interface MarketSummary {
    market: string;
    ask: PriceSummary;
    bid: PriceSummary;
}
