# tasks
text
tag(s)
due - tomorrow@07:00 #config option, maybe
added - now()
completed - now() on completed
priority
    lowest  (zzz)
    low     (yawn)
    normal  (hmm)
    high    (sweat)
    highest (panic)
urgency - automatically determined based on priority, due date
comment
progress - dynamic, only if children
parent
state
  due
  overdue
  done
  abandoned




# optional config dir setup
```shell
aufgaben -c/--config /path/to/config/dir
```

# add a new task
```shell
aufgaben [-a/--add] what you have to do goes here
aufgaben [-a/--add] [-p/--priority value] text goes here
aufgaben [-a/--add] [-t/--tag value] text goes here
aufgaben [-a/--add] [-t/--tag "multi word value"] text goes here
aufgaben [-a/--add] [-t/--tag value] [-t/--tag another] text goes here
aufgaben [-a/--add] [--comment single word comment] # do we need this?
aufgaben [-a/--add] [--comment "multi word config"] # maybe use 'note' instead, open in vim or whatever, and write there?
aufgaben [-a/--add] [--parent hash]
```

# update new task
```shell
aufgaben [-u/--update hash] what you have to do goes here
and all the other options from add
```

# list tasks
```shell
aufgaben -l/--list [--order (u)rgency/(d)ue/(a)dded/(s)] [-t/--tag] terms to search
```


# flags
https://github.com/jessevdk/go-flags
https://lightstep.com/blog/getting-real-with-command-line-arguments-and-goflags/
