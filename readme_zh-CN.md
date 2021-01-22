[中文（简体）版本](https://github.com/KarKLi/MIGA/blob/master/readme_zh-CN.md) | [English Version](https://github.com/KarKLi/MIGA/blob/master/readme.md)
* [MIGA \-  一种用于简化教学用CV\-Dnn部署和实验过程的工具](#miga----一种用于简化教学用cv-dnn部署和实验过程的工具)
  * [1\.数据集准备](#1数据集准备)
    * [I\.文件夹准备](#i文件夹准备)
    * [II\.标注文件的准备](#ii标注文件的准备)
  * [2\.标注验证（由MIGA完成）](#2标注验证由miga完成)
  * [3\.DNN超参数配置（可选）](#3dnn超参数配置可选)
  * [4\.选择你的docker映像（建设中）](#4选择你的docker映像建设中)
  * [5\.选择MIGA运行模式](#5选择miga运行模式)
  * [6\.MIGA在向你招手！](#6miga在向你招手)
  * [7\.一些小提示](#7一些小提示)
# MIGA -  一种用于简化教学用CV-Dnn部署和实验过程的工具
MIGA是一套集成了多种经典计算机视觉网络（ResNet,GoogLeNet,VGGNet,R-CNN等等），可用于教学上快速复现和实验网络的框架。
## 1.数据集准备
### I.文件夹准备
使用MIGA的用户需要自己准备图像数据集，并将它们按照需求划分到三个文件夹中：train文件夹，test文件夹和val文件夹。

注意：除非将MIGA设置为推理模式，否则train和test文件夹都不能为空。

### II.标注文件的准备
标注文件有三种可选格式：JSON,TXT和XML。无论采用哪种格式，都需要遵守MIGA标注规范。而且，标注有两种存放方式：整体标注存放和分图像标注存放。

首先我们来介绍整体标注存放。整体标注存放指的是文件夹下所有图片的标注全部在一个标注文件内，以train文件夹为例：
```
train
|---IMG_0001.jpg
|---IMG_0002.jpg
|---IMG_0003.jpg
|---annotation.txt
|--- ... (Other files)
```

那么什么是分图像标注存放呢？分图像标注存放指的是每一张图片都有一个对应的标注文件用于记录该图像中具有的一个或多个标注，还是以train文件夹为例：
```
train
|---IMG_0001.jpg
|---IMG_0001.txt
|---IMG_0002.jpg
|---IMG_0002.txt
|--- ... (Other files)
|---class.txt
```

讲完了存放方法，现在我们来讲一下MIGA的标注文件格式定义。

对于整体标注存放来说，有三种标注文件的格式定义，其例子见下：
JSON：
```json
{"annotations":{
    "IMG_0001.jpg":[{"x":800,"y":400,"w":30,"h":60,"c":0}],
    "IMG_0002.jpg":[{"x":450,"y":310,"w":100,"h":20,"c":4},{"x":800,"y":400,"w":30,"h":60,"c":0}]
},"class":["car","human","dog","cat","apple"]
}
```

TXT:
```plain
car human dog cat apple
IMG_0001.jpg 800 400 30 60 0
IMG_0002.jpg 450 310 100 20 4
IMG_0002.jpg 800 400 30 60 0
# other image annotations...
```

XML:
```xml
<?xml version="1.0" encoding="utf-8"?>
<root>
    <file filename="IMG_0001.jpg">
        <box x="800" y="400" w="30" h="60" c="0"/>
    </file>
    <file filename="IMG_0002.jpg">
        <box x="450" y="310" w="100" h="20" c="4"/>
        <box x="800" y="400" w="30" h="60" c="0"/>
    </file>
    <!-- other image notations... -->
</root>
```

同样地，分图像标注存放也有三种标注文件的格式定义，例子见下：
JSON：
```json
{"annotations":[{"x":450,"y":310,"w":100,"h":20,"c":4},{"x":800,"y":400,"w":30,"h":60,"c":0}]}
```

TXT:
```plain
450 310 100 20 4
800 400 30 60 0
# other annotations...
```

XML:
```xml
<?xml version="1.0" encoding="utf-8"?>
<root>
    <file filename="IMG_0002.jpg">
        <box x="450" y="310" w="100" h="20" c="4"/>
        <box x="800" y="400" w="30" h="60" c="0"/>
        <!-- other notations... -->
    </file>
</root>
```

如果你选择了分图像标注存放，那么在train/test/val文件夹下，只要有图像，就必须要有一个class.txt以标明你的类目。

一个class.txt的例子见下：
```plain
car human dog cat apple
```

注意：class.txt的下标从0开始计算，也就是说class[0]=="car", class[4]=="apple"

## 2.标注验证（由MIGA完成）
恭喜你！你已经完成了数据集的准备工作。接下来，MIGA会开始验证三个文件夹中标注数据的合理性。

MIGA会检测：

是否有标注的边界会越过图像的边界，这种情况下在MIGA是不被允许的。若发现这样的标注，请修改后再重新执行该步骤。

## 3.DNN超参数配置（可选）
通常来说，DNN docker映像都会有一个默认的配置文件来指定DNN的超参数如：batch size, learning rate, epoch, fliping/cropping/grayscale这些。如果你需要自定义你的DNN超参数，你可以向MIGA提供你自己的配置文件，或直接修改默认的配置文件（在修改默认配置文件前请先备份它）。

配置文件的语法细则受到[Darknet](https://github.com/AlexeyAB/darknet)的启发。

## 4.选择你的docker映像（建设中）
现在是选择网络的时候了，选择一个配置好的docker映像以载入你的DNN进行训练或者推理。MIGA提供了几款用于训练或推理的DNN docker映像，它们将会被逐步上传到 [MIGA-Images](https://github.com/KarKLi/MIGA-Images)。你也可以去下载其他用户提供的DNN docker映像（只要他们遵守MIGA-Images的开发细则），下载的DNN docker映像都会存放到本地的'MIGA-root-path'/images路径。

目前可用的镜像：

❌ ResNet-50

❌ VGGNet-19

## 5.选择MIGA运行模式
验证完镜像的可用性之后，你需要选择MIGA的运行模式。目前有两种运行模式可选：训练或者推理。

训练模式——使用train文件夹内的数据对网络进行训练。

推理模式——使用test文件夹内的数据进行前向传播并输出预测结果。

## 6.MIGA在向你招手！
如果前面的步骤都顺利，你现在应该已经在运行由MIGA进行托管的深度神经网络了！

如果你选择了训练模式，权重文件将会保存在'MIGA-root-path'/weight/'image-name'.weights

如果你选择了推理模式，推理结果将会保存在'MIGA-root-path'/inference/'image-name'+'YYMMDD'+'HHMM'.txt

## 7.一些小提示
我们推荐在每一台机子上只有一个MIGA实例在运行以避免潜在的竞争冲突。