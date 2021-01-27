中文（简体）版本 | [English Version](https://github.com/KarKLi/MIGA/blob/master/param/readme.md)
* [MIGA configuration file parse layer](#miga-configuration-file-parse-layer)
  * [Supported configuration file format](#supported-configuration-file-format)
  * [How to organize the file?](#how-to-organize-the-file)
    * [1\. Text File(\.txt)](#1-text-filetxt)
# MIGA configuration file parse layer
This package contains the code of parsing parameters from configuration files.
## Supported configuration file format
✅ Text File (.txt)

❌ YAML File (.yaml) (Under construction)

❌ JSON File (.json) (Under construction)

## How to organize the file?
### 1. Text File(.txt)
We use '#' symbol for comment's beginning and it's only support single line comment.
for example:
```plain
# This is a comment
# Something will be recorded here
lr = 0.01 # This comment is also supported
```
We use the format 'key=value' to set the parameter. All supported parameters are shown on below table.
| Parameter    | Meaning                                                      |
| ------------ | ------------------------------------------------------------ |
| epoch        | epoch defines how many iterations should be exeucte for training all data. |
| batch        | batch defines how many images will be organised as a batch for forward-propagation. |
| subdivisions | subdivisions defines the ratio of dividing a batch into mini-batch. |
| height       | height defines the network input's image height.             |
| width        | width defines the network input's image height.              |
| channels     | channels defines the input image's channel number. Many traditional object detection/recognition neural network only accept the image which has 1 channel or 3 channels(RGB). |
| momentum     | momentum defines the momentum of changing the learning_rate_delta while back-propagation process. |
| decay        | decay defines the learning rate decay rate, decay on learning rate will be applied after each 100 iterations. |
| lr           | lr defines the ratio of changing the network paramater on back-propagation. |
| cpus         | cpus defines how many CPU cores will be used for transfer image from disk to memory. |

**Notice: Parameters' name are case sensitive.**

An example of .txt configuration file:
```plain
# An example of Text file.
epoch=150
batch = 128
subdivisions    =4
height  = 720
width=  1080
channels=3
# other parameters...
```