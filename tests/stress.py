import copy
import json
import random
import socket
import struct
import threading
import time
from random import choice
from string import ascii_lowercase, digits

chars = ascii_lowercase + digits

REGISTRATION = 4
REQUEST_ACTIONS = 1
mod_list = ['today', 'south', 'project', 'pages', 'version', 'section', 'found', 'sports', 'house', 'related',
            'security', 'county', 'american', 'photo', 'members', 'power', 'while', 'network', 'computer', 'systems']

mod_func_list = ['three', 'total', 'place', 'download', 'without', 'access', 'think', 'north', 'current', 'posts',
                 'media', 'control', 'water', 'history', 'pictures', 'personal', 'since', 'guide', 'board', 'location']

mod_param_list = ['change', 'white', 'small', 'rating', 'children', 'during', 'return', 'students', 'shopping',
                  'account', 'times']

mod_param_types_list = ['String', 'Int', 'Double']


created_modules = []

def getRandomString():
    return str(''.join(choice(chars) for _ in range(random.randint(5, 15))))


def getRandomDigits():
    return str(''.join(choice(digits) for _ in range(random.randint(5, 15))))


def genModule():
    numFuncs = random.randint(1, 9)
    numParams = random.randint(1, 5)
    # numFuncs = 1
    # numParams = 1
    random_mod_name = random.choice(mod_list)
    mod_list.remove(random_mod_name)
    temp_mod_func_list = copy.deepcopy(mod_func_list)
    functions = []
    for i in range(numFuncs):
        random_func_name = random.choice(temp_mod_func_list)
        temp_mod_func_list.remove(random_func_name)
        functions.append({
            "moduleFuncDesc": f"{random_func_name} Description",
            "moduleFuncName": f"{random_func_name}",
            "paramNames": random.sample(mod_param_list, numParams),
            "paramNum": numParams,
            "paramTypes": [choice(mod_param_types_list) for _ in range(numParams)]
        })
    return {
        "moduleDesc": f"{random_mod_name} Description",
        "moduleFuncs": functions,
        "moduleName": f"{random_mod_name}"
    }


def getRandomRegistrationData():
    num = random.randint(1, 10)
    # num = 1

    return {
        "IP": socket.inet_ntoa(struct.pack('>I', random.randint(1, 0xffffffff))),
        "MAC": "02:00:00:%02x:%02x:%02x" % (random.randint(0, 255), random.randint(0, 255), random.randint(0, 255)),
        "OS": ''.join(choice(chars) for _ in range(random.randint(5, 15))),
        "hostname": f"{getRandomString()}",
        "implantName": f"{getRandomString()}",
        "implantVersion": f"{getRandomDigits()}",
        "otherIPs": [socket.inet_ntoa(struct.pack('>I', random.randint(1, 0xffffffff))) for _ in
                     range(random.randint(0, 3))],
        "supportedModules": [choice(created_modules) for _ in range(num)]
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

    created_modules = [genModule() for _ in mod_list]

    for _ in range(500):
        temp = myThread()
        threads.append(temp)
        temp.start()

    for t in threads:
        t.join()
