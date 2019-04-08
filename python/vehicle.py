import threading
import time
import requests
import random
from ThreadManager import ThreadManager
import logging
import protoCar_pb2
import json

class Vehicle:
	def __init__(self, make, model):
		self.make = make
		self.model = model
		self.running = 0
		self.eventGenerator = False
		self.readings = []
		self.message = None
	
	def engine_status(self):
		return self.running
	
	def engine_on(self):
		self.running = 1
		#self.eventGenerator = threading.Thread(target=generateEvents, args=(self,))
		self.eventGenerator = EventGenerator(self)
		ThreadManager.subscribe(self.eventGenerator)
		self.eventGenerator.start()
	
	def engine_off(self):
		ThreadManager.unsubscribe(self.eventGenerator)
		self.running = 0

class EventGenerator(threading.Thread):
	def __init__(self,instance):
		threading.Thread.__init__(self)
		self.car = instance
		self._stopevent = threading.Event()
		self._sleepperiod = 0.05

	def run(self):
		self.car.message = protoCar_pb2.CarMessage()
		self.car.message.make = self.car.make
		self.car.message.model = self.car.model
		self.car.message.color="RED"
		self.car.message.year=2008
		message = self.car.message.SerializeToString()
		#message = {
		#	"make":"Opel",
		#	"model":"Corsa",
		#	"year":2008,
		#	"color":"RED",
		#}
		#message = json.dumps(message)
		message = {
			"make":"Opel",
			"model":"Corsa",
			"year":2008,
			"color":"RED",
		}
		count = 0
		num_requests = 5000
		while count < num_requests:
			if self._stopevent.isSet():
				break
			if self.car.engine_status():
				#data = {
				#	"message" : message,
				#}
				#headers = {
				#	"Content-Type": "application/octet-stream",
				#}
				try:
					r = requests.post("http://localhost:8000/requestbin", data = message)
					print r.text
				except:
					continue
				count += 1
			self._stopevent.wait(self._sleepperiod)
	
	def join(self, timeout=None):
		threading.Thread.join(self, timeout)

	def exitThread(self, timeout=None):
		self._stopevent.set()
		threading.Thread.join(self,timeout)

class VehicleReading:
	def __init__(self, name, upperLimit, lowerLimit):
		self.tendency = 1
		self.name = name
		self.upperLimit = upperLimit
		self.lowerLimit = lowerLimit
		self.value = 0.0

def generateEvents(instance):
	while instance.engine_status():
		time.sleep(0.5)
		for attr in instance.attributes:
			genData(attr)	
		#r = requests.get("http://localhost:8000/"+instance.make+"/"+instance.model)

def genData(attr):
	chance = random.uniform(0.0,100.0)
	if chance < attr[1]:
		return random.uniform(0.0,100.0)
	return None
