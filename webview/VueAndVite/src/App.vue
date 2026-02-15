<template>
  <div class="app">
    <!-- 左侧导航栏 -->
    <nav class="sidebar">
      <div class="logo">
        <h2>我的应用</h2>
      </div>
      <ul class="nav-menu">
        <li 
          v-for="item in menuItems" 
          :key="item.id"
          :class="{ active: activeMenu === item.id }"
          @click="setActiveMenu(item.id)"
        >
          <span class="icon">{{ item.icon }}</span>
          <span class="text">{{ item.name }}</span>
        </li>
      </ul>
      <div class="sidebar-footer">
        <p class="version">v1.0.0</p>
      </div>
    </nav>

    <!-- 右侧内容区 -->
    <main class="content">
      <header class="header">
        <h1>{{ currentMenuItem?.name || '欢迎' }}</h1>
        <div class="window-controls">
          <button @click="minimizeWindow" class="control-btn minimize">─</button>
          <button @click="toggleMaximize" class="control-btn maximize">☐</button>
          <button @click="closeWindow" class="control-btn close">✕</button>
        </div>
      </header>
      <div class="main-content">
        <div v-if="activeMenu === 'dashboard'" class="dashboard">
          <div class="welcome-card">
            <h2>欢迎使用我的应用</h2>
            <p>这是一个基于 Vue 3 + Vite + WebView2 的现代桌面应用, 开发时支持热重载.</p>
          </div>
          <div class="stats-grid">
            <div class="stat-card">
              <div class="stat-number">1,234</div>
              <div class="stat-label">总用户数</div>
            </div>
            <div class="stat-card">
              <div class="stat-number">567</div>
              <div class="stat-label">今日访问</div>
            </div>
            <div class="stat-card">
              <div class="stat-number">89</div>
              <div class="stat-label">新增订单</div>
            </div>
            <div class="stat-card">
              <div class="stat-number">99.9%</div>
              <div class="stat-label">系统可用性</div>
            </div>
          </div>
        </div>
        <div v-else-if="activeMenu === 'settings'" class="settings">
          <h2>设置</h2>
          <div class="settings-list">
            <div class="setting-item">
              <label>主题设置</label>
              <select>
                <option>浅色主题</option>
                <option>深色主题</option>
                <option>跟随系统</option>
              </select>
            </div>
            <div class="setting-item">
              <label>语言设置</label>
              <select>
                <option>简体中文</option>
                <option>English</option>
              </select>
            </div>
            <div class="setting-item">
              <label>自动更新</label>
              <input type="checkbox" checked>
            </div>
          </div>
        </div>
        <div v-else-if="activeMenu === 'about'" class="about">
          <h2>关于</h2>
          <div class="about-content">
            <p>版本: 1.0.0</p>
            <p>构建时间: {{ new Date().toLocaleDateString() }}</p>
            <p>技术栈:</p>
            <ul>
              <li>Vue 3</li>
              <li>Vite</li>
              <li>WebView2</li>
              <li>XCGUI</li>
              <li>Go</li>
            </ul>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const activeMenu = ref('dashboard')

const menuItems = [
  { id: 'dashboard', name: '仪表盘', icon: '📊' },
  { id: 'settings', name: '设置', icon: '⚙️' },
  { id: 'about', name: '关于', icon: 'ℹ️' }
]

const currentMenuItem = computed(() => {
  return menuItems.find(item => item.id === activeMenu.value)
})

const setActiveMenu = (id) => {
  activeMenu.value = id
}

// 窗口控制函数
const minimizeWindow = () => {
  if (window.wnd && window.wnd.minimize) {
    window.wnd.minimize()
  }
}

const toggleMaximize = () => {
  if (window.wnd && window.wnd.toggleMaximize) {
    window.wnd.toggleMaximize()
  }
}

const closeWindow = () => {
  if (window.wnd && window.wnd.close) {
    window.wnd.close()
  }
}
</script>

<style scoped>
.app {
  display: flex;
  height: 100vh;
  overflow: hidden;
  background: #f5f5f5;
}

