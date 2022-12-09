import sys


# load data stream from input file
def import_data_stream(file):
    with open(file) as f:
        line = f.readline()
    return line


# search string for {length} consequtive unique characters
# return index position of last character in the sequence
def find_marker(data_stream, length):
    marker = []
    character_position = 0
    for char in data_stream:
        character_position += 1
        marker.append(char)
        if len(marker) == length:
            if len(set(marker)) == length:
                return character_position
            else:
                marker = marker[1:]
    sys.exit(1)


def main():
    data_stream = import_data_stream("input.txt")
    data_stream.strip()
    start_of_packet_marker = find_marker(data_stream, 4)
    print(f"Start of packet marker (part 1): {start_of_packet_marker}")
    start_of_message_marker = find_marker(data_stream, 14)
    print(f"Start of message marker (part 2): {start_of_message_marker}")


if __name__ == '__main__':
    main()