import requests

response = requests.post(
	"http://localhost:5000/predict", 
	files={"file": open("cat.jpg", "rb")}
)

print(response)