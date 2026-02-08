/**
 * 窗口拖动工具 - 不依赖 CSS 的 app-region: drag;
 * 通过 JS 调用后端 API 实现窗口拖动.
 * 会为元素添加 CSS 的 user-select: none 防止拖动时选中文本.
 */

class WindowDrag {
  /**
   * 为指定元素启用窗口拖动功能
   * @param {string|HTMLElement|Array<string|HTMLElement>} element - 要启用拖动的元素、选择器或数组
   * @param {string|Array<string|HTMLElement>} exclude - 要排除的元素或选择器（点击这些元素时不触发拖动）, 可不填
   * @param {Function} setPosition - 设置窗口位置的函数 (x, y) => void, 可不填，默认使用 wnd.setPos
   * @returns {Function|Array<Function>} 取消绑定函数或函数数组
   */
  static enable(element, exclude, setPosition) {
    // 如果是数组，则批量启用
    if (Array.isArray(element)) {
      return element.map(el => this._enableSingle(el, exclude, setPosition));
    }
    // 单个元素
    return this._enableSingle(element, exclude, setPosition);
  }

  /**
   * 为单个元素启用窗口拖动（内部方法）
   * @private
   */
  static _enableSingle(element, exclude, setPosition) {
    const targetEl = typeof element === 'string'
      ? document.querySelector(element)
      : element;

    if (!targetEl) {
      console.error('[WindowDrag] 目标元素不存在:', element);
      return;
    }

    // 确定设置位置的函数
    const moveFn = typeof setPosition === 'function'
      ? setPosition
      : (x, y) => window['wnd']?.setPos(x, y);

    // 处理排除列表
    const excludeList = exclude
      ? (Array.isArray(exclude) ? exclude : [exclude]).map(item =>
          typeof item === 'string' ? document.querySelector(item) : item
        ).filter(Boolean)
      : [];

    // 检查元素是否在排除列表中
    const isExcluded = (el) => {
      return excludeList.some(excludeEl => excludeEl === el || excludeEl.contains(el));
    };

    let isDragging = false;
    let offsetX = 0;
    let offsetY = 0;

    // 添加 user-select: none 防止拖动时选中文本
    targetEl.style.userSelect = 'none';

    // 鼠标按下事件
    const onMouseDown = (e) => {
      // 只有左键才能拖动
      if (e.button !== 0) return;

      // 检查是否点击在排除元素上
      if (isExcluded(e.target)) return;

      isDragging = true;
      // 获取鼠标在窗口内的相对位置
      offsetX = e.clientX;
      offsetY = e.clientY;
    };

    // 鼠标移动事件
    const onMouseMove = (e) => {
      if (!isDragging) return;

      // 计算新位置
      const newX = e.screenX - offsetX;
      const newY = e.screenY - offsetY;

      // 调用移动函数
      moveFn(newX, newY);
    };

    // 鼠标抬起事件
    const onMouseUp = () => {
      isDragging = false;
    };

    // 绑定事件
    targetEl.addEventListener('mousedown', onMouseDown);
    document.addEventListener('mousemove', onMouseMove);
    document.addEventListener('mouseup', onMouseUp);

    // 返回取消绑定的函数
    return () => {
      targetEl.removeEventListener('mousedown', onMouseDown);
      document.removeEventListener('mousemove', onMouseMove);
      document.removeEventListener('mouseup', onMouseUp);
    };
  }
}

// 默认导出
window.WindowDrag = WindowDrag;
