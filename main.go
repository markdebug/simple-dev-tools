package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @title Port Management API
// @version 1.0
// @description Windows 主机端口管理 API
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// 静态文件服务
	r.Static("/static", "./static")

	// 首页重定向到静态页面
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/index.html")
	})

	// 格式化JSON接口
	r.POST("/api/format-json", func(c *gin.Context) {
		var req struct {
			Input string `json:"input"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
			return
		}
		var obj interface{}
		if err := json.Unmarshal([]byte(req.Input), &obj); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "JSON解析失败"})
			return
		}
		formatted, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "格式化失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": string(formatted)})
	})

	// 压缩JSON接口
	r.POST("/api/compress-json", func(c *gin.Context) {
		var req struct {
			Input string `json:"input"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
			return
		}
		var obj interface{}
		if err := json.Unmarshal([]byte(req.Input), &obj); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "JSON解析失败"})
			return
		}
		compressed, err := json.Marshal(obj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "压缩失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": string(compressed)})
	})

	// 时间戳转换接口（支持时区）
	r.POST("/api/convert-timestamp", func(c *gin.Context) {
		var req struct {
			Timestamp string `json:"timestamp"`
			Timezone  string `json:"timezone"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
			return
		}
		ts, err := strconv.ParseInt(req.Timestamp, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "时间戳解析失败"})
			return
		}
		loc, err := time.LoadLocation(req.Timezone)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "无效的时区"})
			return
		}
		tm := time.Unix(ts, 0).In(loc)
		c.JSON(http.StatusOK, gin.H{"result": tm.Format("2006-01-02 15:04:05")})
	})

	// 获取当前时间戳接口
	r.GET("/api/current-timestamp", func(c *gin.Context) {
		now := time.Now().Unix()
		c.JSON(http.StatusOK, gin.H{"result": now})
	})

	// 时间转时间戳接口（支持时区）
	r.POST("/api/time-to-timestamp", func(c *gin.Context) {
		var req struct {
			DateTime string `json:"datetime"`
			Timezone string `json:"timezone"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
			return
		}
		loc, err := time.LoadLocation(req.Timezone)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "无效的时区"})
			return
		}
		tm, err := time.ParseInLocation("2006-01-02 15:04:05", req.DateTime, loc)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "时间格式错误，请使用 YYYY-MM-DD HH:MM:SS 格式"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": tm.Unix()})
	})

	// 计算时间差接口
	r.POST("/api/time-diff", func(c *gin.Context) {
		var req struct {
			StartTime string `json:"start_time"`
			EndTime   string `json:"end_time"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
			return
		}

		start, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "开始时间格式错误"})
			return
		}

		end, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "结束时间格式错误"})
			return
		}

		// 计算时间差
		duration := end.Sub(start)
		days := int(duration.Hours() / 24)

		// 计算年月日差
		years := end.Year() - start.Year()
		months := int(end.Month()) - int(start.Month())
		dayDiff := end.Day() - start.Day()

		if dayDiff < 0 {
			months--
			// 获取上个月的天数
			prevMonth := end.AddDate(0, -1, 0)
			daysInPrevMonth := time.Date(prevMonth.Year(), prevMonth.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
			dayDiff += daysInPrevMonth
		}

		if months < 0 {
			years--
			months += 12
		}

		result := gin.H{
			"years":      years,
			"months":     months,
			"days":       dayDiff,
			"total_days": days,
		}

		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	// swagger 路由
	r.Run(":8080")
}
