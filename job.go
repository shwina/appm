package appm

import "fmt"

type Job struct {
    Id string
    User string
    Ncpus int
    Mem int
    Walltime int
}

func (j Job) String() string {
    return fmt.Sprintf("%v\t%v\t%v\t%vgb\t%v", j.Id, j.User, j.Ncpus, j.Mem, SecondsToHMS(j.Walltime))
}
