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
                target: 'https://xxxx.com',
                rewrite: path => path.replace(/^\/api/, ''),
                changeOrigin: true,
            }
        }
    }
})
