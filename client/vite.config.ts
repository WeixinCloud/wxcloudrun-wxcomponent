import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
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
                target: 'https://srv20211231143005-1524266-1309057895.ap-shanghai.run.tcloudbase.com/',
                rewrite: path => path.replace(/^\/api/, ''),
                changeOrigin: true,
            }
        }
    }
})
