// RatInMaze.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <vector>

bool GetCheese(const std::vector<std::vector<int>>& maze, int row, int col)
{
    if (row == maze.size() - 1 && col == maze[0].size() - 1)
        return true;

    if (row == maze.size() || col == maze[0].size())
        return false;

    if (!maze[row][col])
        return false;
    // Go down and right
    return GetCheese(maze, row + 1, col) || GetCheese(maze, row, col + 1);
}

bool CanGetCheese(std::vector<std::vector<int>>& maze)
{
    return GetCheese(maze, 0, 0);
}

int main()
{
    std::vector<std::vector<int>> maze{ {1, 1, 1, 0},
        {0, 0, 1, 0},
        {0, 0, 1, 1} };//,
        //{0, 0, 1, 1} };
    std::cout << (CanGetCheese(maze) ? "True" : "False") << std::endl;
}
