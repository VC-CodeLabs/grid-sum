🌟 Go Grid Challenge: The Max Sum Quest 🌟
Embark on an exhilarating journey through a world of numbers! In this Go Grid Challenge, your mission is to navigate a square grid and unearth the path that boasts the highest sum. But remember, your adventure is bound by the laws of the grid – you may only traverse a single row or column. Speed and efficiency are your allies in this quest, as you strive for an algorithmic solution with O(n) complexity.

🧭 The Challenge
You are presented with a grid, a square of numbers, each holding secrets to untold riches. Your challenge is to devise a Go program that can uncover the highest sum possible within the confines of a single row or column.

The Grid's Law
Square Grid: The grid you will encounter is always a square, ranging in size from 2x2 to 50x50 cells.
Single Path: Your journey must be confined to a single row or column of the grid.
Optimize Your Journey: Seek efficiency in your algorithm. Aim for a time complexity of O(n).

🏁 Getting Started
Clone the Repository: Your first step on this adventure is to clone this repository.
Study the Grid: Remember, the grid is a square matrix of integers, varying in size from 2x2 to 50x50.
Implement Your Strategy: Craft your solution in a file named **[name]_solution.go**.
Test for Efficiency: Ensure your solution not only passes the test cases but also operates efficiently; as fast as possible, try to get down to O(n).

📜 Submission Guidelines
To ensure a fair and organized quest:
Naming Your Solution
Unique File Names: Title your solution file as **[name]_solution.go**.
  - Repalace [name] with your name
Respect the Code: Do not alter the original template or other participants' submissions.

Submitting Your Work
Fork and Contribute: Fork this repository, add your unique solution.

🛡️ Code of Honor
To uphold the integrity of the challenge:

Review Process: Each submission will be reviewed for compliance with the guidelines and performance.

⚖️ Evaluation Criteria
Correctness: Does your algorithm find the highest sum accurately?
Efficiency: Does your solution run faster than other submissions?

📊 Leaderboard
Heroes who demonstrate exceptional skill in both correctness and efficiency will be celebrated on our leaderboard. Will you rise to the top?


SUPPORT
- You may run into an issue using a VC machine when trying to run the Go application locally.  To get around the issue you can follow the following steps.

- Create a TMP directory in your VC users folder under projects
- Add an environment variable to tell Go to use the specified TMP directory

- To specify a specific temporary directory for Go to use when building and running your program, you can set the GOTMPDIR environment variable. This variable allows you to override the default temporary directory.

In PowerShell, you can set the GOTMPDIR environment variable before running your Go program. Here's an example:

powershell
Copy code
$env:GOTMPDIR = "C:\path\to\your\desired\temp\directory"
go run .\ababich_solution.go
Replace "C:\path\to\your\desired\temp\directory" with the actual path you want to use as the temporary directory.

This will instruct Go to use the specified directory for temporary files during the build and execution process. Keep in mind that you need to set this environment variable each time you run your Go program, or you can set it globally for your user account if you want it to apply to all Go projects.
