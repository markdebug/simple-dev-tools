# JSON & 时间戳工具

## 项目简介

本项目是一个基于 Go 和 Gin 框架开发的 Web 工具，提供以下功能：
1. **JSON 工具**：
   - 格式化 JSON 数据。
   - 压缩 JSON 数据。
2. **时间戳工具**：
   - 时间戳转时间。
   - 获取当前时间戳。
   - 时间转时间戳。
   - 计算两个时间的差值（年、月、日、总天数）。
3. **文本比较工具**：
   - 比较两个文本是否一致，并在页面上提示结果。

前端使用 HTML、CSS 和 jQuery 实现，提供简洁美观的用户界面。

---

## 功能展示

### JSON 工具
- **格式化 JSON**：将输入的 JSON 数据美化为易读格式。
- **压缩 JSON**：将输入的 JSON 数据压缩为紧凑格式。

### 时间戳工具
- **时间戳转时间**：将时间戳转换为人类可读的时间格式。
- **获取当前时间戳**：获取当前时间的时间戳。
- **时间转时间戳**：将指定时间转换为时间戳。
- **计算时间差**：计算两个时间之间的差值，精确到年、月、日。

### 文本比较工具
- 比较两个输入框中的文本是否一致，并提示结果。

---

## 项目结构

```
tp-k8s/
├── main.go                # 后端主程序，基于 Gin 框架
├── static/
│   ├── index.html         # 前端页面，包含所有功能模块
│   └── ...                # 其他静态资源（如 CSS、JS 等）
└── README.md              # 项目说明文档
```

---

## 使用说明

### 环境要求
- Go 1.18 或更高版本
- 浏览器（推荐 Chrome）

### 启动项目
1. 克隆项目到本地：
   ```bash
   git clone https://github.com/your-repo/tp-k8s.git
   cd tp-k8s
   ```
2. 启动后端服务：
   ```bash
   go run main.go
   ```
3. 打开浏览器，访问 [http://localhost:8080](http://localhost:8080)。

---

## 前端功能截图

### JSON 工具
![JSON 工具](https://via.placeholder.com/800x400?text=JSON+工具)

### 时间戳工具
![时间戳工具](https://via.placeholder.com/800x400?text=时间戳工具)

### 文本比较工具
![文本比较工具](https://via.placeholder.com/800x400?text=文本比较工具)

---

## 贡献指南

欢迎贡献代码！请提交 Pull Request 或 Issue。

---

## 许可证

本项目采用 [MIT License](LICENSE) 开源。
