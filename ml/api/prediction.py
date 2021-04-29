import face_recognition
import cv2
import numpy as np
import os
from PIL import Image, ImageDraw
from io import BytesIO

classNames = []


def loadImages():
    path = './images/images_training'
    images = []
    myList = os.listdir(path)

    for cl in myList:
        curImg = cv2.imread(f'{path}/{cl}')
        images.append(curImg)
        classNames.append(os.path.splitext(cl)[0])
    return images


def trainingImages(images):
    encodeList = []
    for img in images:
        # img = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)
        encode = face_recognition.face_encodings(img)[0]
        encodeList.append(encode)
    print('Training Complete')
    return encodeList
# encodeListKnown = trainingImages(images)


def testImage(fileName, encodeListKnown):
    # path = "./images/images_test/" + fileName
    img = face_recognition.load_image_file(
        fileName)

    face_locations = face_recognition.face_locations(img)
    face_encodings = face_recognition.face_encodings(img, face_locations)

    face_ids = []
    for encodeFace, faceLoc in zip(face_encodings, face_locations):
        matches = face_recognition.compare_faces(encodeListKnown, encodeFace)
        faceDis = face_recognition.face_distance(encodeListKnown, encodeFace)

        matchIndex = np.argmin(faceDis)
        if matches[matchIndex]:
            name = classNames[matchIndex].upper()
            face_ids.append(name)
    print("face==> ", face_ids)
    return face_ids


def read_imagefile(file) -> Image.Image:
    image = Image.open(BytesIO(file))
    image = np.asarray(image.resize((224, 224)))[..., :3]
    image = np.expand_dims(image, 0)
    image = image / 127.5 - 1.0
    return image


# path = "./images/images_test/zellon.jpeg"
# print(face_recognition.load_image_file(
#     path))
