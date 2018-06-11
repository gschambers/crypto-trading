import * as React from "react";

import { ValueMotion } from "~modules/motion";
import { TickerCursor } from "./TickerCursor";
import { TickerScale } from "./TickerScale";
import { PriceSummary } from "../types";
import "./TickerWheel.css";

type Props = PriceSummary;

enum Motion {
    IDLE,
    ACTIVE,
}

interface State {
    motion: Motion;
    prevPrice: number;
}

export class TickerWheel extends React.Component<Props, State> {
    state: State = {
        motion: Motion.IDLE,
        prevPrice: this.props.price,
    };

    static getDerivedStateFromProps(props: Props, state: State): Partial<State> | null {
        if (props.price === state.prevPrice) {
            return null;
        }

        return {
            motion: Motion.ACTIVE,
            prevPrice: props.price,
        };
    }

    handleRest = () => {
        this.setState({
            motion: Motion.IDLE,
        });
    };

    render() {
        const { price, volume } = this.props;
        const { prevPrice } = this.state;

        return (
            <div className="TickerWheel">
                <ValueMotion prevValue={prevPrice} nextValue={price} onRest={this.handleRest}>
                    {(interpolatedValue) => (
                        <React.Fragment>
                            <TickerScale
                                price={Motion.ACTIVE ? prevPrice : price}
                                translateBy={interpolatedValue - prevPrice}
                            />

                            <TickerCursor
                                price={interpolatedValue}
                                volume={volume}
                            />
                        </React.Fragment>
                    )}
                </ValueMotion>
            </div>
        );
    }
}
