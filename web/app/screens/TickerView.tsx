import * as React from "react";
import { Subscription } from "rxjs";
import { sampleTime } from "rxjs/operators";

import { Loading } from "~modules/application";
import { PriceSummary, Stream, TickerWheel } from "~modules/ticker";

import "./TickerView.css";

interface State {
    instrument: string;
    summary: PriceSummary | null;
}

export class TickerView extends React.Component<{}, State> {
    state: State = {
        instrument: "BTC-USD",
        summary: null,
    };

    private stream = new Stream();
    private subscription: Subscription = new Subscription();

    componentDidMount() {
        this.stream.connect();
        this.stream.subscribe(this.state.instrument);
        this.subscription.add(
            this.stream.observe()
                .pipe(sampleTime(500))
                .subscribe((summary) => this.setState({
                    summary: summary.ask,
                }))
        );
    }

    componentWillUnmount() {
        this.stream.disconnect();
        this.subscription.unsubscribe();
    }

    render() {
        const { summary } = this.state;
        const loading = summary === null;

        if (loading) {
            return (
                <div className="TickerView is-loading">
                    <Loading />
                </div>
            );
        }

        return (
            <div className="TickerView">
                <TickerWheel {...(summary as PriceSummary)} />
            </div>
        );
    }
}
