##Discrepencies between the description and the zapserver code

There are some discrepencies between the pre-supplied lab naming in the
`runLab()` method of `zapserver` and the description. Please follow the
description, and change the `zapserver` code accordingly.

(Posted by Hein: Monday 13 Nov)

##Typo in the zaplab description

In Part 1 c) `simplestorage.go` should be replaced with `simplelogger.go`. 
This file resides in the `zlog` folder on github. 

(Posted by Hein: Monday 10 Nov)

##More remote access for *advanced users/developers*

*The following assumes that you have a publicly accessible IP address at your
home, i.e. you are not yourself behind a firewall or NAT router, unless you have
access rights and can configure port forwarding.*

*This tip is not a requirement for the lab; it is only meant to direct those of
you who are more adventurous.*

As I pointed out in my post yesterday, you cannot Dial into the badne machines
because of the firewall. You can only dial into port 22, but then you will only
hit the `ssh` server for interactive login sessions or command execution.

However, if you wish to work on your local system at home, you could set up a
machine with a public IP, and start your zapserver there, listening on the public
IP and some port that you decide (instead of the multicast address). The you
need to develop a forwarding service that needs listen to the multicast stream
and forward that to your home server. Then you can do all your coding on your
home server. That is, you first need to implement the forwarding service and run
that on one of the badne or pitter machines. By having the forwarding service make
connections from inside UiS to your home server, you can circumvent the firewall.
This should be fairly easy to do if you have solid skills, but I’ve not tried it.

This library *may* be helpful:
https://github.com/petar/GoTeleport

(Posted by Hein: Thursday 6 Nov)

##Remote access to work with the zaplab

As mentioned in my walkthrough of the zaplab on Tuesday, there is a firewall that
prevents direct `ssh` connections to the machines in the pitter-lab. However, 
you can use any of the following machines to gain access to the multicast stream:

`{badne5,badne6,badne7,badne8}.ux.uis.no`

However, you need to ensure that your `zapserver` source code is uploaded to
your Unix account before you compile it on one of the badne machines.
Alternatively, you may also compile the code locally on your own machine,
assuming you are working on Linux, and then copy the binary to your Unix account.

