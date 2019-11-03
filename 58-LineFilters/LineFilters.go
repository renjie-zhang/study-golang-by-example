/*
 * @Descripttion: 
 * @version: 
 * @Author: renjie.zhang
 * @Date: 2019-11-03 10:29:53
 * @LastEditTime: 2019-11-03 10:31:43
 */
package main

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
