from flask import Flask, request, redirect
# from werkzeug.utils import secure_filename
import os
import json
# import requests
import base64
import prediction
from flask_cors import CORS
app = Flask(__name__)
CORS(app)


UPLOAD_FOLDER = 'received_files'
ALLOWED_EXTENSIONS = ['png', 'jpg', 'jpeg']


def allowed_file(filename):
    return '.' in filename and \
           filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS


@app.route('/face_match', methods=['POST'])
def face_match():
    if request.method != 'POST':
        return "NOT POST"
    # check if the post request has the file part
    if ('file1' not in request.files):
        return 'No file part'

    file1 = request.files.get('file1')
    # if user does not select file, browser also submit an empty part without filename
    if file1.filename == '':
        return 'No selected file'

    if allowed_file(file1.filename):
        images = prediction.loadImages()
        # print('images loadImages == ', images)

        encodeListKnown = prediction.trainingImages(images)
        # print('encodeListKnown == ', encodeListKnown)

        faceID = prediction.testImage(file1, encodeListKnown)

        resp_data = {"match": faceID}
        return json.dumps(resp_data)
    else:
        return 'Unsupported type'


@app.route('/')
def hello_world():
    return 'Hello, World!'


# Run in HTTP
# When debug = True, code is reloaded on the fly while saved
app.run(host='0.0.0.0', port='8000', debug=True)
