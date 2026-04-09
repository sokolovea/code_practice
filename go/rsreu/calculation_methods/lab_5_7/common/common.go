package common

/*
Найти точку минимума функции одной переменной, определенной вариантом задания,
методами перебора и поразрядного поиска с заданной погрешностью. Для каждого
метода вывести точку минимума, значение функции в точке минимума, число
вычислений значения функции в процессе минимизации. Проверку результата
выполнить с помощью встроенных к Matlab (Scilab) функций минимизации при
погрешности, в 100 раз меньшей заданной.
*/

type F func(float64) float64

type MinimizeResult struct {
	XMin  float64
	YMin  float64
	Iters uint64
}

type Minimize func(left float64, right float64, HandlingFunc F, eps float64) MinimizeResult
