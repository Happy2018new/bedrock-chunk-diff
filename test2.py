import ctypes
import time

LIB = ctypes.cdll.LoadLibrary("./c_api_test.so")

LIB.DO.argtypes = []
LIB.DO2.argtypes = []
LIB.DO.restype = None
LIB.DO2.restype = None

LIB.DO()
time.sleep(1)
print("OK")

LIB.DO2()
time.sleep(1)
print("OK2")