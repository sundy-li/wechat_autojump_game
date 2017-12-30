## wechat_autojump_game
golang实现的自动玩微信跳一跳


#### 如何安装使用


#### 原理
- 利用adb shell截图游戏屏幕
- 读取截屏图片,获取当前位置,下一跳位置,计算跳动距离和触屏事件
- 利用adb shell发送input swipe事件来跳跃



