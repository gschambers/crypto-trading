import { prepend, reverse, slice } from "ramda";
import * as React from "react";
import { Sparklines, SparklinesLine, SparklinesSpots } from "react-sparklines";
import { Subscription } from "rxjs";
import { sampleTime } from "rxjs/operators";
import { Instrument, Stream } from "~modules/ticker";
import "./TickerView.css";

interface State {
    instrument: Instrument;
    priceBuffer: number[];
}

export class TickerView extends React.Component<{}, State> {
    state: State = {
        instrument: { from: "BTC", to: "USD" },
        priceBuffer: [],
    };

    private stream = new Stream();
    private subscription: Subscription = new Subscription();

    componentDidMount() {
        this.stream.connect();
        this.stream.subscribe(this.state.instrument);
        this.subscription.add(
            this.stream.observe()
                .pipe(sampleTime(150))
                .subscribe((tick) => {
                    let priceBuffer = this.state.priceBuffer;
                    priceBuffer = prepend(tick.price, priceBuffer);
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
                    {instrument.from}/{instrument.to}
                </div>

                <div className="TickerView-body">
                    <Sparklines data={reverse(this.state.priceBuffer)} min={95} max={110} svgHeight={240}>
                        <SparklinesLine color="rgba(255, 255, 255, 0.5)" style={{ fill: "none" }} />
                        <SparklinesSpots style={{ fill: "white" }} />
                    </Sparklines>
                </div>
            </div>
        );
    }
}
