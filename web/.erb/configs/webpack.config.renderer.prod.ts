/**
 * Build config for electron renderer process
 */

import path from 'path';
import webpack from 'webpack';
import HtmlWebpackPlugin from 'html-webpack-plugin';
import MiniCssExtractPlugin from 'mini-css-extract-plugin';
import { BundleAnalyzerPlugin } from 'webpack-bundle-analyzer';
import CssMinimizerPlugin from 'css-minimizer-webpack-plugin';
import { merge } from 'webpack-merge';
import TerserPlugin from 'terser-webpack-plugin';
import baseConfig from './webpack.config.base';
import webpackPaths from './webpack.paths';
import checkNodeEnv from '../scripts/check-node-env';
import webpackConfigLoader from './webpack.config.loader';

checkNodeEnv('production');

const configuration: webpack.Configuration = {
    devtool: 'source-map',

    mode: 'production',

    target: [ 'web', 'electron-renderer' ],

    entry: [ path.join(webpackPaths.srcRendererPath, 'index.tsx') ],

    output: {
        path: webpackPaths.distRendererPath,
        publicPath: './',
        filename: 'renderer.js',
        library: {
            type: 'umd',
        },
    },

    module: {
        rules: webpackConfigLoader.webpackProductionLoader,
    },

    optimization: {
        minimize: true,
        minimizer: [
            new TerserPlugin({
                parallel: true,
            }),
            new CssMinimizerPlugin(),
        ],
    },

    plugins: [
        new webpack.EnvironmentPlugin({
            NODE_ENV: 'production',
            DEBUG_PROD: false,
        }),

        new MiniCssExtractPlugin({
            filename: 'style.css',
        }),

        new BundleAnalyzerPlugin({
            analyzerMode: process.env.ANALYZE === 'true' ? 'server' : 'disabled',
            analyzerPort: 8889,
        }),

        new HtmlWebpackPlugin({
            filename: 'index.html',
            template: path.join(webpackPaths.srcRendererPath, 'index.ejs'),
            minify: {
                collapseWhitespace: true,
                removeAttributeQuotes: true,
                removeComments: true,
            },
            isBrowser: false,
            isDevelopment: process.env.NODE_ENV !== 'production',
        }),

        new webpack.DefinePlugin({
            'process.type': '"renderer"',
        }),
    ],
};

export default merge(baseConfig, configuration);
