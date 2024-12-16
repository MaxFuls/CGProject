
from sympy import Matrix, symbols, lcm
import re
import sys

def parse_compound(compound):
    """Парсит химическое соединение и возвращает словарь с элементами и их количеством."""
    elements = re.findall(r'([A-Z][a-z]*)(\d*)', compound)
    parsed = {}
    for element, count in elements:
        parsed[element] = parsed.get(element, 0) + int(count) if count else 1
    return parsed

def build_matrix(equation):
    """Строит матрицу для балансировки уравнения."""
    reactants, products = equation.split('=')
    reactants = reactants.split('+')
    products = products.split('+')

    compounds = reactants + products
    all_elements = set()
    for compound in compounds:
        all_elements.update(parse_compound(compound).keys())

    element_list = sorted(all_elements)
    matrix = []

    for element in element_list:
        row = []
        for compound in reactants:
            row.append(parse_compound(compound).get(element, 0))
        for compound in products:
            row.append(-parse_compound(compound).get(element, 0))
        matrix.append(row)

    return Matrix(matrix), len(reactants)

def balance_chemical_equation(equation):
    """Балансирует химическое уравнение."""
    matrix, split_index = build_matrix(equation)
    symbols_list = symbols(f'x0:{matrix.shape[1]}')

    # Решаем уравнение
    solution = matrix.nullspace()[0]
    lcm_value = lcm([val.q for val in solution])  # Находим НОК для приведения к целым числам
    coefficients = [int(val * lcm_value) for val in solution]

    # Формируем итоговое уравнение
    reactants, products = equation.split('=')
    reactants = reactants.split('+')
    products = products.split('+')

    reactant_side = ' + '.join(f'{coefficients[i]} {reactants[i].strip()}' for i in range(split_index))
    product_side = ' + '.join(f'{coefficients[i + split_index]} {products[i].strip()}' for i in range(len(products)))

    return f'{reactant_side} = {product_side}'

# Пример использования
# requestedData = str(input())

balanced = balance_chemical_equation(sys.argv[1])
print(balanced)