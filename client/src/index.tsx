import * as React from "react";
import * as ReactDOM from "react-dom";
import { App } from "~modules/application";
import { TickerView } from "./screens/TickerView";

ReactDOM.render(
    <App><TickerView /></App>,
    document.querySelector("#root"),
);
