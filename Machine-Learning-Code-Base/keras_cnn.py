import glob
import os


from keras import backend as K
from keras.layers.core import Dense
from keras.layers.core import Dropout
from keras.layers.core import Flatten
from keras.layers.core import Activation
from keras.layers.convolutional import MaxPooling2D
from keras.layers.convolutional import Conv2D
from keras.layers.normalization import BatchNormalization
from keras.models import Sequential
from keras.callbacks import EarlyStopping
from keras.models import load_model
from keras.optimizers import rmsprop
from keras import utils

import cv2
import numpy as np
from matplotlib import pyplot as plt


def get_image(label, x, y):
    for i in range(0, len(y)):
        if np.argmax(y[i]) == label:
            return x[i]
    return None


def image_subset(index, x, y):
    """
    index equals to amount of labels
    """
    xs = []
    ys = []
    print("len(x) is {}".format(len(x)))
    for i in range(len(x)):
        if y[i] < index:
            xs.append(x[i])
            ys.append(y[i])
    return np.array(xs), np.array(ys)


def get_images(path_list):
    """
    read image dataset based on given path
    """
    images = []
    labels = []
    names = []
    i = 0
    for path in path_list:
        for fruit_dir_path in glob.glob(path):
            fruit_label = fruit_dir_path.split("/")[-1]
            for image_path in glob.glob(os.path.join(fruit_dir_path, "*.jpg")):
                image = cv2.imread(image_path, cv2.IMREAD_COLOR)

                image = cv2.resize(image, (45, 45))
                image = cv2.cvtColor(image, cv2.COLOR_RGB2BGR)

                images.append(image)
                names.append(fruit_label)
                labels.append(i)
            i += 1

    images = np.array(images)
    print(images.shape)
    # add a new dimension here
    with np.nditer(images, op_flags=['readwrite']) as it:
        for x in it:
            x = np.expand_dims(x, axis=0)
    labels = np.array(labels)
    return images, labels, i


class CNN(object):
    """
    A convolutional neural network
    """

    def __init__(self, train_directory, test_directory, num_classes, batch_size, epochs):
        self.train_directory = train_directory
        self.test_directory = test_directory
        self.num_classes = num_classes
        self.batch_size = batch_size
        self.epochs = epochs
        self.build_dataset()

    def build_dataset(self):
        """
        split it into proper train/test dataset
        """
        print("reading data of images currently , please wait......")
        x_train, y_train, _ = get_images(self.train_directory)
        x_test, y_test, _ = get_images(self.test_directory)
        x_train, y_train = image_subset(self.num_classes, x_train, y_train)
        x_test, y_test = image_subset(self.num_classes, x_test, y_test)
        x_train = x_train.astype('float32')
        x_test = x_test.astype('float32')
        self.x_train = x_train / 255
        self.x_test = x_test / 255
        self.y_train = utils.to_categorical(y_train, self.num_classes)
        self.y_test = utils.to_categorical(y_test, self.num_classes)

    def build_model(self, method):
        if method == 1:
            model = Sequential()
            model.add(Conv2D(32, (3, 3), activation='relu',
                             padding='same',
                             input_shape=(45, 45, 3)))
            model.add(Conv2D(32, (3, 3),
                             activation='relu'))
            model.add(MaxPooling2D((2, 2)))
            model.add(Dropout(0.25))
            model.add(Conv2D(64, (3, 3),
                             padding='same',
                             activation='relu'))
            model.add(Conv2D(64, (3, 3), activation='relu'))
            model.add(MaxPooling2D((2, 2)))
            model.add(Dropout(0.25))
            model.add(Flatten())
            model.add(Dense(128, activation='relu'))
            model.add(Dropout(0.5))
            model.add(Dense(num_classes, activation='softmax'))
            es = EarlyStopping(monitor='val_loss',
                               patience=15,
                               min_delta=0.01)
