# Problem #2

## Description

Tom works in a warehouse. A billion (1,000,000,000) boxes are
arranged in a row. They are numbered from one to one billion (from
left to right). Some boxes contain a basketball (at most one
basketball in each box). In total, there are N basketballs.

Tom wants to organize the warehouse. He would like to see all the
basketballs arranged next to each other (they should occupy a
consistent interval of boxes). In one move, Tom can take one
basketball and move it into an empty box. What is the minimum
number of moves needed to organize the basketballs in the
warehouse?
Write a function:

func Solution(A []int) int

that, given an array A containing N integers, denotes the positions
of the basketballs (the numbers of the boxes in which they are
placed) and returns the minimum number of moves needed to
organize the basketballs in the warehouse.

For example, given: A = [6,4,1,7,10], your function should return 2
because the minimum number of moves needed to arrange all
basketballs next to each other is 2. There are several ways to do it.
For example, you could move the ball from the first box to the fifth,
and the ball from the tenth box to the eighth. You could also move
the ball from the first box to the fifth, and the ball from the tenth
box to the third instead. In any case, you need at least two moves.
Assume that:

•	N is an integer within the ränge [1 ..50,000];

•	each element of array A is an integer within the ränge

[1..1,000,000,000];

•	numbers in array A are pairwise distinct.

Complexity:

•	expected worst-case time complexity is 0(N*log(N));

•	expected worst-case space complexity is 0(1),
beyond input storage (not counting the storage
required for input arguments).

Elements of input arrays can be modified.