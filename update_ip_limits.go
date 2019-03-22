package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

const maxPodsPath = "./files/eni-max-pods.txt"

func main() {
	cmd := exec.Command("sh", "-c", `curl -s https://raw.githubusercontent.com/awsdocs/amazon-ec2-user-guide/master/doc_source/using-eni.md | grep '^|' | tr -d '` + "`" + `' | sed -e '1,2d ; s/\\./\./g ; s/ *| */|/g' | cut -d '|' --output-delimiter=' ' -f2,3,4`)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	
	limits := make(map[string]int)
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan()  {
		f := strings.Fields(scanner.Text())
		
		instanceType := f[0]
		maxENI, _ := strconv.Atoi(f[1])
		maxIPPerENI, _ := strconv.Atoi(f[2])
		
		limits[instanceType] = maxENI * (maxIPPerENI - 1) + 2
	}
	
	b, err := ioutil.ReadFile(maxPodsPath)
	if err != nil {
		panic(err)
	}

	builder := bytes.NewBuffer(nil)
	scanner = bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		l := scanner.Text()
		if strings.HasPrefix(l, "#") {
			builder.WriteString(l)
			builder.WriteString("\n")
			continue
		}
		
		f := strings.Fields(l)
		
		instanceType := f[0]
		limit, ok := limits[instanceType]
		if !ok {
			panic("limit not found for instance type " + instanceType)
		}
		
		builder.WriteString(fmt.Sprintf("%s %d\n", instanceType, limit))
	}
	
	err = ioutil.WriteFile(maxPodsPath, builder.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}