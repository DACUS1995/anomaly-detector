class ModelHandler:
	__model = None
	@staticmethod 
	def get_model():
		if ModelHandler.__model == None:
			ModelHandler()
		return ModelHandler.__model

	def __init__(self):
		""" Virtually private constructor. """
		if ModelHandler.__model != None:
			raise Exception("This class is a singleton!")
		else:
			# Need implementation
			model = None
			model.to(device)
			model.eval()
			ModelHandler.__model = model
			