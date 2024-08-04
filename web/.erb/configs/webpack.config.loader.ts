/**
 * webpack默认基础的loader配置
 * TODO 特别注意这边是基础的loader，如果对应开发需要添加loader请到对应文件当中去修改
 */
import webpack from "webpack";
import MiniCssExtractPlugin from "mini-css-extract-plugin";

const webpackTransformationLoader = [
    {
        test: /\.[jt]sx?$/,
        exclude: /node_modules/,
        use: {
            loader: 'ts-loader',
            options: {
                transpileOnly: true,
                compilerOptions: {
                    module: 'esnext',
                },
            },
        },
    },
];

const webpackDefaultLoader: webpack.RuleSetRule[] = [
    {
        test: /\.(woff|woff2|eot|ttf|otf)$/i,
        type: 'asset/resource',
    },
    {
        test: /\.(png|jpg|jpeg|gif)$/i,
        type: 'asset/resource',
    },
    {
        test: /\.(wav|mp3|mp4|ogg|webm)$/,
        use: [
            {
                loader: "file-loader",
                options: {
                    name: "assets/file/[name].[hash:16].[ext]"
                }
            }
        ],
    },
    {
        test: /\.svg$/,
        use: [
            {
                loader: '@svgr/webpack',
                options: {
                    prettier: false,
                    svgo: false,
                    svgoConfig: {
                        plugins: [ { removeViewBox: false } ],
                    },
                    titleProp: true,
                    ref: true,
                },
            },
        ],
    },
];

const webpackDevelopLoader = [
    {
        test: /\.l?(c|e)ss$/,
        use: [
            'style-loader',
            'css-loader',
            {
                loader: 'less-loader',
                options: {
                    lessOptions: {
                        javascriptEnabled: true,
                    },
                },
            },
        ],
    },
    ...webpackDefaultLoader
];

const webpackProductionLoader = [
    {
        test: /\.l?(c|e)ss$/,
        use: [
            MiniCssExtractPlugin.loader,
            'css-loader',
            {
                loader: 'less-loader',
                options: {
                    lessOptions: {
                        javascriptEnabled: true,
                    },
                },
            },
        ],
    },
    ...webpackDefaultLoader
];

export default {
    webpackDevelopLoader,
    webpackTransformationLoader,
    webpackProductionLoader
};
