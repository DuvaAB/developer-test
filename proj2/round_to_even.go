package main

import(
    "os"
    "strconv"
    "encoding/csv"
)

func main() {
    in, _ := csv.NewReader(os.Stdin).ReadAll()
    out := csv.NewWriter(os.Stdout)

    for i, row := range in {
        if i == 0 { continue }
        x, err := strconv.Atoi(row[1])
        if err != nil { panic(err) }
        y, err := strconv.Atoi(row[2])
        if err != nil { panic(err) }
        z, err := strconv.Atoi(row[3])
        if err != nil { panic(err) }
        if x % 2 != 0 {
            x = x + 1
        }
        if y % 2 != 0 {
            y = x + 1
        }
        if z % 2 != 0 {
            z = x + 1
        }
        row[1] = strconv.Itoa(x)
        row[2] = strconv.Itoa(y)
        row[3] = strconv.Itoa(z)
        out.Write(row)
    }
    out.Flush()
}
