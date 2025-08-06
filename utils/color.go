package utils

import (
	"github.com/fatih/color"
)

// 由专业的库来处理多平台的颜色输出效果
var (
	Red     = color.New(color.FgRed)                // 红色 31
	Green   = color.New(color.FgGreen)              // 绿色 32
	Yellow  = color.New(color.FgYellow)             // 黄色 33
	Blue    = color.New(color.FgBlue, color.Bold)   // 蓝色 34
	Magenta = color.New(color.FgMagenta)            // 紫红色 35
	White   = color.New(color.FgWhite)              // 白色 37
	Pink    = color.RGB(255, 182, 193).Add(color.Bold) //粉色,感觉好看
)



