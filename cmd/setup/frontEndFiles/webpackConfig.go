package frontEndFiles

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateWebpackConfig(filePath string) {
  const template = `const path = require('path');
const webpack = require('webpack');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
  entry: {
    output: ['css-hot-loader/hotModuleReplacement', path.resolve(__dirname, 'src/Application.js')]
  },
  mode: 'development',
  devServer: {
    contentBase: './dist',
    hot: true
  },
  output: {
    path: path.resolve(__dirname, './dist/'),
    filename: 'bundle.js'
  },
  module: {
    rules: [{
      test: /\.css$/,
      use: ['css-hot-loader'].concat(ExtractTextPlugin.extract({
        fallback: 'style-loader',
        use: [ 'css-loader' ]
      }))
    }, {
      test: /\.scss$/,
      use: ['css-hot-loader'].concat(ExtractTextPlugin.extract({
        fallback: 'style-loader',
        use: [
          'css-hot-loader',
          'css-loader',
          'sass-loader'
        ]
      }))
    }, {
      test: /\.(png|gif|jpg|jpeg|svg|xml)$/,
      use: [ 'url-loader' ]
    }, {
      test: /\.js$/,
      exclude: /node_modules/,
      use: {
        loader: 'babel-loader',
        options: {
          presets: ['babel-preset-react']
        }
      }
    }]
  },
  plugins: [
    new CopyWebpackPlugin([{
      from: 'src/index.html',
      to: path.join(__dirname, 'dist')
    }]),
    new ExtractTextPlugin('styles.css')
  ]
};
`
  files.CreateFromTemplate(filePath, template, nil)
}
