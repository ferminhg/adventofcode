from typing import List


def main():
    input_data = get_input_data_from_file("day3.input.txt")
    slopes = [[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]]
    result = 1
    for slope in slopes:
        num_trees = calculate_num_of_threes(input_data, slope[0], slope[1])
        result *= num_trees
    print(result)

def get_input_data_from_file(filename: str) -> List[str]:
    input_file = open(filename, "r")
    input_data = []
    for line in input_file:
        input_data.append(line.replace('\n', ''))
    return input_data


def calculate_num_of_threes(data: List[str], inc_col:int, inc_row: int) -> int:
    num_trees = pos_col = pos_row = 0

    row_limit = len(data)-1

    while pos_row <= row_limit:
        if (get_value_from_data(data, pos_row, pos_col) == '#'):  num_trees += 1
        pos_col += inc_col
        pos_row += inc_row

    return num_trees

def get_value_from_data(data: List[str], pos_row:int, pos_col:int ) -> str:
    max_col = len(data[pos_row])
    if pos_col >= max_col:
        pos_col = pos_col % max_col

    return data[pos_row][pos_col]


main()