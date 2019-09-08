const path = require('path');
var webpack = require('webpack');

module.exports = {
    entry: __dirname + '/src/index.js',
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: 'main.js'
    },
    devServer: {
        inline: true,
        contentBase: __dirname + '/dist',
        port: 9000
    },
    mode: 'development',
    watch: true,
    watchOptions: {
        aggregateTimeout: 300,
        poll: 1000,
        ignored: /node_modules/
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader'
                }
            },
            {
                test: /\.scss$/,
                use: {
                    loader: 'style-loader!css-loader!sass-loader'
                }
            }
        ]
    },
    plugins: [
        new webpack.optimize.OccurrenceOrderPlugin()
    ]
};