A few tips to ease this process of copying files from a local development machine
to your Unix account is to install [sshfs](http://fuse.sourceforge.net/sshfs.html) (Linux)
or [osxfuse](http://osxfuse.github.io) (Mac OS X). These tools allow you to mount
your Unix account folder as a remote drive on your local machine, meaning that you
can work almost as if your files were local (except for the added latency.)

From a Windows environment, I don't know which is the best tool, but
[WinSCP](http://winscp.net/eng/index.php)
seems to be one possibility that requires manually copying of files.

(Posted by Hein: Wednesday 5 Nov)

##No lectures next week after all

Today we wrapped up the theory lectures. Next week, you are free to work on the lab.
If you need help with lab7, feel free to stop by my office E424 in the 8:15-10:00
time frame on Monday and Tuesday.
Morten and Heine will be in the lab in the regular lab hours.

(Posted by Hein: Tuesday 4 Nov)

##Lecture also in Week 46

Due to my travels and illness earlier this semester, I will be lecturing also in Week 46.

(Posted by Hein: Sunday 26 Oct)

##Lab unavailable on November 18th

FYI: While I expect that both the lab and lectures will be wrapped up by November 18th,
I just wanted to let you know that another course has requested to use the Linux lab on
this day between 10-12.

(Posted by Hein: Tuesday 21 Oct)

##Lab exam info

I've posted a document in this folder titled `lab-exam.md` describing the lab exam procedure.

(Posted by Hein: Tuesday 21 Oct)

##Exams from previous years

I have posted exams and solutions from previous years on It's learning.
Note that we used a different book and different lab project in 2012,
so it may be somewhat less relevant.

(Posted by Hein: Tuesday 21 Oct)

##Use of inline assembly in gcc

```
//make a global lock variable
int asmMutex  = 1;
 
//To lock
asm("spin: lock btr $0, asmMutex ");
asm(" jnc spin ");
.
//To unlock
asm(" bts $0, asmMutex ");
```

##Use of commandline arguments i C
To get commandline arg into c try this:

``` 
#include<stdio.h>
#include<stdlib.h>
 
 
int main( int argc, char *argv[] )
{
        int i;
        printf("Number of arguments entered=%d\n", argc);
        printf("And the arguments are:\n");
        for(i=0;i<argc;i++)
                printf("Arg %d=%s\n", i, argv[i] );
        return 0;
}
``` 
Usage:
```
mortenm@pitter23:~/OpSys/2013/cmdline$ ./main 1 heisan hoppsan
Number of arguments entered=4
And the arguments are:
Arg 0=./main
Arg 1=1
Arg 2=heisan
Arg 3=hoppsan
mortenm@pitter23:~/OpSys/2013/cmdline$
``` 
PS: All arguments are 'strings' (char[]), and may need to be converted to in,float, etc. before furter usage.
Use atoi(), atof() etc....



##Share folder between host and guest on VirtualBox
It is possible to share folders between the host (the machine where the VirtualBox application is running on) and the virtual image.
 
You first need to set up a folder on the host (settings->shared folders) and give the share a name. 
Inside SlackWare you need to run the following command to mount the share:
 
Make a dir where you do the mount:
mkdir share (or another name you choose )
sudo mount -t vboxsf share ¨/share (where the first 'share' is the name given in settings->shared folder, and ~/share is the directory just created)
PS: mount is reserved to 'root' to execute. But by use of sudo we can run 'root' commands as 'user'. Use man sudo for more info


##Lab handin deadlines postponed

Due to the problems with github, we will be postponing the deadline for lab 2 until Sunday 14th, and likewise for the next labs. Please consult the lecture plan for the details. Tomorrow I'll walk through a few examples with git and github to explain a few things. I have also found a new room which should also have better space and facilities. For those who read this, please let others know that we will meet in E458 (across the hall from the eletro-lab).

(Posted by Hein: Sunday 7 Sep)

##Lab repositories update

If you find your repo for lab2 missing, it has most likely been deleted (the remaining will be deleted soon). The reason for deleting, see below.

(Posted by Hein: Friday 5 Sep)

##Lab repositories

Hey all. Due to a misunderstanding of github's instructions, I have created too many repositories (one per student per lab exercise). This does not scale, and yesterday we ran out of private repositories. Thus, we will instead be creating only one repository for all lab exercises, with lab2 and lab3 etc as subfolders. Thus, those who have already been given a repo for lab2, this will be deleted. If you have made changes locally related to lab2, please take care to keep your local copy safe, so that you can move those changes over to the new copy later. I will start to create new repositories soon, and I'm asking that you please begin using the new `username-labs` repos.

(Posted by Hein: Thursday 4 Sep)

##Lab 2 will be private

**We will need your githut account name.** You should not fork the lab2 repo, when you get access to it. Instead we will create a separate repo for each of you. To give us your github account name, clone this repo, and add your account name to the file `students`, and do a *Pull request.* (see the README.md description for lab1 for instructions related to doing a Pull request.)

(Posted by Hein: Tuesday 2 Sep)

##Lab 1 Handin

If you experience trouble with submitting lab 1 due to difficultites with git, don't worry. We will not write down any slip days due to this.

(Posted by Hein: Tuesday 2 Sep)

##External access to the Linux machines

Information about how to do a remote login to the Linux machines can be found here: http://wiki.ux.uis.no/foswiki/Info/WebHome and here: http://wiki.ux.uis.no/foswiki/Info/HvordanLoggeInnP%E5Unix-anlegget

##Important: Sign up for Unix account

To be effective in the lab on Tuesday, it is important that you sign up for an account on the Unix system prior to coming to the lab. The latest time to sign up is **Monday at 15:00**, in order to have your account ready for the lab on Tuesday. But there is no reason to delay; do it now! The instructions for sign up are here http://user.ux.uis.no/.

(Posted by Hein: Friday 22 Aug)

##Lab schedule update

Turns out that the lab room has a scheduling conflict so we won't have lab on Mondays. I've requested that the schedule be updated. You can choose when you come to the lab, but we, the teaching staff, will only be there at some given time (to be announced later). Attendance is not mandatory. Beyond that you can be in the lab when it is reserved for the course or when it is not being used by another course. 

