from random import choice
import random
from string import ascii_lowercase, digits
import socket
import struct
import json
import time
import threading

chars = ascii_lowercase + digits

REGISTRATION = 4
REQUEST_ACTIONS = 1


def getRandomString():
    return str(''.join(choice(chars) for _ in range(random.randint(5, 15))))


def getRandomDigits():
    return str(''.join(choice(digits) for _ in range(random.randint(5, 15))))


def getRandomRegistrationData():
    num = random.randint(1, 10)
    return {
        "IP": socket.inet_ntoa(struct.pack('>I', random.randint(1, 0xffffffff))),
        "MAC": "02:00:00:%02x:%02x:%02x" % (random.randint(0, 255), random.randint(0, 255), random.randint(0, 255)),
        "OS": ''.join(choice(chars) for _ in range(random.randint(5, 15))),
        "hostname": f"{getRandomString()}",
        "implantName": f"{getRandomString()}",
        "implantVersion": f"{getRandomDigits()}",
        "otherIPs": [socket.inet_ntoa(struct.pack('>I', random.randint(1, 0xffffffff))) for _ in
                     range(random.randint(0, 3))],
        "supportedModules": [
            {
                "moduleDesc": "ModuleDesc",
                "moduleFuncs": [
                    {
                        "moduleFuncDesc": "Module Func Desc",
                        "moduleFuncName": "Module Func Name",
                        "paramNames": [''.join(choice(chars) for _ in range(random.randint(5, 15))) for _ in
                                       range(num)],
                        "paramNum": num,
                        "paramTypes": [''.join(choice(chars) for _ in range(random.randint(5, 15))) for _ in range(num)]
                    }
                ],
                "moduleName": f"{getRandomString()}"
            }
        ]
    }


def generatePacket(data, packet_type, uuid=''):
    return {
        "data": data,
        "fingerprint": "fingerprint",
        "implantInfo": {
            "UUID": uuid,
            "primaryIP": socket.inet_ntoa(struct.pack('>I', random.randint(1, 0xffffffff)))
        },
        "numLeft": 0,
        "packetType": packet_type
    }


def sendData(data, host, port):
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect((host, port))
    s.send(json.dumps(data).encode('utf-8'))
    temp = s.recv(1024)
    s.close()
    # print(f'Received from server: {temp}')
    if temp.strip() != b'':
        info = json.loads(temp)
        return info
    return None


def replicateBot():
    port = 1234
    # print(host)
    host = 'localhost'
    uuid = ''
    # time.sleep(random.randint(0, 10))
    p = generatePacket(json.dumps(getRandomRegistrationData()), REGISTRATION)
    # print(f'THIS IS THE PACKET BEING SENT: LENGTH: {len(json.dumps(p))} DATA: {p}')
    info = sendData(p, host, port)
    # print(info)
    if info is not None:
        uuid = info[0]['implantInfo']['UUID']
    else:
        print('ERROR WITH SERVER')
        exit()

    for _ in range(1000):
        p = generatePacket('', REQUEST_ACTIONS, uuid)
        time.sleep(5)
        info = sendData(p, host, port)


class myThread(threading.Thread):
    def __init__(self):
        threading.Thread.__init__(self)

    def run(self):
        replicateBot()


if __name__ == '__main__':
    threads = []

    for _ in range(500):
        temp = myThread()
        threads.append(temp)
        temp.start()

    for t in threads:
        t.join()
