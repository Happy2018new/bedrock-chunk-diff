import ctypes

from python.package.internal.types import CString, as_c_string

LIB = ctypes.cdll.LoadLibrary("./c_api_test.so")

LIB.DO.argtypes = []
LIB.DO2.argtypes = []
LIB.DO.restype = None
LIB.DO2.restype = None

LIB.DO3.argtypes = [CString]
LIB.DO3.restype = None

# LIB.DO()
# print("OK")

LIB.DO3(as_c_string("ssss"))
print("OK2")
