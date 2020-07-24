from sklearn.tree import DecisionTreeClassifier

class ModelHandler:
	__model = None
	@staticmethod 
	def get_model(name):
		if ModelHandler.__model == None:
			ModelHandler(name)
		return ModelHandler.__model

	# TODO Add support for multiple model types 
	def __init__(self, name):
		""" Virtually private constructor. """
		if ModelHandler.__model != None:
			raise Exception("This class is a singleton!")
		else:
			# Need implementation
			model = None
			model.to(device)
			model.eval()
			ModelHandler.__model = model
			