from typing import List


def main():
    input_data = get_input_data_from_file("day2.input.txt")
    good_passwords = []
    good_passwords_second_half = []
    for element in input_data:
        if is_good_password(element):
            good_passwords.append(element)
        if is_good_password_two(element):
            good_passwords_second_half.append(element)

    print(len(good_passwords))
    print(len(good_passwords_second_half))

def get_input_data_from_file(filename: str) -> List[dict]:
    input_file = open(filename, "r")
    input_data = []
    for line in input_file:
        splited_line = line.split()
        min_occurences = int(splited_line[0].split('-')[0])
        max_occurences = int(splited_line[0].split('-')[1])
        character = splited_line[1].replace(':', '')
        password = splited_line[2]
        element = {
            'min': min_occurences,
            'max': max_occurences,
            'character': character,
            'password': password,
        }
        input_data.append(element)
    return input_data


def is_good_password(element: dict) -> bool:
    ocurrences = element['password'].count(element['character'])
    if ocurrences >= element['min'] and ocurrences <= element['max']:
        return True
    return False


def is_good_password_two(element: dict) -> bool:
    password = element['password']
    min_index = element['min']-1
    max_index = element['max']-1

    character = element['character']
    if (password[min_index] == character and password[max_index] != character) or \
        (password[min_index] != character and password[max_index] == character):
        return True
    return False


main()