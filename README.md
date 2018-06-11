# Image Classification - TensorFlow in Go
## Run locally
### Prerequisites
* Go 1.8+
* Node 6+
* Python 2.7+

### Install TensorFlow C binding
[Source](https://www.tensorflow.org/install/install_c): Execute the following shell command
```
 TF_TYPE="cpu" # Change to "gpu" for GPU support
 OS="linux" # Change to "darwin" for macOS
 TARGET_DIRECTORY="/usr/local"
 curl -L \
   "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-${TF_TYPE}-${OS}-x86_64-1.8.0.tar.gz" |
   sudo tar -C $TARGET_DIRECTORY -xz
```

Then configure the linker
```
sudo ldconfig
```

### Install Project Dependencies
Once TensorFlow is installed, next is to install Go dependency
```
dep ensure
```

Then compile Go source code
```
go install
```

Install node modules for building the frontend
```
npm install
```

Then build it
```
npm run build
```

