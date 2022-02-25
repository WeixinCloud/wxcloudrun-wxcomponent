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
                target: 'https://wxcomponent-1421032-1304480914.ap-shanghai.run.tcloudbase.com/',
                rewrite: path => path.replace(/^\/api/, ''),
                changeOrigin: true,
            }
        }
    }
})
