package common

import (
	"testing"
)

func TestRun(t *testing.T) {
	testConn := ec2.EC2TestConn{}
	state := multistep.StateBag{}
	state.Put("ec2", testConn)
	stepRunSourceInstance := StepRunSourceInstance{}

	stepRunSourceInstance.Run(state)

	if !testConn.RunInstanceCalled {
		t.Error("RunInstance not called")
	}

	if testConn.RunInstanceParams.KeyName != "theKeyName" {
		t.Error("RunInstance called with wrong keyname")
	}
}