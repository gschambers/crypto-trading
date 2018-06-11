import * as React from "react";
import { TickerPips } from "./TickerPips";
import "./TickerScale.css";

interface Props {
    price: number;
    translateBy: number;
}

export const TickerScale: React.SFC<Props> = ({ price, translateBy }) => (
    <div className="TickerScale" style={{ transform: `translateY(${translateBy * 400}px)` }}>
        <TickerPips price={price} />
    </div>
);
