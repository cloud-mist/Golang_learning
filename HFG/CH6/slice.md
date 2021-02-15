# 切片

## 定义
- 与数组相同：有多个相同类型元素构成
- 不同：允许我们在结尾追加更多元素

## 声明
- 不指定大小：`var mySlice []string`
- make创建:`notes = make([]string, 7)`
- 切片字面量：`notes := []string{8, 19, 23}`
