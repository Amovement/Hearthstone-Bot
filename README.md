# Hearthstone-Bot
- Hearthstone 炉石传说脚本 战令挂机 PVE投降 非天梯 PVE automatically surrenders
- 国服挂机巫妖王(pve) 定时自动投降 刷战令经验
- 摸鱼炉石

# 叠甲
- 本项目仅供学习交流 `Golang` 以及 `炉石传说` 玩法，请勿违反法律法规及游戏协议！
- 该项目仅用于学习 golang 桌面版 app 开发！
- 本项目运行时会在特定时间内占用鼠标使用，请勿在缺少备份的环境使用该项目！
- 脚本挂机每分钟只有 2 经验 请勿高强度使用！

# 开发环境
- Windows
- Go 1.14+

# 特性说明
- 提供了 `显示\隐藏` 炉石传说的功能，~~方便上班摸鱼~~
- 在冰霜堡垒界面点击开始挂机会进行每分钟 2 经验的自动投降策略， 这个行为会占用鼠标，即每隔固定的时间会抢占你的鼠标进行一些操作（如结束回合、投降等）
- 目前没有暂停脚本的功能，可以直接关掉 `HearthStoneBot.exe` 脚本自动停止，如果这个时候你的炉石界面被隐藏了，那么就重新启动 `HearthStoneBot.exe` 再点显示炉石即可。
- ![](https://s3.bmp.ovh/imgs/2025/03/20/25b63e881009b699.png)


# How to use
### 0. 下载 app
- 直接下载 release 文件
- 或本地编译执行 `buildgui.ps1` 
- 得到 `HearthStoneBot.exe`

### 1. 启动炉石传说并手动定位到冒险模式巫妖王（冰冠堡垒）如图所示位置
   ![](https://s3.bmp.ovh/imgs/2025/03/20/32735e8c68bafef8.png)

### 2. 运行 `HearthStoneBot.exe`
![](https://s3.bmp.ovh/imgs/2025/03/20/25b63e881009b699.png)

### 3. 按下开始挂机按钮