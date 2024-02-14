package main

import (
	"fmt"
	s "strings" // strings as s
)

func stringFunctions() {
	name := "sefa"
	fmt.Println(s.Count(name, "a"))     // how many "a" is there in <name> //* 1
	fmt.Println(s.Contains(name, "b"))  //* false
	fmt.Println(s.Index(name, "e"))     // index of first "e" in <name>. if not exists, returns -1 //* 1
	fmt.Println(s.ToUpper(name))        //* SEFA
	fmt.Println(s.ToLower(name))        //* sefa
	fmt.Println(s.HasPrefix(name, "e")) // = name.startsWith("e") //* false
	fmt.Println(s.HasSuffix(name, "a")) // = name.endsWith("a") //* true
	strArr := s.Split(name, "")         // = ["s","e","f","a"]
	newStr := s.Join(strArr, "*")       // = s*e*f*a

	fmt.Println(s.Replace(newStr, "*", "-", -1)) //* s-e-f-a
	// -1 is the "*"" amount the replace. Since it's given negative, there is no limit in this case

	fmt.Println(s.Repeat(name, 5)) //* = sefa * 5 = sefasefasefasefasefa
}
