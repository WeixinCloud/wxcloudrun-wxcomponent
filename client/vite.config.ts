import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'

// 详细配置信息 https://vitejs.dev/config/
export default defineConfig({
    build: {
        // wxcomponent 为微管家使用路径前缀
        assetsDir: 'wxcomponent/assets'
    },
    plugins: [react()],
    css: {
        preprocessorOptions: {
            less: {
                modifyVars: {
                },
            },
        },
    },
    server: {
        proxy: {
            '/api': {
                target: 'https://xxxxxxxxxx.com/',
                rewrite: path => path.replace(/^\/api/, ''),
                changeOrigin: true,
            }
        }
    }
})
