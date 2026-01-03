import pytest


@pytest.fixture
def small_graph_positive_undirected():
    return (
        {'A', 'B', 'C', 'D'},
        [('A', 'B', 2),
         ('B', 'A', 2),
         ('A', 'C', 4),
         ('C', 'A', 4),
         ('B', 'C', 10),
         ('C', 'B', 10),
         ('B', 'D', 1),
         ('D', 'B', 1),
         ('C', 'D', 5),
         ('D', 'C', 5)]
    )

@pytest.fixture
def small_graph_negative_undirected():
    return (
        {'A', 'B', 'C', 'D'},
        [('A', 'B', 2),
         ('B', 'A', 2),
         ('A', 'C', 4),
         ('C', 'A', 4),
         ('B', 'C', 10),
         ('C', 'B', -1),
         ('B', 'D', 2),
         ('D', 'B', 2),
         ('C', 'D', -1)]
    )

@pytest.fixture
def small_graph_positive_directed():
    return (
        {'A', 'B', 'C', 'D'},
        [('A', 'B', 2),
         ('A', 'C', 4),
         ('B', 'C', 10),
         ('B', 'D', 1),
         ('C', 'D', 5)]
    )

@pytest.fixture
def average_graph_positive_mixed():
    return (
        {'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'},
        [('A', 'C', 4),
         ('C', 'A', 4),
         ('B', 'C', 7),
         ('C', 'B', 7),
         ('B', 'D', 9),
         ('D', 'B', 9),
         ('D', 'C', 3),
         ('D', 'E', 2),
         ('E', 'D', 2),
         ('E', 'F', 2),
         ('F', 'E', 2),
         ('A', 'F', 4),
         ('F', 'A', 3),
         ('G', 'G', 9),
         ('H', 'G', 30),
         ('H', 'G', 20),
         ('G', 'H', 50)]
    )