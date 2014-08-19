##Tentative Syllabus

The course is based on this [textbook](http://www.recursivebooks.com): 
Operating Systems: Principles and Practice, Tom Anderson and Mike Dahlin. 
You do not need the second edition. I only have the beta edition still.

All chapters listed below and lab exercises 1-7 are syllabus. 
I will notify of any changes that may arise.

| Chapter | Details                                                                           |
|:-------:|-----------------------------------------------------------------------------------|
| 1       | All. The concepts in 1.2 are important.                                           |
| 2       | All except 2.4 and 2.5. Sidebar on p. 93 gives a good summary.                    |
| 3       | All except Windows process management on p. 108-109.                              |
| 4       | All                                                                               |
| 5       | All except p. 221-225 (Linux 2.6 mutex lock). Sidebars on p. 211 and p. 213-215 not important for exam. *Rules listed on p. 237-239 are important.* |
| 6       | All except fig. 6.11 p. 276, fig. 6.13 p. 282, and p. 289-291 (Implementing RCU). *6.1 and 6.2: important.* |
| 7       | All. 7.1 and 7.2: important. 7.3 not important for exam.                          |
| 8       | All. 8.1-8.3: important. 8.4 not important for exam.                              |
| 9       | All important.                                                                    |
| 10-14   | Read on your own material; will not play a significant role on the exam.          |


##Lecture and Lab Plan

| W    |  Date | Chap. | Topic                                  | Teacher | Travels      |
|:----:|:-----:|:-----:|----------------------------------------|:-------:|:------------:|
|  35  |  25/8 |   1   | Introduction to Operating Systems      |  Morten |              |
|      |  26/8 |       | Introduction to C programming          |  Morten |              |
|      |  26/8 | Lab 1 | Unix, programming tools and C          |  Morten |              |
|      |  31/8 |       | **[Lab 1 Handin](https://github.com/uis-dat320-fall2014/lab1)** |         |              |
|  36  |  1/9  |       | Variables, pointers, and memory        |  Morten |              |
|      |  2/9  |   2   | The Kernel Abstraction                 |   Hein  |              |
|      |  2/9  | Lab 2 | Introduction to Go programming         |   Hein  |              |
|      |  7/9  |       | **Lab 2 Handin**                       |         |              |
|  37  |  8/9  |   3   | The Programming Interface              |   Hein  | Morten@CP    |
|      |  9/9  |   4   | Concurrency and Threads                |   Hein  |     "        |
|      |  9/9  | Lab 3 | Network Programming with Go            |   Hein  |     "        |
|      |  14/9 |       | **Lab 3 Handin**                       |         |              |
|  38  |  15/9 |   4   | Concurrency and Threads                |   Hein  |              |
|      |  16/9 |   5   | Synchronizing Access to Shared Objects |   Hein  |              |
|      |  16/9 | Lab 4 | Threads and Protection                 |  Morten |              |
|      |  21/9 |       | **Lab 4 Handin**                       |         |              |
|  39  |  22/9 |   5   | Synchronizing Access to Shared Objects |   Hein  | Morten@ICTSS |
|      |  23/9 |   5   | Synchronizing Access to Shared Objects |   Hein  |     "        |
|      |  23/9 | Lab 5 | Programming Tools                      |   Hein  |     "        |
|  40  |  29/9 |   6   | Advanced Synchronization               |   Hein  |              |
|      |  30/9 |   6   | Advanced Synchronization               |   Hein  | Morten@Oslo  |
|      |  30/9 | Lab 5 | Programming Tools                      |   Hein  |              |
|      |  5/10 |       | **Lab 5 Handin**                       |         |              |
|  41  |  6/10 |       | *No lectures*                          |  Morten | Hein@OSDI    |
|      |  7/10 |       | *No lectures*                          |  Morten |     "        |
|      |  7/10 | Lab 6 | Linux Kernel IO Driver                 |  Morten |     "        |
|  42  | 13/10 |       | *No lectures*                          |  Morten | Hein@DISC    |
|      | 14/10 |       | *No lectures*                          |  Morten |     "        |
|      | 14/10 | Lab 6 | Linux Kernel IO Driver                 |  Morten |     "        |
|  43  | 20/10 |   7   | Scheduling                             |   Hein  |              |
|      | 21/10 |   7   | Scheduling                             |   Hein  |              |
|      | 21/10 | Lab 6 | Linux Kernel IO Driver                 |  Morten |              |
|      | 26/10 |       | **Lab 6 Handin**                       |         |              |
|  44  | 27/10 |   8   | Address Translation                    |   Hein  |              |
|      | 28/10 |   8   | Address Translation                    |   Hein  |              |
|      | 28/10 | Lab 7 | ChanStat: TV channel statistics        |   Hein  |              |
|  45  |  3/11 |   9   | Caching and Virtual Memory             |   Hein  |              |
|      |  4/11 |   9   | Caching and Virtual Memory             |   Hein  |              |
|      |  4/11 | Lab 7 | ChanStat: TV channel statistics        |   Hein  |              |
|  46  | 10/11 |       | *No lectures*                          |   Hein  |              |
|      | 11/11 |       | *No lectures*                          |   Hein  |              |
|      | 11/11 | Lab 7 | ChanStat: TV channel statistics        |   Hein  |              |
|      | 16/11 |       | **Lab 7 Handin**                       |         |              |
|  47  | 21/11 |       | Lab Handin (w/5 slip days)             |         |              |
|  48  | 24/11 |       | Final Handin Date (w/reduced grade)    |         |              |
|  50  | 13/12 |       | **Exam**                               |         |              |


##Lab Overview

| Lab    | Topic                           | Grading   | Submission         | Deadline | 
|:------:|---------------------------------|-----------|--------------------|:--------:|
| 1      | Unix, programming tools and C   | Pass/Fail | Submit on your own | 31/8     |
| 2      | Introduction to Go programming  | Pass/Fail | Submit on your own | 7/9      |
| 3      | Network Programming with Go     | Pass/Fail | Submit on your own | 14/9     |
| 4      | Threads and Protection          | Pass/Fail | Submit as group    | 21/9     |
| 5      | Programming Tools               | Pass/Fail | Submit as group    | 5/10     |
| 6      | Linux Kernel IO Driver          | Graded    | Submit as group    | 26/10    |
| 7      | ChanStat: TV channel statistics | Graded    | Submit as group    | 16/11    |
