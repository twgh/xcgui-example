
// 选择文件
async function selectFile() {
    const path = await go.openFile();
    if (path === "") {
        return;
    }
    document.getElementById('filePath').value = path;
    await calculate(path);
}

// 计算MD5
async function calculate(path) {
    if (path === "") {
        return;
    }

    const btnSelectFile = document.getElementById('selectFile');
    btnSelectFile.disabled = true;
    btnSelectFile.textContent = '计算中...'

    try {
        document.getElementById('result').textContent = await go.calculateMD5(path)
    } catch (error) {
        showError(error.message);
    } finally {
        btnSelectFile.disabled = false
        btnSelectFile.textContent = '选择文件'
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
