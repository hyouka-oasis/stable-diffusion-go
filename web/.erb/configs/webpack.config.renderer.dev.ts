import 'webpack-dev-server';
import path from 'path';
import fs from 'fs';
import webpack from 'webpack';
import HtmlWebpackPlugin from 'html-webpack-plugin';
import chalk from 'chalk';
import { merge } from 'webpack-merge';
import { execSync, spawn } from 'child_process';
import ReactRefreshWebpackPlugin from '@pmmmwh/react-refresh-webpack-plugin';
import baseConfig from './webpack.config.base';
import webpackPaths from './webpack.paths';
import checkNodeEnv from '../scripts/check-node-env';
import webpackConfigLoader from './webpack.config.loader';

if (process.env.NODE_ENV === 'production') {
    checkNodeEnv('development');
}

const port = process.env.PORT || 2580;
const manifest = path.resolve(webpackPaths.dllPath, 'renderer.json');
const skipDLLs = module.parent?.filename.includes('webpack.config.renderer.dev.dll') || module.parent?.filename.includes('webpack.config.eslint');

/**
 * 校验webpack-dll文件是否存在
 */
if (!skipDLLs && !(fs.existsSync(webpackPaths.dllPath) && fs.existsSync(manifest))) {
    console.log(chalk.black.bgYellow.bold('缺少webpack-dll依赖正在通过 "npm run build:dll" 安装请耐心等待.'));
    execSync('npm run build:dll');
}
const configuration: webpack.Configuration = {
    devtool: 'inline-source-map',

    mode: 'development',

    target: ['web', 'electron-renderer'],

    entry: [`webpack-dev-server/client?http://localhost:${port}/dist`, 'webpack/hot/only-dev-server', path.join(webpackPaths.srcRendererPath, 'index.tsx')],

    output: {
        path: webpackPaths.distRendererPath,
        publicPath: '/',
        filename: 'renderer.dev.js',
        library: {
            type: 'umd',
        },
    },

    module: {
        rules: webpackConfigLoader.webpackDevelopLoader,
    },
    plugins: [
        ...(skipDLLs
            ? []
            : [
                  new webpack.DllReferencePlugin({
                      context: webpackPaths.dllPath,
                      manifest: require(manifest),
                      sourceType: 'var',
                  }),
              ]),

        new webpack.NoEmitOnErrorsPlugin(),

        new webpack.EnvironmentPlugin({
            NODE_ENV: 'development',
        }),

        new webpack.LoaderOptionsPlugin({
            debug: true,
        }),

        new ReactRefreshWebpackPlugin(),

        new HtmlWebpackPlugin({
            filename: path.join('index.html'),
            template: path.join(webpackPaths.srcRendererPath, 'index.ejs'),
            minify: {
                collapseWhitespace: true,
                removeAttributeQuotes: true,
                removeComments: true,
            },
            isBrowser: false,
            env: process.env.NODE_ENV,
            isDevelopment: process.env.NODE_ENV !== 'production',
            nodeModules: webpackPaths.appNodeModulesPath,
        }),
    ],

    node: {
        __dirname: false,
        __filename: false,
    },

    devServer: {
        port,
        compress: true,
        hot: true,
        headers: { 'Access-Control-Allow-Origin': '*' },
        static: {
            publicPath: '/',
        },
        historyApiFallback: {
            verbose: true,
        },
        setupMiddlewares(middlewares) {
            console.log('正在加载客户端...');
            const preloadProcess = spawn('npm', ['run', 'start:preload'], {
                shell: true,
                stdio: 'inherit',
            })
                .on('close', (code: number) => process.exit(code!))
                .on('error', spawnError => console.error(spawnError));

            console.log('正在启动客户端...');
            let args = ['run', 'start:main'];
            if (process.env.MAIN_ARGS) {
                args = args.concat(['--', ...process.env.MAIN_ARGS.matchAll(/"[^"]+"|[^\s"]+/g)].flat());
            }
            spawn('npm', args, {
                shell: true,
                stdio: 'inherit',
            })
                .on('close', (code: number) => {
                    preloadProcess.kill();
                    process.exit(code!);
                })
                .on('error', spawnError => console.error(spawnError));
            return middlewares;
        },
    },
};

export default merge(baseConfig, configuration);
