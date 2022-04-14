import {MessagePlugin} from "tdesign-react";

// TODO: ie的复制有问题，要看下啥报错，然后对应提示下用chrome
export const copyMessage = (msg: string) => {
    navigator.clipboard.writeText(msg).then(() => {
        MessagePlugin.success('复制成功', 2000)
    }).catch(() => {
        // 确保视图已经更新了再执行复制操作
        setTimeout(() => {
            const range = document.createRange();
            // @ts-ignore
            range.selectNode(document.getElementById('copyInner'));
            const selection = window.getSelection();
            // @ts-ignore
            if (selection.rangeCount > 0) selection.removeAllRanges();
            // @ts-ignore
            selection.addRange(range);
            document.execCommand('copy') ? MessagePlugin.success('复制成功', 2000) : MessagePlugin.error(`当前浏览器不支持复制该文本，请手动复制: ${msg}`, 10000);
        }, 0)
    })
}

export const objToQueryString = (obj: Object) => {
    let str = ''
    Object.entries(obj).forEach(([key, value]) => {
        str += `${encodeURIComponent(key)}=${encodeURIComponent(value)}&`;
    });
    return str
}
