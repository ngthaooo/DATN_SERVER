import json
import nltk
from nltk.stem import WordNetLemmatizer
import numpy as np
import tensorflow as tf
from tensorflow.keras.models import load_model

nltk.download('punkt')
nltk.download('wordnet')

lemmatizer = WordNetLemmatizer()

# Load the trained model
model = load_model('/home/huythang/PycharmProjects/pythonProject/train/chatbot_model.h5')

# Load intents file
with open('/home/huythang/PycharmProjects/pythonProject/data/intents.json') as file:
    data = json.load(file)

words = []
classes = []
documents = []
ignore_words = ['?', '!', '.']

# Process intents file
for intent in data['intents']:
    for pattern in intent['patterns']:
        word_list = nltk.word_tokenize(pattern)
        words.extend(word_list)
        documents.append((word_list, intent['tag']))
        if intent['tag'] not in classes:
            classes.append(intent['tag'])

words = [lemmatizer.lemmatize(w.lower()) for w in words if w not in ignore_words]
words = sorted(list(set(words)))
classes = sorted(list(set(classes)))


def clean_up_sentence(sentence):
    sentence_words = nltk.word_tokenize(sentence)
    sentence_words = [lemmatizer.lemmatize(word.lower()) for word in sentence_words]
    return sentence_words


def bag_of_words(sentence):
    sentence_words = clean_up_sentence(sentence)
    bag = [0] * len(words)
    for w in sentence_words:
        for i, word in enumerate(words):
            if word == w:
                bag[i] = 1
    return np.array(bag)


def predict_class(sentence):
    bow = bag_of_words(sentence)
    res = model.predict(np.array([bow]))[0]
    ERROR_THRESHOLD = 0.25
    results = [[i, r] for i, r in enumerate(res) if r > ERROR_THRESHOLD]

    results.sort(key=lambda x: x[1], reverse=True)
    return_list = []
    for r in results:
        return_list.append({'intent': classes[r[0]], 'probability': str(r[1])})
    return return_list


def get_response(intents_list, intents_json):
    tag = intents_list[0]['intent']
    list_of_intents = intents_json['intents']
    for i in list_of_intents:
        if i['tag'] == tag:
            result = np.random.choice(i['responses'])
            break
    return result


# Chatbot loop
print("Chatbot: Hi! How can I help you today? (type 'quit' to exit)")
while True:
    message = input("You: ")
    if message.lower() == 'quit':
        break

    ints = predict_class(message)
    res = get_response(ints, data)
    print("Chatbot:", res)

print("Chatbot: Goodbye!")