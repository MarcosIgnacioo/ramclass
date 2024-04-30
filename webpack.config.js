module.exports = {
  module: {
    rules: [
      {
        test: /\.(png|jpg|gif|ttf|otf)$/i,
        use: [
          {
            loader: 'url-loader',
            options: {
              limit: 8192,
            }
          },
        ],

        type: 'javascript/auto'
      },
    ]
  },
}
