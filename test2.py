import ctypes
import struct

CPtr = ctypes.c_void_p
CSlice = CPtr
CString = CPtr
CInt = ctypes.c_int
CLongLong = ctypes.c_longlong

LIB = ctypes.cdll.LoadLibrary("./c_api_test.so")
LIB.FreeMemory.argtypes = [CPtr]
LIB.FreeMemory.restype = None


def free_memory(address: CPtr) -> None:
    LIB.FreeMemory(address)


def as_c_bytes(b: bytes) -> CSlice:
    return ctypes.cast(ctypes.c_char_p(struct.pack("<I", len(b)) + b), CSlice)


def as_python_bytes(slice: CSlice) -> bytes:
    length = struct.unpack("<I", ctypes.string_at(slice, 4))[0]
    result = ctypes.string_at(slice, 4 + length)[4:]
    free_memory(slice)
    return result


def as_c_string(string: str) -> CString:
    return ctypes.cast(ctypes.c_char_p(bytes(string, encoding="utf-8")), CString)


def as_python_string(c_string: CString) -> str:
    result = ctypes.c_char_p(c_string).value.decode(encoding="utf-8")  # type: ignore
    free_memory(c_string)
    return result


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
