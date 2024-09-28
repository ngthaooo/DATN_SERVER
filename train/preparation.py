import json
import nltk
from nltk.stem import WordNetLemmatizer
import numpy as np
import pickle

# Tải các gói NLTK cần thiết
nltk.download('punkt')
nltk.download('wordnet')

lemmatizer = WordNetLemmatizer()

# Đọc dữ liệu từ file JSON
with open('/home/huythang/PycharmProjects/pythonProject/data/intents.json') as file:
    data = json.load(file)

words = []
classes = []
documents = []
ignore_words = ['?', '!', '.']

# Tokenize và lemmatize dữ liệu
for category, intents in data.items():  # Duyệt qua từng nhóm trong dữ liệu
    for intent in intents:
        pattern = intent['question']
        word_list = nltk.word_tokenize(pattern)
        words.extend(word_list)
        documents.append((word_list, intent['intent']))
        if intent['intent'] not in classes:
            classes.append(intent['intent'])

# Chuẩn hóa từ
words = [lemmatizer.lemmatize(w.lower()) for w in words if w not in ignore_words]
words = sorted(list(set(words)))

classes = sorted(list(set(classes)))

training = []
output_empty = [0] * len(classes)

# Tạo bag of words và đầu ra cho các mẫu
for doc in documents:
    bag = []
    word_patterns = doc[0]
    word_patterns = [lemmatizer.lemmatize(word.lower()) for word in word_patterns]

    for word in words:
        bag.append(1) if word in word_patterns else bag.append(0)

    output_row = list(output_empty)
    output_row[classes.index(doc[1])] = 1

    training.append([bag, output_row])

# Chuyển đổi sang numpy array
training = np.array(training, dtype=object)

# Tách thành X (dữ liệu đầu vào) và y (dữ liệu đầu ra)
train_x = np.array(list(training[:, 0]))
train_y = np.array(list(training[:, 1]))

# Lưu dữ liệu đã chuẩn bị
with open('training_data.pickle', 'wb') as f:
    pickle.dump((words, classes, train_x, train_y), f)

print("Dữ liệu đã được chuẩn bị và lưu vào 'training_data.pickle'")
