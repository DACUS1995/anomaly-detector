import json
from typing import List

import torch
from torchvision import models
import torchvision.transforms as transforms
from PIL import Image
from fastapi import FastAPI, HTTPException, File, UploadFile
from pydantic import BaseModel
from fastapi.middleware.cors import CORSMiddleware

from config import Config
from predict_classification import get_classification_prediction

device = torch.device(Config.device)
app = FastAPI()

origins = [
	"http://localhost:8080"
]

app.add_middleware(
	CORSMiddleware,
	allow_origins=origins,
	allow_credentials=True,
	allow_methods=["*"],
	allow_headers=["*"],
)

@app.post("/predict")
async def predict(file: UploadFile = File(...)):
	if file is None:
		raise HTTPException(status_code=404, detail="No file detected!")

	file_bytes = await file.read()
	class_id, class_name = get_classification_prediction(raw_input=file_bytes)
	return jsonify({"class_id": class_id, "class_name": class_name})

@app.post("/batchpredict")
async def batch_predict(files: List[UploadFile] = File(...)):
	if len(files) == 0:
		raise HTTPException(status_code=404, detail="No file detected!")

	results = []
	for file in files:
		file_bytes = await file.read()
		class_id, class_name = get_classification_prediction(raw_input=file_bytes)
		results.append({
			"class_id": class_id,
			"class_name": class_name
		})
	
	return jsonify(results)

@app.get("/")
async def root():
	return "Hello"


def main():
	print("running")


if __name__ == "__main__":
	main()