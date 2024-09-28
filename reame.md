Loại mô hình: Đây là một mô hình mạng nơ-ron nhân tạo (Artificial Neural Network - ANN) đa lớp, còn được gọi là mạng nơ-ron sâu (Deep Neural Network).
Kiến trúc:

Mô hình sử dụng kiến trúc Sequential của Keras/TensorFlow.
Nó bao gồm các lớp Dense (fully connected) và Dropout.


Cấu trúc chi tiết:

Lớp đầu vào: Dense layer với 128 nơ-ron, sử dụng hàm kích hoạt ReLU.
Lớp Dropout (50%) để giảm overfitting.
Lớp ẩn: Dense layer với 64 nơ-ron, sử dụng hàm kích hoạt ReLU.
Lớp Dropout khác (50%).
Lớp đầu ra: Dense layer với số nơ-ron bằng số lớp (classes), sử dụng hàm kích hoạt softmax.


Mục đích: Mô hình này được thiết kế để phân loại ý định (intent classification) trong một hệ thống chatbot. Nó nhận đầu vào là một vector biểu diễn câu (dưới dạng bag-of-words) và đưa ra xác suất cho mỗi ý định có thể.
Phương pháp huấn luyện:

Loss function: Categorical Crossentropy (phù hợp cho bài toán phân loại nhiều lớp).
Optimizer: Adam
Metric: Accuracy


Preprocessing: Mô hình sử dụng các kỹ thuật xử lý ngôn ngữ tự nhiên (NLP) cơ bản như tokenization và lemmatization (sử dụng NLTK) để chuẩn bị dữ liệu đầu vào.


```
gen .proto python:

python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. pb.proto
```

``` 
gen code golang

protoc --go_out=. --go-grpc_out=. pb.proto
```

```commandline
# Cài đặt các công cụ cần thiết
sudo apt install -y autoconf automake libtool curl make g++ unzip

# Tải mã nguồn
wget https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protobuf-all-21.12.zip

# Giải nén và cài đặt
unzip protobuf-all-21.12.zip
cd protobuf-21.12
mkdir build
cd build
cmake ..
make
sudo make install
sudo ldconfig

```

```commandline
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

```

```commandline
protoc --version
protoc-gen-go --version
protoc-gen-go-grpc --version

```

```commandline
sudo apt update
sudo apt install -y protobuf-compiler

```