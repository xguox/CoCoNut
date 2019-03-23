const path              = require('path');
const webpack           = require('webpack');

let WEBPACK_ENV = process.env.WEBPACK_ENV || 'dev';
console.log(WEBPACK_ENV); 
module.exports = {
    entry: './src/app.jsx',
    output: {
        path: path.resolve(__dirname, 'dist'),
        publicPath: WEBPACK_ENV === 'dev' 
            ? '/dist/' : '//s.jianliwu.com/admin-v2-fe/dist/',
        filename: 'js/app.js'
    },
　　 // 别名设置，项目中引入的话 需要找到项目的文件，有时候深的话不好找，而且文件目录变化的话 引入路径就也得变化，利用这个设置别名
    resolve: {
        alias : {
            page        : path.resolve(__dirname, 'src/page'),
            component   : path.resolve(__dirname, 'src/component'),
            // util        : path.resolve(__dirname, 'src/util'),
            // service     : path.resolve(__dirname, 'src/service')
        }
    },
    module: {
        rules: [
          {
            test: /\.js$/,
            exclude: /node_modules/,
            use: [
              {
                loader: 'babel-loader',
                options: {
                  presets: ['es2015', 'react']
                }
              }
            ]
          }, {
            test: /\.css$/,
            use: [
              {
                loader: 'style-loader',
              }, {
                loader: 'css-loader',
              }
            ]
          }
        ],
        noParse: /node_modules\/reactstrap-tether\/dist\/js\/tether.js/,
    }
};