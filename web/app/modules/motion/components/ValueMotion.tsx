import * as React from "react";
import { Motion, PlainStyle, spring, Style } from "react-motion";

interface Props {
    children: (interpolatedValue: number) => React.ReactElement<any>;
    nextValue: number;
    prevValue: number;
    onRest: () => void;
}

export class ValueMotion extends React.Component<Props> {
    deriveDefaultStyle(): PlainStyle {
        return {
            value: this.props.prevValue,
        };
    }

    deriveStyle(): Style {
        return {
            value: spring(this.props.nextValue),
        };
    }

    render() {
        const { children } = this.props;

        return (
            <Motion defaultStyle={this.deriveDefaultStyle()} style={this.deriveStyle()} onRest={this.props.onRest}>
                {({ value }) => children(value)}
            </Motion>
        )
    }
}
