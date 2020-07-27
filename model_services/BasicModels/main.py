import json
from typing import List, Optional

from fastapi import FastAPI, HTTPException, File, UploadFile
from pydantic import BaseModel
from fastapi.middleware.cors import CORSMiddleware

from config import Config
from predict_classification import get_classification_prediction

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

class Data(BaseModel):
	values: List[float]

class BatchData(BaseModel):
	batch: List[Data]

class Result(BaseModel):
	class_id: str
	class_name: str

class BatchResults(BaseModel):
	results: Result


@app.post("/detect", response_model=Result)
async def predict(data: Data):
	if len(data.values) == 0:
		raise HTTPException(status_code=404, detail="Empty data point sent!")

	class_id, class_name = get_classification_prediction(input=data.values)
	return Result(class_id=class_id[0], class_name=class_name[0])


# TODO Optimize this using vectorization based batch processing
@app.post("/detect/batch", response_model=BatchResults)
async def batch_predict(batchData: BatchData):
	if len(batchData.values) == 0:
		raise HTTPException(status_code=404, detail="Empty data batch point sent!")

	results = []
	for data in batchData.batch:
		file_bytes = await file.read()
		class_id, class_name = get_classification_prediction(raw_input=file_bytes)
		results.append({
			"class_id": class_id,
			"class_name": class_name
		})
	
	return BatchResults(results=results)

@app.get("/")
async def root():
	return "Hello"


def main():
	print(f"Running service: {Config.service_name}")


if __name__ == "__main__":
	main()