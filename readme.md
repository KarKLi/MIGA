中文（简体）版本 | English Version
# MIGA -  A system for simpilfying computer vision dnn deployment & experiment.
MIGA is a system unified many traditional computer vision dnns for teaching usage.
## 1.Dataset prepartion
### I folder prepartion
MIGA users need to prepare the data by themselve. For preparation, divide your dataset into three parts - train, test and validate.
How to divide? Just create three folders on your dataset path: train, test and val.
Now you can put your images into the folders.
Notice: Except val folder, train and test folder should NOT be empty when start running MIGA.

### II annotation prepartion
MIGA also need to verify your images' annotation in order to prevent the potential train/inference error.
There are three optional formats for annotations - JSON, TXT and XML. Each format should follow the annotation format defined by MIGA.
Also, there are two ways to organize your annotations - total or each.

So, what is 'total annotations'?

It means there is only one JSON/TXT/XML file in your train/test/val folder, contains every annotations of the data, for example:
```
train
|___IMG_0001.jpg
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
```

OK, now let's talk about the format of annotation. In total annotations, there are 3 versions of 'total annotations'.

JSON：
```json
{"annotations":{
    "IMG_0001.jpg":[{"x":800,"y":400,"w":30,"h":60,"c":0}],
    "IMG_0002.jpg":[{"x":450,"y":310,"w":100,"h":20,"c":4},{"x":800,"y":400,"w":30,"h":60,"c":0}],
    // other image annotations...
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

Also, there are 3 versions of each annotations.

JSON：
```json
{"annotations":[{"x":450,"y":310,"w":100,"h":20,"c":4},{"x":800,"y":400,"w":30,"h":60,"c":0}/*other annotations...*/]}
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

By the way, if MIGA user select 'each annotations' expression, there must exist a file named 'class.txt' in each folder (train/test/val) to define the class.

An example of class.txt:
```plain
car human dog cat apple
```

Notice: the index of the class array begin with 0. For example, class[0] == "car" and class[4]=="apple"