import ctypes
import time

CPtr = ctypes.c_void_p
CSlice = CPtr
CString = CPtr
CInt = ctypes.c_int
CLongLong = ctypes.c_longlong

LIB1 = ctypes.cdll.LoadLibrary("./c_api_test.so")
LIB2 = ctypes.cdll.LoadLibrary("./c_api_2.so")


LIB1.DO.argtypes = []
LIB1.DO.restype = None
LIB2.DO.argtypes = []
LIB2.DO.restype = None

LIB1.DO()
LIB2.DO()
time.sleep(1)
print("OK")
