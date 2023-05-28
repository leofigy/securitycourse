#!/usr/env/bin python
import ctypes
lib = ctypes.CDLL("libdoubler.so")
print(lib.DoubleIt(9))