/* 侧边栏样式 */
.sidebar {
  width: 250px;
  background: linear-gradient(180deg, #1a1a2e 0%, #16213e 100%);
  color: white;
  display: flex;
  flex-direction: column;
  box-shadow: 2px 0 10px rgba(0, 0, 0, 0.1);
  app-region: drag; /* 可拖动 */
}

.logo {
  padding: 30px 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.logo h2 {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
  background: linear-gradient(90deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.nav-menu {
  list-style: none;
  padding: 20px 0;
  margin: 0;
  flex: 1;
}

.nav-menu li {
  padding: 15px 20px;
  display: flex;
  align-items: center;
  cursor: pointer;
  transition: all 0.3s ease;
  border-left: 3px solid transparent;
  app-region: no-drag; /* 不可拖动 */
}

.nav-menu li:hover {
  background: rgba(255, 255, 255, 0.1);
}

.nav-menu li.active {
  background: rgba(102, 126, 234, 0.2);
  border-left-color: #667eea;
}

.nav-menu .icon {
  font-size: 20px;
  margin-right: 12px;
}

.nav-menu .text {
  font-size: 15px;
}

.sidebar-footer {
  padding: 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.version {
  margin: 0;
  font-size: 12px;
  color: rgba(255, 255, 255, 0.5);
}

/* 内容区样式 */
.content {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: #f5f5f5;
}

.header {
  background: white;
  padding: 20px 30px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  app-region: drag; /* 可拖动 */
}

.header h1 {
  margin: 0;
  font-size: 24px;
  color: #333;
  font-weight: 600;
}

.window-controls {
  display: flex;
  gap: 8px;
  app-region: no-drag; /* 不可拖动 */
}

.control-btn {
  width: 36px;
  height: 36px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.control-btn:hover {
  opacity: 0.8;
  transform: scale(1.05);
}

.control-btn.minimize {
  background: #ffd93d;
  color: #333;
}

.control-btn.maximize {
  background: #6bcb77;
  color: white;
}

.control-btn.close {
  background: #ff6b6b;
  color: white;
}

.main-content {
  flex: 1;
  padding: 30px;
  overflow-y: auto;
}

/* 仪表盘样式 */
.welcome-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 40px;
  border-radius: 12px;
  margin-bottom: 30px;
  box-shadow: 0 10px 30px rgba(102, 126, 234, 0.3);
}

.welcome-card h2 {
  margin: 0 0 10px 0;
  font-size: 28px;
}

.welcome-card p {
  margin: 0;
  font-size: 16px;
  opacity: 0.9;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
}

.stat-card {
  background: white;
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  text-align: center;
  transition: transform 0.3s, box-shadow 0.3s;
}

.stat-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
}

.stat-number {
  font-size: 36px;
  font-weight: bold;
  color: #667eea;
  margin-bottom: 10px;
}

.stat-label {
  color: #666;
  font-size: 14px;
}

/* 设置页面样式 */
.settings h2 {
  margin-bottom: 30px;
  color: #333;
}

.settings-list {
  background: white;
  border-radius: 12px;
  padding: 30px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
}

.setting-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 0;
  border-bottom: 1px solid #eee;
}

.setting-item:last-child {
  border-bottom: none;
}

.setting-item label {
  font-size: 16px;
  color: #333;
}

.setting-item select,
.setting-item input[type="checkbox"] {
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
}

/* 关于页面样式 */
.about h2 {
  margin-bottom: 30px;
  color: #333;
}

.about-content {
  background: white;
  border-radius: 12px;
  padding: 30px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
}

.about-content p {
  margin-bottom: 15px;
  color: #666;
  font-size: 16px;
}

.about-content ul {
  margin: 0;
  padding-left: 20px;
}

.about-content li {
  color: #666;
  margin-bottom: 8px;
}

/* 滚动条样式 */
.main-content::-webkit-scrollbar {
  width: 8px;
}

.main-content::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

.main-content::-webkit-scrollbar-thumb {
  background: #667eea;
  border-radius: 4px;
}

.main-content::-webkit-scrollbar-thumb:hover {
  background: #5568d3;
}
</style>
