from pykeepass import PyKeePass

kp = PyKeePass("./example.kdbx", password="abcdefg12345678")
response = kp.find_entries(title='foo', first=True)
print("without deref -> "+response.password)
print("with deref -> "+kp.deref(response.password))