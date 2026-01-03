import math as mt

# TODO
def floyd_warshall(graph : tuple[set, list[tuple[str, str, int]]]) -> dict[str, dict[str, float]]:
    """
    :param graph: example:
    graph = (
        {'A', 'B', 'C'}, - set of vertexes
        [('A', 'B', 2), - the list of tuple (start, stop, length), directed
        ('B', 'A', -2),
        ('B', 'C', -5),
        ('A', 'C', 4)]
    )
    :return: the shortest paths between all vertexes
    """
    distances = {x : {y: mt.inf for y in graph[0]} for x in graph[0]}
    for vertex in distances.keys():
        graph[vertex][vertex] = 0 # diagonal
    for edge in graph[1]:
        distances[edge[0]][edge[1]] = edge[2]
    for i in range(len(distances)):
        for j in range(len(distances)):
            for k in range(len(distances)):
                pass

    return distances
