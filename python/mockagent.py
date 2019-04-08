#!/usr/bin/env python
# -*- coding: utf-8 -*-
#
#  mockagent.py
#  
#  Copyright 2019  <tval@va>
#  
#  This program is free software; you can redistribute it and/or modify
#  it under the terms of the GNU General Public License as published by
#  the Free Software Foundation; either version 2 of the License, or
#  (at your option) any later version.
#  
#  This program is distributed in the hope that it will be useful,
#  but WITHOUT ANY WARRANTY; without even the implied warranty of
#  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#  GNU General Public License for more details.
#  
#  You should have received a copy of the GNU General Public License
#  along with this program; if not, write to the Free Software
#  Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
#  MA 02110-1301, USA.
#  
#  

#mock agent
#generating data

import random
import time
import requests
from vehicle import Vehicle
from ThreadManager import ThreadManager

import signal
import sys
def signal_handler(sig, frame):
	print "Signal received, stopping"
	ThreadManager.terminate()
	sys.exit(42)

def main(args):
	vehicle = Vehicle("Opel", "Corsa")
	vehicle.engine_on()
	signal.signal(signal.SIGINT, signal_handler)
	signal.signal(signal.SIGTERM, signal_handler)
	print('Press Ctrl+C to exit')
	ThreadManager.waitForFinish()
	signal.pause()
	return 0
	
if __name__ == '__main__':
    import sys
    sys.exit(main(sys.argv))

