# Problem #1

## Description

There is a soccer match between team A and team B. Each team
tries to score more goals than their opponents.

N goals are scored in the match. You are given a string S with N
characters. Each character is either 'A' or 'B'. The i-th goal is scored
by a team S[i]. For example, for a string S = "abbaaa” the match
progresses as follows:

•	Initially, there is a draw condition, 0:0.

•	Team A scores, so they are winning 1:0.

•	Team B scores, so the result is now 1:1 (another
draw).

•	Team B scores, so B is winning 2:1.

•	Team A scores, so the Status is again a draw, 2:2.

•	Team A scores, so A is now winning 3:2.

•	Team A scores, so Afmallywins 4:2.

You cheer team A on, and you celebrate each time team A takes
the lead (i.e. team A scores a goal after a draw). Count the number
of times you celebrate.
Write a function:

func Solution(S string) int

that, given a string S with N characters (representing all goals in
the match), returns the number of times team A takes the lead.

For example, for a string S = "abbaaa" (analysed above) your
function should return 2. Team A takes the lead after scoring the
very first goal (when the score becomes 1:0), and they also take the
lead when the score becomes 3:2.

Given a string S = "babbaa", your function should return 0 since
team A doesn't take the lead even once.

Assume that:

•	N is an integer within the ränge [1 ..50];

•	Each character in S is 'A' or 'B'.

In your solution, focus on correctness. The performance of your
solution will not be the focus of the assessment.