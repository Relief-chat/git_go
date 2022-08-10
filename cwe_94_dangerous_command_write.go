import (
	"fmt"
	"os"
	"os/exec"
  )
  
  func test1(password string) {
	cmd := exec.Command("bash")
	cmdWriter, _ := cmd.StdinPipe()
	cmd.Start()
  
	cmdString := fmt.Sprintf("sshpass -p %s", password)
  
	// bad
	cmdWriter.Write([]byte(cmdString + "\n"))
  
	cmd.Wait()
  }
  
  func okTest1() {
	cmd := exec.Command("bash")
	cmdWriter, _ := cmd.StdinPipe()
	cmd.Start()
  
	// ok 
	cmdWriter.Write([]byte("sshpass -p 123\n"))
	cmdWriter.Write([]byte("exit"    + "\n"))
  
	cmd.Wait()
  }
  