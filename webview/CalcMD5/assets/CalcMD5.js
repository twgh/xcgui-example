
// 选择文件
async function selectFile() {
    document.getElementById('filePath').value = await go.openFile();
}

// 计算MD5
async function calculate() {
    const path = document.getElementById('filePath').value;
    if (path === "") {
        showError('请先选择文件');
        return;
    }

    const btnCalcMd5 = document.getElementById('calcMd5');
    btnCalcMd5.disabled = true;
    btnCalcMd5.textContent = '计算中...'

    try {
        document.getElementById('result').textContent = await go.calculateMD5(path)
    } catch (error) {
        showError(error.message);
    } finally {
        btnCalcMd5.disabled = false
        btnCalcMd5.textContent = '计算MD5'
    }
}

// 显示错误
function showError(message) {
    document.getElementById('result').textContent = '错误: ' + message
}

// 拖放文件的区域
const dragArea = document.body;

dragArea.addEventListener('drop', (e) => {
    e.preventDefault();

    // 重要：通过 WebView2 的 chrome.webview.postMessageWithAdditionalObjects 发送消息,
    // 将 e.dataTransfer.files 中的文件作为附加对象传递
    if (window.chrome && chrome.webview) {
        chrome.webview.postMessageWithAdditionalObjects('drag_files', e.dataTransfer.files);
    }
});

dragArea.addEventListener('dragover', (e) => {
    e.preventDefault(); // 必须阻止默认行为以允许拖放
});
