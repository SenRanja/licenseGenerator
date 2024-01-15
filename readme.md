
# LicenseGenerator

因为发现fyne有非常强劲的跨平台能力，所以选择fyne作为GUI的开发组件。经过测试，生成的GUI界面包不大，且稳定可以以exe、apk的方式运行。

### license因素

经过我的考虑，决定在考试系统中插入暗桩，多因素校验是否符合license，来防止系统滥用。考试系统部署在docker中。

无法考量因素：
* 客户环境部分离线，license无法线上校验

考量因素：
* 运行时长（三处暗桩校验：数据库、本地文件、隐写方式校验，每小时叠加计算时间，目标运行时间藏在另一个图片隐写内容中。篡改其文件属性的Ctime时间防止被人发现这三个文件被频繁读取修改。）
* docker容器的eth0网卡的mac地址（docker生成容器默认具有不同的mac地址）

### 使用者

使用我提供的考试系统的docker命令安装后，需要

1. docker inspect进入获取eth0的mac地址
2. 运行时长

生成license，然后登陆考试系统，输入license，用来延长客户对考务系统的使用时间。

### license生成算法

IV、Key值固定（不写在此处）

1. 字符串拼接：`Ip + "shenyanjian" + Time`，编码为`[]byte`类型
2. 然后`AES-GCM`加密得到密文，直接对byte类型的密文进行`md5`生成hex摘要
3. 取前16字节
4. 四四分割，得到license，样子如:`34a6-0b6d-a1de-e321`

### 须知

为了避免本项目使用fyne编译时泄露机密，因此代码中不加注释用以破解困难。传递代码勿传递本`readme.md`文件

apk版安卓暂时由于未充分测试其安全性与是否泄露生成license算法，暂不发布。

windows版exe由于使用fyne生成，但是为了安全性使用 go build生成，未插入图标，因此我只会给go build生成的exe的license生成器作为交付。

