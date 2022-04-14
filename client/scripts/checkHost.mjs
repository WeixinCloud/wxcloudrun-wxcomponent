import fs from 'fs'

const container = fs.readFileSync('vite.config.ts').toString()

const lines = container.split('\n')

const apiIndex = lines.findIndex(l => l.includes('\'/api\''))

const apiArr = lines.slice(apiIndex)

const target = apiArr.find(e => e.includes('target:')).trim()

if (target.includes('https://xxxxxxxxxx.com/')) {
    process.exit()
} else {
    console.log(`🤔 请不要把线上的网址 push 到 github 上，请把 vite.config.ts 里的 proxy.api.target 改为 https://xxxxxxxxxx.com/`)
    process.exit(-1)
}
