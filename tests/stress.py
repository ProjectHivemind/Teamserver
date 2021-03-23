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


def get_random_string():
    return str(''.join(choice(chars) for _ in range(random.randint(5, 15))))


def get_random_digits():
    return str(''.join(choice(digits) for _ in range(random.randint(5, 15))))


def gen_module():
    num_funcs = random.randint(1, 9)
    num_params = random.randint(1, 5)

    random_mod_name = random.choice(mod_list)
    mod_list.remove(random_mod_name)
    temp_mod_func_list = copy.deepcopy(mod_func_list)
    functions = []
    for i in range(num_funcs):
        random_func_name = random.choice(temp_mod_func_list)
        temp_mod_func_list.remove(random_func_name)
        functions.append({
            "moduleFuncDesc": f"{random_func_name} Description",
            "moduleFuncName": f"{random_func_name}",
            "paramNames": random.sample(mod_param_list, num_params),
            "paramNum": num_params,
            "paramTypes": [choice(mod_param_types_list) for _ in range(num_params)]
        })
    return {
        "moduleDesc": f"{random_mod_name} Description",
        "moduleFuncs": functions,
        "moduleName": f"{random_mod_name}"
    }


def get_random_registration_data():
    num = random.randint(1, 10)

    return {
        "IP": socket.inet_ntoa(struct.pack('>I', random.randint(1, 0xffffffff))),
        "MAC": "02:00:00:%02x:%02x:%02x" % (random.randint(0, 255), random.randint(0, 255), random.randint(0, 255)),
        "OS": ''.join(choice(chars) for _ in range(random.randint(5, 15))),
        "hostname": f"{get_random_string()}",
        "implantName": f"{get_random_string()}",
        "implantVersion": f"{get_random_digits()}",
        "otherIPs": [socket.inet_ntoa(struct.pack('>I', random.randint(1, 0xffffffff))) for _ in
                     range(random.randint(0, 3))],
        "supportedModules": [choice(created_modules) for _ in range(num)]
    }


def generate_packet(data, packet_type, uuid=''):
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


def send_data(data, host, port):
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    is_open = sock.connect_ex((host, port))
    while is_open != 0:
        time.sleep(250)
        sock.connect_ex((host, port))
    sock.send(json.dumps(data).encode('utf-8'))
    data_received = sock.recv(1024)
    sock.close()

    if data_received.strip() != b'':
        info = json.loads(data_received)
        return info
    return None


def replicate_bot():
    port = 1234
    host = 'localhost'
    uuid = ''

    p = generate_packet(json.dumps(get_random_registration_data()), REGISTRATION)

    info = send_data(p, host, port)

    if info is not None:
        uuid = info[0]['implantInfo']['UUID']
    else:
        print('ERROR WITH SERVER')
        exit()

    for _ in range(1000):
        p = generate_packet('', REQUEST_ACTIONS, uuid)
        time.sleep(5)
        info = send_data(p, host, port)


class MyThread(threading.Thread):
    def __init__(self):
        threading.Thread.__init__(self)

    def run(self):
        replicate_bot()


if __name__ == '__main__':
    threads = []

    created_modules = [gen_module() for _ in mod_list]

    for _ in range(1000):
        temp = MyThread()
        threads.append(temp)
        temp.start()

    for t in threads:
        t.join()
