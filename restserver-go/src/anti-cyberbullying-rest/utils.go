/**
* Copyright (c) 2015 EMC Corporation
* All Rights Reserved
*
* This software contains the intellectual property of EMC Corporation
* or is licensed to EMC Corporation from third parties.  Use of this
* software and the intellectual property contained therein is expressly
* limited to the terms and conditions of the License Agreement under which
* it is provided by or on behalf of EMC.
*
**/

/**
*
* LogCourier utility module
*
**/

package main

import (
	"os/exec"
	"strings"
	"log"
	)





func CmdExec(cmd string)(output string, err error){
	parts := strings.Fields(cmd)
	out := []byte{};
    out, err = exec.Command(parts[0], parts[1:]...).Output();
    if err != nil {
        log.Println("error occured")
        log.Printf("%s", err)
        return "",err
    }
    log.Printf("cmd exec:%s", out)
    output = string(out)
    return output, nil
}

