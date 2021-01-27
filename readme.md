[中文（简体）版本](https://github.com/KarKLi/MIGA/blob/master/readme_zh-CN.md) | [English Version](https://github.com/KarKLi/MIGA/blob/master/readme.md)
* [MIGA \-  A system for simpilfying computer vision dnn deployment &amp; experiment](#miga----a-system-for-simpilfying-computer-vision-dnn-deployment--experiment)
  * [1\.Dataset prepartion](#1dataset-prepartion)
    * [I folder prepartion](#i-folder-prepartion)
    * [II annotation prepartion](#ii-annotation-prepartion)
  * [2\. Annotations verfication (MIGA's work)](#2-annotations-verfication-migas-work)
  * [3\.Deep Neural Network Hyperparameter Configuration (Optional)](#3deep-neural-network-hyperparameter-configuration-optional)
  * [4\.Choose the docker image (under construction)](#4choose-the-docker-image-under-construction)
  * [5\.Select the mode](#5select-the-mode)
  * [6\.Enjoy it\!](#6enjoy-it)
  * [7\. Some tips](#7-some-tips)
# MIGA -  A system for simpilfying computer vision dnn deployment & experiment
MIGA is a system unified many traditional computer vision dnns for teaching usage.

MIGA's architecture:
![](https://github.com/KarKLi/MIGA/blob/master/resources/MIGA_zh-CN.png)
## 1.Dataset prepartion
### I folder prepartion
MIGA users need to prepare the data by themselve. For preparation, divide your dataset into three parts - train, test and validate.

How to divide? Just create three folders on your dataset path: train, test and val.
Now you can put your images into the folders.

Notice: Unless inference mode, train and test folder should **NOT** be empty when start running MIGA.

### II annotation prepartion
MIGA also need to verify your images' annotation in order to prevent the potential train/inference error.

There are three optional formats for annotations - JSON, TXT and XML. Each format should follow the annotation format defined by MIGA.
Also, there are two ways to organize your annotations - total or each.

So, what is 'total annotations'?

It means there is only one JSON/TXT/XML file in your train/test/val folder, contains every annotations of the data, for example:
```
train
|---IMG_0001.jpg
|---IMG_0002.jpg
|---IMG_0003.jpg
|---annotation.txt
|--- ... (Other files)
```

and what is 'each annotations'?

It means each images has its own annotation file with the same name, for example:
```
train
|---IMG_0001.jpg
|---IMG_0001.txt
|---IMG_0002.jpg
|---IMG_0002.txt
|--- ... (Other files)
|---class.txt
```

OK, now let's talk about the format of annotation. In total annotations, there are 3 versions of 'total annotations'.

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

Also, there are 3 versions of 'each annotations'.

JSON：
```json
{"annotations":[{"x":450,"y":310,"w":100,"h":20,"c":4},{"x":800,"y":400,"w":30,"h":60,"c":0}],"img_name":"IMG_0001.jpg"}
```

TXT:
```plain
IMG_0001.jpg
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

By the way, if MIGA user select 'each annotations' expression, there must exist a file named 'class.txt' in each folders (train/test/val) to define the class.

An example of class.txt:
```plain
car human dog cat apple
```

Notice: the index of the class array begin with 0. For example, class[0] == "car" and class[4]=="apple"

## 2. Annotations verfication (MIGA's work)
Congratulations! Now you can have a rest and take a cup of tea, because MIGA will take over the dataset and start checking the annotations.

MIGA will check:

If an annotation coordinates cross over the image's boundary for each image's annotation(s).

## 3.Deep Neural Network Hyperparameter Configuration (Optional)
As usual, the DNN docker image will have a default configuration file to specify the network's hyperparameters such as batch size, learning rate, epoch, filping/cropping/grayscale, etc. If you want to change it, provide your own configuration file, or just change the default file (Remember backup the default one before you modify it).

The configuration file's format is inspried from [Darknet](https://github.com/AlexeyAB/darknet).

## 4.Choose the docker image (under construction)
OK, now it's your turn. Choose the network image that you want to load. MIGA provides several DNN images for training or inferencing.

Images are now avaliable at [MIGA-Images](https://github.com/KarKLi/MIGA-Images).
You can also download the image creating by others, just put it on 'MIGA-root-path'/images folder and it's done.

Current Avaliable images:

❌ ResNet-50

❌ VGGNet-19

## 5.Select the mode
After validating the image, you can select mode 'train' or 'inference'.

train - Using the data from train folder to learn the network param.

inference - Using the data from test folder to get the prediction.

## 6.Enjoy it!
Now you are running a network managed by MIGA!

If you choose 'train' mode, the weight file will save on 'MIGA-root-path'/weight/'image-name'.weights

If you choose 'inference' mode, the result annotation will save on 'MIGA-root-path'/inference/'image-name'+'YYMMDD'+'HHMM'.txt

## 7. Some tips
We recommend there is always only one MIGA running instance on your PC.