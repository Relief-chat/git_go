package main

import (
    "fmt"
    "os"
    "os/exec"
)

func test1(userInput string) {

    cmdPath,_ := userInput;

    // bad
    cmd := &exec.Cmd {
        Path: cmdPath,
        Args: []string{ "foo", "bar" },
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }

    cmd.Start();

}

func test2(userInput string) {

    cmdPath,_ := exec.LookPath("foo");

    // bad
    cmd := &exec.Cmd {
        Path: cmdPath,
        Args: []string{ userInput, "bar" },
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }

    cmd.Start();

}

func test3(userInput string) {

    cmdPath,_ := exec.LookPath("bash");

    // bad
    cmd := &exec.Cmd {
        Path: cmdPath,
        Args: []string{ cmdPath, "-c", userInput },
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }

    cmd.Start();

}

func test4(userInput string) {

    cmdPath,_ := exec.LookPath("bash");

    args = []string{ cmdPath, "-c", userInput }

    // bad
    cmd := &exec.Cmd {
        Path: cmdPath,
        Args: args,
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }

    cmd.Start();

}

func okTest1(userInput string) {

    cmdPath,_ := exec.LookPath("go");

    // ok 
    cmd := &exec.Cmd {
        Path: cmdPath,
        Args: []string{ cmdPath, "bar" },
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }

    cmd.Start();

}
