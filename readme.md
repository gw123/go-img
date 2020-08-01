#  
##加载网络图片
- LoadImageFromUrl(url string) (image.Image, error)
## CombineV 合并两张图片
- CombineV(upImg image.Image, downImg image.Image) (image.Image, error) 

# golang 官方包
```cassandraql
/***
dst 画布
r   在画布上的起点和终点两个坐标
src 要贴上的贴图
sp 贴图的图片的起点如果要全部贴上 设置image.Pt(0,0)
op 贴图和画布的叠加关系
**/
Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
```
