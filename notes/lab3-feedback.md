
##General feedback on lab3

Many of you have received **Not Approved** on lab3 (14 of 33 submissions). I wanted to summarize some of the general issues that prevented approval and provide other recommendations to everyone. For the most part I've reported these comments also to each individual student.

##Resubmission will be allowed until October 28th

If you have problems completing the lab3 assignments please seek help in the lab exercise time slot on Tuesdays (11-14). If you need more time, please get in touch.

You should resubmit your updated version to github with the commit message: `username lab3 resubmission`

**Approval of resubmissions will be done in the lab on October 28th.**

###Issue 1: Not enough tasks completed

I expect a minimum of two of the three main tasks to be completed to a satisfactory level. For a task to be considered completed to satisfaction, I expect:

1. It should compile without errors
2. It should run and do the main task expected of it according to the specification

This is the main reason for **Not Approving** a submission. Many of you may have some code that looks close to working, but have not passed the bar due to compile errors. 

For lab3, we use the following to compile your code (similar for the others):

```
cd techo
go build
```

This produces a single binary file in the `techo` folder called `techo`. We can run this with `techo -server` to start the server, and in another terminal we can run just `techo` without an argument to start a single client. This client is interactive, so it expects user input. The output on the client side should be the string echoed back from the server.

**Recommendation: Seek help to resolve compile errors before submitting.** Also remember that google can be quite useful to find help regarding compile error messages.

Most of you have managed to complete both echo programs, and that is enough for approval.


###Issue 2: Not following the provided code templates

Many of you have written some code that does not conform to the provided template source files. This makes it more difficult for us to test, especially if no instructions are provided.

Things you should **not do**:

1. Rename the package from `main` to something else.
2. Make new/separate package for the provided template source files
3. Provide your own command line parsing; this is already provided to allow us to test your code in standardized way.

Note that in lab7 you may wish to create new packages, but for lab3 it is expected that you do not mess around with the package structure.

In addition, it should be mentioned that there are lots of hints in the code provided for the two echo programs (TCP client and UDP server) that can be used to implement the other client and server. It is not a simple copy, but you should try to take advantage of the same code structure. This will make your code easily compatible with the provided counterpart program.

##Some other recommendations

These recommendations are not related to approval or not, but should help you avoid some common mistakes.

###Issue 1: Avoid adding backup files or binary files etc. to your git repository

Many of you have added, committed, and pushed binary files, backup files (typically ending with `.bak` or `~`), swap files (`.swo` and `.swp`), and other files that are unrelated and not useful to store in the git repository (e.g. `.DS_Store`). Binary files take up a lot of space and may not even run on my OS X system.

You can prevent that files gets accidentally added to the repository by adding a wildcard pattern, such as `*.bak` or `*~`, or an actual file name, e.g. `techo` to the `lab3/.gitignore` file. 

###Issue 2: One-way communication only

Some of you have only implemented one-way communication. While I've mostly accepted this, you should really read back the reply from the server and output it to the terminal.

###Issue 3: When committing changes

When committing changes to your local repository, please avoid touching many lab exercises at the same time and committing it under one commit. This is my recommended workflow:

1. Work on one task at a time (e.g. techo)
2. Commit changes as you finish different parts of it
3. Use whatever commit messages make sense to you
4. (Optional) Push your code changes to github
5. Repeat 1-4 for each task
6. When ready to submit it all, make a small *unimportant* change in one of the files
7. Commit with the message `username lab3 submission`
8. Push your submission to github (this makes it available for us)

###Issue 4: No lock on server-side RPC methods

This is how you can add a lock on the server-side RPC methods:

```
type KVStore struct {
        lock  *sync.Mutex
        store map[string]string
}

func (kv *KVStore) Insert(input Pair, reply *bool) error {
        kv.lock.Lock()
        defer kv.lock.Unlock()
        kv.store[input.Key] = input.Value
        *reply = true
        return nil
}
```

Note also that the `Keys()` method must also lock on access to the `kv.store` map. 

One optimization that can be done here is to use a read-write lock, and use read mode on `Lookup()` and `Keys()` and write mode on `Insert()`. 


