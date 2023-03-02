import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  // server: {// 代理属性
  //   host: '127.0.0.1', // can be overwritten by process.env.HOST
  //   port: 6666,
  //   proxy: {
  //     //axios跨域改造 by zhengkai.blog.csdn.net
  //     '/api': {
  //       target: 'http://127.0.0.1:9999/', // 你请求的第三方接口
  //       changeOrigin: true, // 在本地会创建一个虚拟服务端，然后发送请求的数据，并同时接收请求的数据，这样服务端和服务端进行数据的交互就不会有跨域问题
  //     }
  //   },
  // },
})
