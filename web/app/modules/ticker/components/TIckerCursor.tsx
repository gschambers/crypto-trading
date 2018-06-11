import * as React from "react";
import "./TickerCursor.css";

interface Props {
    price: number;
    volume: number;
}

export const TickerCursor: React.SFC<Props> = ({ price, volume }) => (
    <div className="TickerCursor">
        <div className="TickerCursor-price">
            {price.toFixed(2)}
        </div>

        <div className="TickerCursor-volume">
            {volume}
        </div>
    </div>
);
