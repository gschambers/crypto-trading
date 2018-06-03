const path = require("path");
const mode = "development";

module.exports = {
    mode,

    entry: "./src/index.tsx",

    output: {
        filename: "index.js",
        path: path.resolve(__dirname, "../static"),
        publicPath: "/",
    },

    resolve: {
        alias: {
            "~modules": path.resolve(__dirname, "src/modules"),
        },

        extensions: [".ts", ".tsx", ".js", ".json"],
    },

    module: {
        rules: [
            { test: /\.tsx?$/, exclude: /node_modules/, use: "ts-loader" },
            { test: /\.css$/, exclude: /node_modules/, use: ["style-loader", "css-loader"] },
        ],
    },

    devServer: {
        proxy: {
            "/stream": {
                target: "ws://localhost:3000",
                ws: true,
            },
        },
    },
};
