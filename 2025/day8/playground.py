from collections import defaultdict

import math
import sys


def get_input(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


def main():
    # d = sqrt((x1-x2) ** 2 + (y1-y2) ** 2 + (z1-z2) ** 2)
    
    # Part 1
    data = get_input("sample.txt")

    distances = {}
    for i, row in enumerate(data):
        row = row[:-1]
        x1, y1, z1 = row.split(",")
        p1 = f"{x1}-{y1}-{z1}"
        other = {}
        for row2 in data:
            row2 = row2[:-1]
            x2, y2, z2 = row2.split(",")
            p2 = f"{x2}-{y2}-{z2}"
            if p1 == p2:
                continue
            other[p2] = math.sqrt((int(x1)-int(x2)) ** 2 + (int(y1)-int(y2)) ** 2 + (int(z1)-int(z2)) ** 2)
        distances[p1] = other
    
    closest = {}
    for k, v in distances.items():
        c = sorted(v, key=v.get)[0]
        cl = {}
        cl["point"] = c
        cl["distance"] = v[c]
        closest[k] = cl
    
    closest = dict(sorted(closest.items(), key=lambda item: item[1]['distance']))
    
    max_connections = 10
    connections_made = 0
    connected = defaultdict(list) # simple mapping of direct connections
    connections = defaultdict(set) # full mapping of connections and circuits
    connections["remaining"] = list(closest.keys())
    while True:
        connections_made += 1
        if connections_made > max_connections or len(connections["remaining"]) <= 0:
            break
        to_connect = connections["remaining"].pop(0)
        to_connect_to = closest[to_connect]["point"]
        # if to_connect_to in connections["remaining"]:
        #     connections["remaining"].remove(to_connect_to)
        if to_connect in connected.get(to_connect_to, []) or to_connect_to in connected.get(to_connect, []):
            continue
        connected[to_connect].append(to_connect_to)
        connected[to_connect_to].append(to_connect)
        circuit = to_connect # default if not in a circuit
        for k, v in connections.items():
            if k == "remaining":
                continue
            if to_connect in v or to_connect_to in v:
                circuit = k
                break
        connections[circuit].add(to_connect)
        connections[circuit].add(to_connect_to)

    # print(connections_made)
    print("#### CLOSEST ####")
    for k,v in closest.items():
        print(k,v)

    # print("#### CONNECTED ####")
    # for k,v in connected.items():
    #     print(k, v)

    # print("#### CONNECTIONS ####")
    # for k, v in connections.items():
    #     print(k, v)
    


if __name__ == '__main__':
    main()
