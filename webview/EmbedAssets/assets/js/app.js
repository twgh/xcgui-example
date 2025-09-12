
const maximizeBtn = document.querySelector('.maximize-btn');
let isMaximized = false;
// 切换窗口最大化, 并更新按钮图标
function toggleMaximize2(){
    isMaximized = !isMaximized;

    const icon = maximizeBtn.querySelector('i');
    if (isMaximized) {
        icon.classList.remove('fa-window-maximize');
        icon.classList.add('fa-window-restore');
    } else {
        icon.classList.remove('fa-window-restore');
        icon.classList.add('fa-window-maximize');
    }
    go.toggleMaximize();
}

// 切换侧边栏
const toggleBtn = document.querySelector('.toggle-btn');
const sidebar = document.querySelector('.sidebar');

toggleBtn.addEventListener('click', () => {
    sidebar.classList.toggle('collapsed');
});

// 切换页面
const navLinks = document.querySelectorAll('.nav-link');
const pages = document.querySelectorAll('.page');
const pageTitle = document.querySelector('.page-title');

navLinks.forEach(link => {
    link.addEventListener('click', (e) => {
        e.preventDefault();

        // 移除所有活动状态
        navLinks.forEach(l => l.classList.remove('active'));
        pages.forEach(p => p.classList.remove('active'));

        // 添加当前活动状态
        link.classList.add('active');

        const targetPage = link.getAttribute('href').substring(1);
        document.getElementById(targetPage).classList.add('active');

        // 更新页面标题
        pageTitle.textContent = link.querySelector('.nav-text').textContent;
    });
});

// 切换主题
const themeToggle = document.querySelector('.theme-toggle');

themeToggle.addEventListener('click', () => {
    document.body.classList.toggle('dark-theme');

    const icon = themeToggle.querySelector('i');
    if (document.body.classList.contains('dark-theme')) {
        icon.classList.remove('fa-moon');
        icon.classList.add('fa-sun');
    } else {
        icon.classList.remove('fa-sun');
        icon.classList.add('fa-moon');
    }
});

// 响应式处理
function handleResponsive() {
    if (window.innerWidth < 992) {
        sidebar.classList.add('collapsed');
    } else {
        sidebar.classList.remove('collapsed');
    }
}

window.addEventListener('resize', handleResponsive);
handleResponsive(); // 初始调用