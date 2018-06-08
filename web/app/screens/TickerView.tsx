import { prepend, reverse, slice } from "ramda";
import * as React from "react";
import { Sparklines, SparklinesLine, SparklinesSpots } from "react-sparklines";
import { Subscription } from "rxjs";
import { Stream } from "~modules/ticker";
import "./TickerView.css";

interface State {
    instrument: string;
    priceBuffer: number[];
}

export class TickerView extends React.Component<{}, State> {
    state: State = {
        instrument: "BTC-USD",
        priceBuffer: [],
    };

    private stream = new Stream();
    private subscription: Subscription = new Subscription();

    componentDidMount() {
        this.stream.connect();
        this.stream.subscribe(this.state.instrument);
        this.subscription.add(
            this.stream.observe()
                .subscribe((summary) => {
                    let priceBuffer = this.state.priceBuffer;
                    priceBuffer = prepend(summary.ask.price, priceBuffer);
                    priceBuffer = slice(0, 25, priceBuffer);
                    this.setState({ priceBuffer });
                })
        );
    }

    componentWillUnmount() {
        this.stream.disconnect();
        this.subscription.unsubscribe();
    }

    render() {
        const { instrument } = this.state;

        return (
            <div className="TickerView">
                <div className="TickerView-title">
                    {instrument}
                </div>

                <div className="TickerView-body">
                    <Sparklines data={reverse(this.state.priceBuffer)} svgHeight={240}>
                        <SparklinesLine color="rgba(255, 255, 255, 0.5)" style={{ fill: "none" }} />
                        <SparklinesSpots style={{ fill: "white" }} />
                    </Sparklines>
                </div>
            </div>
        );
    }
}
