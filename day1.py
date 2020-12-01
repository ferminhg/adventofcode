import itertools
import numpy
from typing import List, Tuple


# Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

# For example, suppose your expense report contained the following:

# 1721
# 979
# 366
# 299
# 675
# 1456
# In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying t
# hem together produces 1721 * 299 = 514579, so the correct answer is 514579.

# Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you multiply them together?


SUM_OF_ITEMS = 2020

class NotFoundException(Exception):
    pass


def main():
    input_data = get_input_data_from_file("day1.input.txt")
    for num in [2, 3]:
        result = get_valid_new_year_items(input_data,num)
        print(result, numpy.prod(result))


def get_input_data_from_file(filename: str) -> List[int]:
    input_file = open(filename, "r")
    input_data = []
    for line in input_file:
        input_data.append(int(line))
    return input_data


def get_valid_new_year_items(input_data: List[int], items_to_validate: int) -> List[int]:
    combinations = itertools.combinations(range(len(input_data)), items_to_validate)
    for items in combinations:
        items_to_add = get_items_to_check(input_data, items_to_validate, items)
        if check_items_are_valid(items_to_add):
            return items_to_add
    raise NotFoundException()


def get_items_to_check(input_data: List[int], items_to_validate: int, items: Tuple) -> List[int]:
    items_to_add = []
    for pos in range(items_to_validate):
        items_to_add.append(input_data[items[pos]])
    return items_to_add


def check_items_are_valid(items_to_add: List[int]) -> bool:
    return sum(items_to_add) == SUM_OF_ITEMS


main()