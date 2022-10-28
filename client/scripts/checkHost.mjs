import fs from 'fs'
import { execSync } from 'child_process'
const command = 'git diff --name-status --staged HEAD'

const checkTarget = () => {
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
}

const checkGitStagedStatus = () => {
    const result = execSync(command).toString()
    // 只有git staged change的时候才需要做检查
    // 不然没提交进版本库也检查这也太蠢了
    if (result.includes('vite.config.ts')) {
        checkTarget()
    }
}

checkGitStagedStatus()
