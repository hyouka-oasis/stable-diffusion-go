/**
 * webpack默认配置，无特殊必要请修改以下文件
 * webpack.config.renderer.dev
 * webpack.config.renderer.prod
 */
import webpack from 'webpack';
import TsconfigPathsPlugins from 'tsconfig-paths-webpack-plugin';
import webpackPaths from './webpack.paths';
import { dependencies as externals } from '../../release/app/package.json';
import webpackConfigLoader from './webpack.config.loader';

const configuration: webpack.Configuration = {
    /**
     * 指定不应由webpack解析的依赖项
     */
    externals: [...Object.keys(externals || {})],

    stats: 'errors-only',

    module: {
        rules: webpackConfigLoader.webpackTransformationLoader,
    },

    output: {
        path: webpackPaths.srcPath,
        library: {
            type: 'commonjs2',
        },
    },

    resolve: {
        extensions: ['.js', '.jsx', '.json', '.ts', '.tsx'],
        modules: [webpackPaths.srcPath, 'node_modules'],
        plugins: [new TsconfigPathsPlugins({})],
        alias: {
            renderer: webpackPaths.srcRendererPath,
            main: webpackPaths.srcMainPath,
            declaration: webpackPaths.srcDeclarationPath,
        },
    },

    plugins: [
        new webpack.EnvironmentPlugin({
            NODE_ENV: 'production',
        }),
    ],
};

export default configuration;
