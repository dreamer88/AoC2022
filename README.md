My solutions for AoC2022 using GoLang.

# Thoughts on Go
This was my first time using Go so many of the solutions and techniques could definitely be improved. This submitted solutions are not always my original implementations; I've gone back and cleaned up many solutions as I learned more about the language. One item in particular that would have made my life easier from the beginning was realizing that I could use `map[Point2d]type` instead of using `map[int]map[int]type` which added a lot of extra compexity to the 2D problems! Life was also much easier after I wrote a `HashSet[T]` implementation; since Go doesn't provide a native `set` implementation creating a basic implementation wrapped around a `map[T]bool` was incredibly helpful. I also had several issues with how Go handles string types and iterators; the fact that `string[i]` returns a `byte` but `for _, c := range string` returns a `rune` was fairly annoying. Everything being copy-by-value was also a buggaboo on a few problems.

# Problem notes
## Problems I particularly enjoyed:
- Day 6: This made creating the `HashSet` class worth it and was more validating than anything.
- Day 10: I always like the puzzles that make actual visuals :)
- Day 11: Brushing up on modulo arithmetic is always nice.
- Day 16: This one I struggled with a little more than I wanted to. Part one was straightforward; I could just jump directly to openable valves and calculate times directly. Trying to expand that to Part Two though, I had a lot of mis-steps forgetting about the 1-turn-delay when opening a valve. Adding a proper cache/early out system so that it didn't take forever also took some configuring and while this could still be faster, I'm happy enough with the overall approach.
- Day 18: I really enjoyed part 2; I liked the mental visual of "filling the object with water".
- Day 21 Part 2: I'm mostly just happy that the problem didn't have `humn` in multiple places; trying to solve a complex set of equations would have made my head hurt.

## Problems I had particular trouble with:
- Day 5: Parsing the input was just generally annoying and had a few problems. This was also the first instance of modifying arrays passed by value causing me issues.
- Day 9: Attempting puzzles _after_ a few cocktails is much harder than attempting them before cocktails. After much consternation (and sleep), my issue was just a bad regex which worked on the sample (`\d`) but not on the real input (`\d+`) which was incredibly frustrating!
- Day 13: I found this question very awkwardly worded and was returning bad values when encountering two lists; I was returning out early if the left list was shorter than the right list, not if `first[i] < second[i]`.
- Day 17 Part Two: I was being dumb and detected a loop as "the first round where I enouncter a rock `X` with wind `Y` that I have seen before". The issue with that approach is that the rock may settle on different rows in that scenario, causing your sample input to succeed and your real input to fail. I ended up adding a dirty check to cache the previous `N` rows up until we have encountered a solid object in each column. This would break if you were playing tetris and were leaving a column empty for the long piece but fortunately that doesn't happen with this input.
- Day 19: Programming the solution was easy. Finding ways to effectively early out so it didn't take forever was less easy and definitely aren't the types of problems I enjoy.
- Day 20: I envy everyone using Python that could just `rotate`
- Day 22 Part 2: Any time I need to take a piece of paper and cut out the shape to visualize what I'm doing... No thanks :D

# Leaderboard
Using a new language, I don't target hitting the global leaderboard but was close a few times towards the end! December was a busy month so I didn't get to try every puzzle immediately but I'm glad I was able to polish all of them off in time to finish on Day 25 :)

```
      --------Part 1--------   --------Part 2--------
Day       Time   Rank  Score       Time   Rank  Score
 25   01:38:04   3320      0   01:38:20   2595      0
 24   00:51:03    839      0   01:00:34    828      0
 23   01:09:14   1859      0   01:10:11   1626      0
 22   23:19:35  12854      0       >24h   7531      0
 21   00:20:11   2101      0   00:56:44   1697      0
 20   03:10:20   4058      0   03:30:30   3739      0
 19       >24h   9228      0       >24h   8160      0
 18   02:01:31   5269      0   02:44:35   3921      0
 17   18:35:30  13260      0       >24h   9926      0
 16   03:55:22   3541      0       >24h  13767      0
 15   13:07:47  21558      0   20:50:16  21173      0
 14   00:32:53   2099      0   00:35:22   1572      0
 13   00:56:45   3998      0   01:14:23   4076      0
 12   00:24:47   1525      0   00:26:52   1252      0
 11   00:53:15   5023      0   00:56:30   2375      0
 10   03:28:22  17585      0   03:44:56  15091      0
  9       >24h  55236      0       >24h  46330      0
  8   00:20:26   3283      0   00:34:52   2874      0
  7   00:32:23   2331      0   00:44:58   2834      0
  6   00:08:01   4653      0   00:10:13   4694      0
  5   00:58:12  11081      0   01:01:23  10104      0
  4   00:16:59   7936      0   00:19:04   6384      0
  3   00:17:09   6117      0   00:23:14   4713      0
  2   00:10:56   3449      0   00:18:15   3841      0
  1   00:06:43   3763      0   00:12:47   4776      0
  ```