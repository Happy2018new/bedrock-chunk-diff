import ctypes
import time

LIB = ctypes.cdll.LoadLibrary("./c_api_test.so")

LIB.DO.argtypes = []
LIB.DO.restype = None

LIB.DO()
time.sleep(1)
print("OK")