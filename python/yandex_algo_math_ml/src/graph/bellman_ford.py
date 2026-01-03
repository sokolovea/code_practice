import math as mt
from queue import PriorityQueue

def bellman_ford(graph : tuple[set, list[tuple[str, str, int]]], start: str, end: str) -> float:
    """
    :param graph: example:
    graph = (
        {'A', 'B', 'C'}, - set of vertexes
        [('A', 'B', 2), - the list of tuple (start, stop, length), directed
        ('B', 'A', -2),
        ('B', 'C', -5),
        ('A', 'C', 4)]
    )
    :param start: start path vertex
    :param end: end path vertex
    :return: the shortest path from start to end
    """
    distances = {vertex: mt.inf for vertex in graph[0]}
    distances[start] = 0
    priority_queue = PriorityQueue()
    priority_queue.put((distances[start], start))
    passed_vertices = set()

    while not priority_queue.empty():
        # print(distances)
        item = priority_queue.get()
        vertex = item[1]
        passed_vertices.add(vertex)
        for edge in graph[1]:
            if edge[0] == vertex:
                if distances[edge[0]] + edge[2] < distances[edge[1]]:
                    distances[edge[1]] = distances[edge[0]] + edge[2]
                    priority_queue.put((distances[edge[1]], edge[1]))
    return distances[end]
