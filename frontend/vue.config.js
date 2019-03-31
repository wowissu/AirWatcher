const baseUrl = process.env.NODE_ENV === 'production' ? '/' : '/'

const host = '0.0.0.0'
const port = 8080

module.exports = {
  // base url
  publicPath: baseUrl,
  devServer: {
      port,
      host,
      https: false,
      // hotOnly: true,
      disableHostCheck: true,
      // clientLogLevel: 'warning',
      // inline: false,
      // headers: {
      //   'Access-Control-Allow-Origin': 'http://localhost',
      //   'Access-Control-Allow-Credentials': false,
      //   'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, PATCH, OPTIONS',
      //   'Access-Control-Allow-Headers': 'X-Requested-With, content-type, Authorization'
      // },
      proxy: {
        '/api': {
          target: 'http://127.0.0.1:6543',
          ws: true,
          changeOrigin: true
        }
      }
  },

  css: {
    loaderOptions: {
      sass: {
        data: `
          $baseUrl: "${baseUrl}";
          @import "@/scss/_variables.scss";
          @import "@/scss/_mixins.scss";
        `
      }
    }
  },
}