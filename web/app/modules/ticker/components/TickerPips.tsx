import { both, complement, cond, range, reverse } from "ramda";
import * as React from "react";

interface Props {
    price: number;
}

const createPipRange = (value: number) =>
    reverse(range((value - 2) * 100, (value + 2) * 100 + 1));

export class TickerPips extends React.PureComponent<Props> {
    render() {
        const { price } = this.props;

        return (
            <div className="TickerPips">
                {createPipRange(price).map(renderPip)}
            </div>
        );
    }
}

const isFull = (value: number) => value % 100 === 0;
const isTenth = both(complement(isFull), (value: number) => value % 10 === 0);
const isHundredth = both(complement(isFull), complement(isTenth));

const renderFull = (value: number) => (
    <div key={value} className="TickerScale-pip TickerScale-pip--full">
        <div className="TickerScale-pipValue">
            {(value / 100).toFixed(2)}
        </div>
    </div>
);

const renderTenth = (value: number) => (
    <div key={value} className="TickerScale-pip TickerScale-pip--tenth" />
);

const renderHundredth = (value: number) => (
    <div key={value} className="TickerScale-pip TickerScale-pip--hundredth" />
);

const renderPip = cond([
    [isFull, renderFull],
    [isTenth, renderTenth],
    [isHundredth, renderHundredth],
]);
