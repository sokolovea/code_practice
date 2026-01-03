import pytest
from src.graph.dijkstra import dijkstra


@pytest.mark.parametrize("start, end, expected", [
    ('A', 'D', 3),
    ('B', 'C', 6),
    ('C', 'D', 5),
    ('B', 'D', 1),
    ])
def test_dijkstra_small_graph_positive_undirected(start: str, end: str, expected: float, small_graph_positive_undirected):
    assert dijkstra(small_graph_positive_undirected, start, end) == expected

@pytest.mark.parametrize("start, end, expected", [
    ('A', 'D', 3),
    ('D', 'A', float('inf')),
    ('B', 'C', 10),
    ('C', 'B', float('inf')),
    ('C', 'D', 5),
    ('D', 'C', float('inf')),
    ('B', 'D', 1),
    ('D', 'B', float('inf')),
    ])
def test_dijkstra_small_graph_positive_directed(start: str, end: str, expected: float, small_graph_positive_directed):
    assert dijkstra(small_graph_positive_directed, start, end) == expected

@pytest.mark.parametrize("start, end, expected", [
    ('A', 'C', 4),
    ('A', 'F', 4),
    ('F', 'A', 3),
    ('C', 'F', 8),
    ('F', 'C', 7),
    ('A', 'G', float('inf')),
    ('F', 'H', float('inf')),
    ('C', 'E', 10),
    ('B', 'D', 9),
    ('H', 'G', 20),
    ('G', 'H', 50)
    ])
def test_dijkstra_average_graph_positive_mixed(start: str, end: str, expected: float, average_graph_positive_mixed):
    assert dijkstra(average_graph_positive_mixed, start, end) == expected

# @pytest.mark.parametrize("start, end, expected", [
#     ('A', 'B', 2),
#     ('A', 'C', 4),
#     ('B', 'A', 2),
#     ('C', 'B', -1),
#     ('B', 'C', 6),
#     ('B', 'D', 2),
#     ('D', 'A', 4),
#     ('A', 'D', 3)
#     ])
# def test_dijkstra_small_graph_negative_undirected(start: str, end: str, expected: float, small_graph_negative_undirected):
#     """
#     Not for Dijkstra
#     """
#     assert dijkstra(small_graph_negative_undirected, start, end) == expected

if __name__ == '__main__':
    pytest.main()