import numpy as np
import pickle
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import Dense, Dropout

# Đọc dữ liệu đã chuẩn bị
with open('training_data.pickle', 'rb') as f:
    words, classes, train_x, train_y = pickle.load(f)

# Xây dựng mô hình
model = Sequential()
model.add(Dense(128, input_shape=(len(train_x[0]),), activation='relu'))
model.add(Dropout(0.5))
model.add(Dense(64, activation='relu'))
model.add(Dropout(0.5))
model.add(Dense(len(train_y[0]), activation='softmax'))

# Biên dịch mô hình
model.compile(loss='categorical_crossentropy', optimizer='adam', metrics=['accuracy'])

# Huấn luyện mô hình
model.fit(train_x, train_y, epochs=200, batch_size=5, verbose=1)

# Lưu mô hình
model.save('chatbot_model.h5')

print("Mô hình đã được huấn luyện và lưu vào 'chatbot_model.h5'")