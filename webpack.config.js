'use strict'

const { VueLoaderPlugin } = require('vue-loader')

module.exports = {
    entry: `${__dirname}/assets/main.js`,
    output: {
        path: `${__dirname}/public`,
        filename: 'app.js'
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loader: 'babel-loader'
            }, {
                test: /\.vue$/,
                loader: 'vue-loader'
            }, {
                test: /\.css$/,
                loader: ['style-loader', 'css-loader']
            }, {
                test: /\.s[ac]ss$/,
                loader: ['vue-style-loader', 'css-loader', 'sass-loader']
            }, {
                test: /\.(eot|woff|woff2|ttf|svg|png|jpe?g|gif|pdf)(\?[\s\S]+)?$/,
                loader: 'file-loader',
                options: { name: '[name].[ext]' }
            }
        ]
    },
    resolve: {
        alias: {
            'vue$': 'vue/dist/vue.esm.js',
            '@': `${__dirname}/assets`,
            'font-awesome': 'font-awesome/css/font-awesome.min.css'
        }
    },
    plugins: [
        new VueLoaderPlugin()
    ]
}