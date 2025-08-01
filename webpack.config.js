const path = require('path');
const webpack = require('webpack');

const ForkTsCheckerWebpackPlugin = require('fork-ts-checker-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin'); // インポート


module.exports = {
  entry: './public/src/index.tsx', // エントリポイント
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js',
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx|ts|tsx)$/,
        exclude: /node_modules/,
        use: {
          loader: 'ts-loader',
        },
      },
      {
        test: /\.css$/, // CSSファイルの処理
        use: [
          'style-loader',
          {
            loader: 'css-loader',
            options: {
              url: true,
            },
          }
        ], // ローダーの順序に注意
      },
      {
        test: /\.(png|jpg|gif|svg)$/,
        type: 'asset/resource', // Webpack 5 ではこれが推奨
        generator: {
          filename: 'images/[name].[contenthash].[ext]', // 出力形式
        },
        // use: [
        //   {
        //     loader: 'file-loader',
        //     options: {
        //       name: '[name].[contenthash].[ext]',
        //       outputPath: 'images/',
        //       publicPath: '/images/',
        //     },
        //   },
        // ],
      }
    ],
  },
  target: 'web',
  resolve: {
    extensions: ['.js', '.jsx', '.ts', '.tsx'],
  },
  devServer: {
    setupMiddlewares: (middlewares, devServer) => {
      devServer.app.use((req, res, next) => {
        console.log('Request URL:', req.url);
        next();
      });
      return middlewares;
    },
    static: {
      directory: path.join(__dirname, 'public'),
    },
    historyApiFallback: true,
    compress: true,
    port: 3000,
  },
  plugins: [
    new ForkTsCheckerWebpackPlugin(), // TypeScript型チェックを有効化
    new HtmlWebpackPlugin({
      template: './public/index.html', // テンプレートHTML
    }),
    new webpack.DefinePlugin({
      'process.env.GOOGLE_CLIENT_ID': JSON.stringify(process.env.GOOGLE_CLIENT_ID),
      'process.env.GOOGLE_CLIENT_SECRET': JSON.stringify(process.env.GOOGLE_CLIENT_SECRET),
      'process.env.REDIRET_URI': JSON.stringify(process.env.REDIRECT_URI),
      'process.env.REACT_APP_API_BASE': JSON.stringify(process.env.REACT_APP_API_BASE),
    }),
  ],
};


console.log("REDIRECT_URI = ", process.env.REDIRECT_URI);
console.log("REACT_APP_API_BASE = ", process.env.REACT_APP_API_BASE);